package client

import (
	"bytes"
	"compress/flate"
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
		InternalDataFormat: (*DataFormat)(prtb.GetDformat()),
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
		RefDformat:    (*DataFormat)(header.RefDformat),
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

// Next implements Iterator
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

		// Reset currval
		internals := make([]*InternalMeta, len(header.DatasetMeta.InternalsMeta))
		cit.currval = CubeElem{
			Records:      make([]*Record, len(header.GroupedRecords.Records)),
			DatasetsMeta: &DatasetMeta{Internals: internals},
		}

		for i, r := range header.GroupedRecords.Records {
			cit.currval.Records[i] = recordFromPb(r)
		}
		for i, d := range header.DatasetMeta.InternalsMeta {
			cit.currval.DatasetsMeta.Internals[i] = NewInternalMetaFromPb(d)
		}
		if header.GetError() != "" {
			cit.currval.Err = header.GetError()
			return true
		}

		cit.currval.Shape = [3]int32{header.Shape.Dim1, header.Shape.Dim2, header.Shape.Dim3}
		cit.currval.DType = header.GetDtype()
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

// Header (global Header)
func (cit *CubeIterator) Header() CubeHeader {
	return cit.header
}

// Value implements Iterator
func (cit *CubeIterator) Value() *CubeElem {
	return &cit.currval
}

// Err implements Iterator
func (cit *CubeIterator) Err() error {
	return cit.err
}

func (c Client) getCube(req *pb.GetCubeRequest) (*CubeIterator, error) {
	if c.dlClient != nil && !req.HeadersOnly {
		req.HeadersOnly = true
		it, err := c.getCube(req)
		if err != nil {
			return nil, err
		}
		return c.dlClient.DownloadCube(it, Format(req.Format))
	}

	stream, err := c.gcc.GetCube(c.ctx, req)
	if err != nil {
		return nil, grpcError(err)
	}
	return NewCubeIterator(ClientStream{stream}, req.Size.Width, req.Size.Height)
}

// GetCubeFromRecords gets a cube from a list of records
func (c Client) GetCubeFromRecords(instancesID, recordsID []string, crs string, pix2crs [6]float64, sizeX, sizeY int64, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	return c.getCube(&pb.GetCubeRequest{
		RecordsLister:    &pb.GetCubeRequest_Records{Records: &pb.RecordIdList{Ids: recordsID}},
		InstancesId:      instancesID,
		Crs:              crs,
		PixToCrs:         &pb.GeoTransform{A: pix2crs[0], B: pix2crs[1], C: pix2crs[2], D: pix2crs[3], E: pix2crs[4], F: pix2crs[5]},
		Size:             &pb.Size{Width: int32(sizeX), Height: int32(sizeY)},
		CompressionLevel: int32(compression),
		HeadersOnly:      headersOnly,
		Format:           pb.FileFormat(format),
	})
}

// GetCube gets a cube from a list of filters
func (c Client) GetCube(instancesID []string, tags map[string]string, fromTime, toTime time.Time, crs string, pix2crs [6]float64, sizeX, sizeY int64, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	fromTs := timestamppb.New(fromTime)
	if err := fromTs.CheckValid(); err != nil {
		return nil, err
	}
	toTs := timestamppb.New(toTime)
	if err := toTs.CheckValid(); err != nil {
		return nil, err
	}

	return c.getCube(&pb.GetCubeRequest{
		RecordsLister:    &pb.GetCubeRequest_Filters{Filters: &pb.RecordFilters{Tags: tags, FromTime: fromTs, ToTime: toTs}},
		InstancesId:      instancesID,
		Crs:              crs,
		PixToCrs:         &pb.GeoTransform{A: pix2crs[0], B: pix2crs[1], C: pix2crs[2], D: pix2crs[3], E: pix2crs[4], F: pix2crs[5]},
		Size:             &pb.Size{Width: int32(sizeX), Height: int32(sizeY)},
		CompressionLevel: int32(compression),
		HeadersOnly:      headersOnly,
		Format:           pb.FileFormat(format),
	})
}

func (c Client) GetCubeFromTile(instancesID []string, tags map[string]string, fromTime, toTime time.Time, tile *Tile, format Format, compression int, headersOnly bool) (*CubeIterator, error) {
	fromTs := timestamppb.New(fromTime)
	if err := fromTs.CheckValid(); err != nil {
		return nil, err
	}
	toTs := timestamppb.New(toTime)
	if err := toTs.CheckValid(); err != nil {
		return nil, err
	}
	geoTransform := pb.GeoTransform{A: tile.Transform[0],
		B: tile.Transform[1], C: tile.Transform[2],
		D: tile.Transform[3], E: tile.Transform[4],
		F: tile.Transform[5]}
	size := pb.Size{Width: tile.Width, Height: tile.Height}
	return c.getCube(&pb.GetCubeRequest{
		RecordsLister:    &pb.GetCubeRequest_Filters{Filters: &pb.RecordFilters{Tags: tags, FromTime: fromTs, ToTime: toTs}},
		InstancesId:      instancesID,
		Crs:              tile.CRS,
		PixToCrs:         &geoTransform,
		Size:             &size,
		CompressionLevel: int32(compression),
		HeadersOnly:      headersOnly,
		Format:           pb.FileFormat(format),
	})
}

func (d DownloaderClient) DownloadCube(iter *CubeIterator, format Format) (*CubeIterator, error) {
	var dsMeta []*pb.DatasetMeta
	var groupedRecords []*pb.GroupedRecords
	for iter.Next() {
		headers := iter.Value()
		internals := headers.DatasetsMeta.Internals
		internalsMeta := make([]*pb.InternalMeta, len(internals))
		for i, element := range internals {
			m := pb.InternalMeta{
				ContainerUri:    element.ContainerURI,
				ContainerSubdir: element.ContainerSubDir,
				Bands:           element.Bands,
				Dformat:         (*pb.DataFormat)(element.InternalDataFormat),
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
	stream, err := d.gdcc.DownloadCube(d.ctx,
		&pb.GetCubeMetadataRequest{
			DatasetsMeta:   dsMeta,
			GroupedRecords: groupedRecords,
			RefDformat:     (*pb.DataFormat)(iter.header.RefDformat),
			ResamplingAlg:  iter.header.ResamplingAlg,
			PixToCrs:       iter.header.Geotransform,
			Crs:            iter.header.Crs,
			Size:           &pb.Size{Width: iter.header.Width, Height: iter.header.Height},
			Format:         pb.FileFormat(format),
		})
	if err != nil {
		return nil, grpcError(err)
	}
	return NewCubeIterator(DownloaderStream{stream}, iter.header.Width, iter.header.Height)
}
