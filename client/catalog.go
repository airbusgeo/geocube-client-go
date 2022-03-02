package client

import (
	"bytes"
	"compress/flate"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	pb "github.com/airbusgeo/geocube-client-go/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Format pb.FileFormat

const (
	Format_Raw   = Format(pb.FileFormat_Raw)
	Format_GTiff = Format(pb.FileFormat_GTiff)
)

type CubeHeader struct {
	Count         int64
	NbDatasets    int64
	RefDformat    *DataFormat
	ResamplingAlg pb.Resampling
	Geotransform  *pb.GeoTransform
	Crs           string
	Width, Height int32
}

type CubeElem struct {
	Data         []byte
	Shape        [3]int32
	DType        pb.DataFormat_Dtype
	Records      []*Record
	DatasetsMeta *DatasetMeta
	Order        pb.ByteOrder
	Err          string
}

type CubeIterator struct {
	stream  CubeStream
	currval CubeElem
	header  CubeHeader
	err     error
}

type CubeStream interface {
	Recv() (CubeResponse, error)
}

type CubeResponse interface {
	GetGlobalHeader() *pb.GetCubeResponseHeader
	GetHeader() *pb.ImageHeader
	GetChunk() *pb.ImageChunk
}

type ClientStream struct{ pb.Geocube_GetCubeClient }

func (s ClientStream) Recv() (CubeResponse, error) {
	return s.Geocube_GetCubeClient.Recv()

}

type DownloaderStream struct {
	pb.GeocubeDownloader_DownloadCubeClient
}

func (s DownloaderStream) Recv() (CubeResponse, error) {
	return s.GeocubeDownloader_DownloadCubeClient.Recv()
}

type DatasetMeta struct {
	Internals []*InternalMeta
}

func NewInternalMetaFromPb(prtb *pb.InternalMeta) *InternalMeta {
	if prtb == nil {
		panic("Proto message is empty")
	}
	return &InternalMeta{
		ContainerURI:       prtb.GetContainerUri(),
		ContainerSubDir:    prtb.GetContainerSubdir(),
		Bands:              prtb.GetBands(),
		InternalDataFormat: NewDataFormatFromPb(prtb.GetDformat()),
		ExternalMinValue:   prtb.GetRangeMin(),
		ExternalMaxValue:   prtb.GetRangeMax(),
		Exponent:           prtb.GetExponent(),
	}
}

type InternalMeta struct {
	ContainerURI       string
	ContainerSubDir    string
	Bands              []int64
	InternalDataFormat *DataFormat
	ExternalMinValue   float64
	ExternalMaxValue   float64
	Exponent           float64
}

func NewCubeIterator(stream CubeStream, width, height int32) (*CubeIterator, error) {
	cit := CubeIterator{stream: stream}

	// Get global header
	resp := cit.next()
	if resp == nil {
		if cit.err != nil {
			return nil, cit.err
		}
		return nil, fmt.Errorf("empty response : expecting a global header")
	}
	header := resp.GetGlobalHeader()
	if header == nil {
		return nil, fmt.Errorf("excepting a global header")
	}
	cit.header = CubeHeader{
		Count:         header.Count,
		NbDatasets:    header.NbDatasets,
		RefDformat:    NewDataFormatFromPb(header.RefDformat),
		ResamplingAlg: header.ResamplingAlg,
		Geotransform:  header.Geotransform,
		Crs:           header.Crs,
		Width:         width,
		Height:        height,
	}
	return &cit, nil
}

func (cit *CubeIterator) next() CubeResponse {
	resp, err := cit.stream.Recv()
	if err != nil {
		if err != io.EOF {
			cit.err = err
		}
		return nil
	}
	return resp
}

// Header (global Header)
func (cit *CubeIterator) Header() CubeHeader {
	return cit.header
}

// Value returns a pointer to the current element. The pointer remains valid as long as cubeiterator is.
// The value is updated when cit.Next() is called.
// Check CubeElem.Err for a potential error on this element.
func (cit *CubeIterator) Value() *CubeElem {
	return &cit.currval
}

// Err returns the error if the iterator failed to go to next Element.
func (cit *CubeIterator) Err() error {
	return cit.err
}

// Next loads the next element of the cube and returns true.
// It returns false if there is no more elements or an error occured.
// In the latter case, check .Err().
func (cit *CubeIterator) Next() bool {
	var resp CubeResponse

	// Get header
	var header *pb.ImageHeader
	var data bytes.Buffer
	{
		if resp = cit.next(); resp == nil {
			return false
		}
		// Parse header
		if header = resp.GetHeader(); header == nil {
			cit.err = errors.New("fatal: excepting a header")
			return false
		}

		// Reset currval (and the values returned by Value())
		internals := make([]*InternalMeta, len(header.DatasetMeta.InternalsMeta))
		cit.currval = CubeElem{
			Records:      make([]*Record, len(header.GroupedRecords.Records)),
			DatasetsMeta: &DatasetMeta{Internals: internals},
			Order:        header.Order,
			Err:          header.GetError(),
			DType:        header.GetDtype(),
		}

		for i, r := range header.GroupedRecords.Records {
			cit.currval.Records[i] = recordFromPb(r)
		}
		for i, d := range header.DatasetMeta.InternalsMeta {
			cit.currval.DatasetsMeta.Internals[i] = NewInternalMetaFromPb(d)
		}
		if cit.currval.Err != "" {
			return true
		}
		cit.currval.Shape = [3]int32{header.Shape.Dim1, header.Shape.Dim2, header.Shape.Dim3}

		data.Grow(int(header.GetSize()))
		data.Write(header.GetData())
	}

	// Get chunks
	for i := int32(1); i < header.NbParts; i++ {
		if resp = cit.next(); resp == nil {
			return false
		}

		// Parse chunk
		chunk := resp.GetChunk()
		if chunk == nil || chunk.GetPart() != i {
			cit.err = errors.New("fatal: excepting a chunk")
			return false
		}
		data.Write(chunk.GetData())
	}

	if header.NbParts > 0 {
		if header.Compression {
			inflater := flate.NewReader(&data)
			var b bytes.Buffer
			b.Grow(int(cit.currval.Shape[0]) * int(cit.currval.Shape[1]) * int(cit.currval.Shape[2]) * sizeOf(cit.currval.DType))

			if _, err := io.Copy(&b, inflater); err != nil {
				cit.currval.Err = err.Error()
			} else if err := inflater.Close(); err != nil {
				cit.currval.Err = err.Error()
			} else {
				cit.currval.Data = b.Bytes()
			}
		} else {
			cit.currval.Data = data.Bytes()
		}
	}
	return true
}

func sizeOf(dt pb.DataFormat_Dtype) int {
	switch dt {
	case pb.DataFormat_UInt8:
		return 1
	case pb.DataFormat_Int16, pb.DataFormat_UInt16:
		return 2
	case pb.DataFormat_Int32, pb.DataFormat_UInt32, pb.DataFormat_Float32:
		return 4
	case pb.DataFormat_Float64, pb.DataFormat_Complex64:
		return 8
	}
	return 0
}

type FromOption func(*fromOptions)

type fromOptions struct {
	RecordsID        []string
	GroupedRecordsID [][]string
	Tags             map[string]string
	FromTime, ToTime time.Time
}

// FromRecords to get a cube from records
func FromRecords(recordsID []string) func(*fromOptions) {
	return func(fo *fromOptions) {
		fo.RecordsID = recordsID
	}
}

// FromGroupeRecords to get a cube from a list of grouped records
func FromGroupeRecords(groupedRecordsID [][]string) func(*fromOptions) {
	return func(fo *fromOptions) {
		fo.GroupedRecordsID = groupedRecordsID
	}
}

// FromFilters to get a cube from a list of tags and an interval of time
func FromFilters(tags map[string]string, fromTime, toTime time.Time) func(*fromOptions) {
	return func(fo *fromOptions) {
		fo.Tags = tags
		fo.FromTime = fromTime
		fo.ToTime = toTime
	}
}

type TileOption func(*tileOptions)

type tileOptions struct {
	Crs          string
	Pix2Crs      [6]float64
	SizeX, SizeY int32
}

// OnTile to get data for the given tile
func OnTile(tile *Tile) func(*tileOptions) {
	return func(to *tileOptions) {
		to.Crs = tile.CRS
		to.Pix2Crs = tile.Transform
		to.SizeX = tile.Width
		to.SizeY = tile.Height
	}
}

// OnGeoTransform to get data for the given geoTransform
func OnGeoTransform(crs string, pix2crs [6]float64, sizeX, sizeY int32) func(*tileOptions) {
	return func(to *tileOptions) {
		to.Crs = crs
		to.Pix2Crs = pix2crs
		to.SizeX = sizeX
		to.SizeY = sizeY
	}
}

type GetCubeOption func(*getCubeOptions)

type getCubeOptions struct {
	HeadersOnly bool
	Compression int32
	Format      Format
	DataFormat  *DataFormat
	Predownload bool
}

// WithCompression to compress data from the server to the client
func WithCompression(compression int) func(*getCubeOptions) {
	return func(gco *getCubeOptions) {
		gco.Compression = int32(compression)
	}
}

// WithHeadersOnly to only get headers
func WithHeadersOnly() func(*getCubeOptions) {
	return func(gco *getCubeOptions) {
		gco.HeadersOnly = true
	}
}

// WithFileFormat to get a file instead of row data
func WithFileFormat(f Format) func(*getCubeOptions) {
	return func(gco *getCubeOptions) {
		gco.Format = f
	}
}

// The following options only apply when using a downloader)

// WithDataFormat to change the dataformat of the output
func WithDataFormat(df *DataFormat) func(*getCubeOptions) {
	return func(gco *getCubeOptions) {
		gco.DataFormat = df
	}
}

// WithPredownload advise the downloader to predownload the files before merging them. When the dataset is remote and the whole dataset is required, it is more efficient to predownload it
func WithPredownload() func(*getCubeOptions) {
	return func(gco *getCubeOptions) {
		gco.Predownload = true
	}
}

// GetCube gets a cube from a list of records, groupedRecords or filters
func (c Client) GetCube(ctx context.Context, fromOpt FromOption, instanceID string, onTileOpt TileOption, opts ...GetCubeOption) (*CubeIterator, error) {
	onTile := tileOptions{}
	onTileOpt(&onTile)
	req := &pb.GetCubeRequest{
		InstancesId: []string{instanceID},
		Crs:         onTile.Crs,
		PixToCrs:    &pb.GeoTransform{A: onTile.Pix2Crs[0], B: onTile.Pix2Crs[1], C: onTile.Pix2Crs[2], D: onTile.Pix2Crs[3], E: onTile.Pix2Crs[4], F: onTile.Pix2Crs[5]},
		Size:        &pb.Size{Width: onTile.SizeX, Height: onTile.SizeY},
	}

	from := fromOptions{}
	{
		fromOpt(&from)
		if len(from.RecordsID) > 0 {
			req.RecordsLister = &pb.GetCubeRequest_Records{Records: &pb.RecordIdList{Ids: from.RecordsID}}
		} else if len(from.GroupedRecordsID) > 0 {
			groupedRecords := make([]*pb.GroupedRecordIds, 0, len(from.GroupedRecordsID))
			for _, records := range from.GroupedRecordsID {
				groupedRecords = append(groupedRecords, &pb.GroupedRecordIds{Ids: records})
			}
			req.RecordsLister = &pb.GetCubeRequest_GroupedRecords{GroupedRecords: &pb.GroupedRecordIdsList{Records: groupedRecords}}
		} else {
			fromTs := timestamppb.New(from.FromTime)
			if err := fromTs.CheckValid(); err != nil {
				return nil, err
			}
			toTs := timestamppb.New(from.ToTime)
			if err := toTs.CheckValid(); err != nil {
				return nil, err
			}
			req.RecordsLister = &pb.GetCubeRequest_Filters{Filters: &pb.RecordFilters{Tags: from.Tags, FromTime: fromTs, ToTime: toTs}}
		}
	}

	options := getCubeOptions{}
	{
		for _, opt := range opts {
			opt(&options)
		}
		req.CompressionLevel = options.Compression
		req.HeadersOnly = options.HeadersOnly
		req.Format = pb.FileFormat(options.Format)
	}

	return c.getCube(ctx, req, opts...)
}

func (c Client) getCube(ctx context.Context, req *pb.GetCubeRequest, opts ...GetCubeOption) (*CubeIterator, error) {
	if c.dlClient != nil && !req.HeadersOnly {
		req.HeadersOnly = true
		it, err := c.getCube(ctx, req)
		if err != nil {
			return nil, err
		}
		return c.dlClient.DownloadCube(ctx, it, Format(req.Format), opts...)
	}

	stream, err := c.gcc.GetCube(ctx, req)
	if err != nil {
		return nil, grpcError(err)
	}
	return NewCubeIterator(ClientStream{stream}, req.Size.Width, req.Size.Height)
}

// GetCubeFromRecords gets a cube from a list of records
//
// Deprecated: Use GetCube(..., WithRecords(), ...)
func (c Client) GetCubeFromRecords(ctx context.Context, recordsID []string, instanceID string, crs string, pix2crs [6]float64, sizeX, sizeY int32, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	return c.getCube(ctx, &pb.GetCubeRequest{
		RecordsLister:    &pb.GetCubeRequest_Records{Records: &pb.RecordIdList{Ids: recordsID}},
		InstancesId:      []string{instanceID},
		Crs:              crs,
		PixToCrs:         &pb.GeoTransform{A: pix2crs[0], B: pix2crs[1], C: pix2crs[2], D: pix2crs[3], E: pix2crs[4], F: pix2crs[5]},
		Size:             &pb.Size{Width: sizeX, Height: sizeY},
		CompressionLevel: int32(compression),
		HeadersOnly:      headersOnly,
		Format:           pb.FileFormat(format),
	})
}

// GetCubeFromGroupedRecords gets a cube from a list of grouped records
//
// Deprecated: Use GetCube(..., WithGroupedRecords(), ...)
func (c Client) GetCubeFromGroupedRecords(ctx context.Context, recordsID [][]string, instanceID string, crs string, pix2crs [6]float64, sizeX, sizeY int32, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	groupedRecords := make([]*pb.GroupedRecordIds, 0, len(recordsID))
	for _, records := range recordsID {
		groupedRecords = append(groupedRecords, &pb.GroupedRecordIds{Ids: records})
	}

	return c.getCube(ctx, &pb.GetCubeRequest{
		RecordsLister:    &pb.GetCubeRequest_GroupedRecords{GroupedRecords: &pb.GroupedRecordIdsList{Records: groupedRecords}},
		InstancesId:      []string{instanceID},
		Crs:              crs,
		PixToCrs:         &pb.GeoTransform{A: pix2crs[0], B: pix2crs[1], C: pix2crs[2], D: pix2crs[3], E: pix2crs[4], F: pix2crs[5]},
		Size:             &pb.Size{Width: sizeX, Height: sizeY},
		CompressionLevel: int32(compression),
		HeadersOnly:      headersOnly,
		Format:           pb.FileFormat(format),
	})
}

// GetCubeFromFilters gets a cube from a list of filters
//
// Deprecated: Use GetCube(..., WithFilters(), ...)
func (c Client) GetCubeFromFilters(ctx context.Context, tags map[string]string, fromTime, toTime time.Time, instanceID string, crs string, pix2crs [6]float64, sizeX, sizeY int32, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	fromTs := timestamppb.New(fromTime)
	if err := fromTs.CheckValid(); err != nil {
		return nil, err
	}
	toTs := timestamppb.New(toTime)
	if err := toTs.CheckValid(); err != nil {
		return nil, err
	}

	return c.getCube(ctx, &pb.GetCubeRequest{
		RecordsLister:    &pb.GetCubeRequest_Filters{Filters: &pb.RecordFilters{Tags: tags, FromTime: fromTs, ToTime: toTs}},
		InstancesId:      []string{instanceID},
		Crs:              crs,
		PixToCrs:         &pb.GeoTransform{A: pix2crs[0], B: pix2crs[1], C: pix2crs[2], D: pix2crs[3], E: pix2crs[4], F: pix2crs[5]},
		Size:             &pb.Size{Width: sizeX, Height: sizeY},
		CompressionLevel: int32(compression),
		HeadersOnly:      headersOnly,
		Format:           pb.FileFormat(format),
	})
}

// GetCubeFromTile gets a cube from a tile and a list of records or filters
//
// Deprecated: use GetCube(..., OnTile(), ...)
func (c Client) GetCubeFromTile(ctx context.Context, tile *Tile, instanceID string, recordIDs []string, gRecordIDs [][]string, tags map[string]string, fromTime, toTime time.Time, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	if recordIDs != nil {
		if gRecordIDs != nil || !toTime.IsZero() || !fromTime.IsZero() || len(tags) != 0 {
			return nil, fmt.Errorf("records, groupedRecords and time or tags are defined: records, grecords and filters are mutually exclusive")
		}
		return c.GetCubeFromRecords(ctx, recordIDs, instanceID, tile.CRS, tile.Transform, tile.Width, tile.Height, format, compression, headersOnly)
	}
	if gRecordIDs != nil {
		if !toTime.IsZero() || !fromTime.IsZero() || len(tags) != 0 {
			return nil, fmt.Errorf("groupedRecords and time or tags are defined: grecords and filters are mutually exclusive")
		}
		return c.GetCubeFromGroupedRecords(ctx, gRecordIDs, instanceID, tile.CRS, tile.Transform, tile.Width, tile.Height, format, compression, headersOnly)

	}
	return c.GetCubeFromFilters(ctx, tags, fromTime, toTime, instanceID, tile.CRS, tile.Transform, tile.Width, tile.Height, format, compression, headersOnly)
}

func (d DownloaderClient) DownloadCube(ctx context.Context, iter *CubeIterator, format Format, opts ...GetCubeOption) (*CubeIterator, error) {
	options := getCubeOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	var dsMeta []*pb.DatasetMeta
	var groupedRecords []*pb.GroupedRecords
	for headers := iter.Value(); iter.Next(); {
		internals := headers.DatasetsMeta.Internals
		internalsMeta := make([]*pb.InternalMeta, len(internals))
		for i, element := range internals {
			m := pb.InternalMeta{
				ContainerUri:    element.ContainerURI,
				ContainerSubdir: element.ContainerSubDir,
				Bands:           element.Bands,
				Dformat:         element.InternalDataFormat.toPb(),
				RangeMin:        element.ExternalMinValue,
				RangeMax:        element.ExternalMaxValue,
				Exponent:        element.Exponent,
			}
			internalsMeta[i] = &m
		}
		records := make([]*pb.Record, len(headers.Records))
		for i, element := range headers.Records {
			r := element.ToPb()
			records[i] = &r
		}
		groupedRecords = append(groupedRecords, &pb.GroupedRecords{Records: records})
		dsMeta = append(dsMeta, &pb.DatasetMeta{InternalsMeta: internalsMeta})
	}
	toDataFormat := iter.header.RefDformat
	if options.DataFormat != nil {
		toDataFormat = options.DataFormat
	}
	stream, err := d.gdcc.DownloadCube(ctx,
		&pb.GetCubeMetadataRequest{
			DatasetsMeta:   dsMeta,
			GroupedRecords: groupedRecords,
			RefDformat:     toDataFormat.toPb(),
			ResamplingAlg:  iter.header.ResamplingAlg,
			PixToCrs:       iter.header.Geotransform,
			Crs:            iter.header.Crs,
			Size:           &pb.Size{Width: iter.header.Width, Height: iter.header.Height},
			Format:         pb.FileFormat(format),
			Predownload:    options.Predownload,
		})
	if err != nil {
		return nil, grpcError(err)
	}
	return NewCubeIterator(DownloaderStream{stream}, iter.header.Width, iter.header.Height)
}

type Dataset struct {
	RecordId        string
	InstanceId      string
	ContainerSubdir string
	Bands           []int64
	Dformat         *DataFormat
	RealMinValue    float64
	RealMaxValue    float64
	Exponent        float64
}
type Container struct {
	Uri      string
	Managed  bool
	Datasets []*Dataset
}

func NewContainerFromPb(c *pb.Container) *Container {
	datasets := make([]*Dataset, 0, len(c.Datasets))
	for _, d := range c.Datasets {
		datasets = append(datasets, &Dataset{
			RecordId:        d.RecordId,
			InstanceId:      d.InstanceId,
			ContainerSubdir: d.ContainerSubdir,
			Bands:           d.Bands,
			Dformat:         NewDataFormatFromPb(d.Dformat),
			RealMinValue:    d.RealMinValue,
			RealMaxValue:    d.RealMaxValue,
			Exponent:        d.Exponent,
		})
	}
	return &Container{
		Uri:      c.Uri,
		Managed:  c.Managed,
		Datasets: datasets,
	}
}

// GetContainers gets information on containers
func (c Client) GetContainers(ctx context.Context, uris []string) ([]*Container, error) {
	resp, err := c.gcc.GetContainers(ctx, &pb.GetContainersRequest{
		Uris: uris,
	})

	containers := make([]*Container, 0, len(resp.Containers))
	for _, c := range resp.Containers {
		containers = append(containers, NewContainerFromPb(c))
	}

	return containers, grpcError(err)
}
