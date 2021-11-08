package client

import (
	"fmt"
	"io"
	"strings"

	pb "github.com/airbusgeo/geocube-client-go/pb"
)

type Layout pb.Layout

type Tile struct {
	Transform     [6]float64
	CRS           string
	Width, Height int32
	err           error
}

// TileError is an erroneous tile
func TileError(err error) Tile {
	return Tile{err: err}
}

// Error returns an error if the tile is an erroneous tile
func (t *Tile) Error() error {
	return t.err
}

func NewTileFromPb(pbt *pb.Tile) *Tile {
	return &Tile{
		Transform: [6]float64{pbt.Transform.A, pbt.Transform.B, pbt.Transform.C, pbt.Transform.D, pbt.Transform.E, pbt.Transform.F},
		CRS:       pbt.Crs,
		Width:     pbt.SizePx.Width,
		Height:    pbt.SizePx.Height,
	}
}

func (c Client) CreateLayout(name string, gridFlags []string, gridParameters map[string]string, blockXSize, blockYSize, maxRecords int64) error {
	if _, err := c.gcc.CreateLayout(c.ctx,
		&pb.CreateLayoutRequest{Layout: &pb.Layout{
			Name:           name,
			GridFlags:      gridFlags,
			GridParameters: gridParameters,
			BlockXSize:     blockXSize,
			BlockYSize:     blockYSize,
			MaxRecords:     maxRecords}}); err != nil {
		return grpcError(err)
	}

	return nil
}

func (c Client) ListLayouts(nameLike string) ([]*Layout, error) {
	resp, err := c.gcc.ListLayouts(c.ctx, &pb.ListLayoutsRequest{NameLike: nameLike})

	if err != nil {
		return nil, grpcError(err)
	}

	var layouts []*Layout
	for _, l := range resp.Layouts {
		layouts = append(layouts, (*Layout)(l))
	}

	return layouts, nil
}

func (c Client) TileAOI(aoi AOI, crs string, resolution float32, width_px, height_px int32) (<-chan Tile, error) {
	stream, err := c.gcc.TileAOI(c.ctx,
		&pb.TileAOIRequest{
			Aoi:        pbFromAOI(aoi),
			Crs:        crs,
			Resolution: resolution,
			SizePx:     &pb.Size{Width: width_px, Height: height_px},
		})

	if err != nil {
		return nil, grpcError(err)
	}

	tiles := make(chan Tile)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				tiles <- TileError(err)
				break
			}
			for _, tile := range resp.Tiles {
				tiles <- *NewTileFromPb(tile)
			}
		}
		close(tiles)
	}()

	return tiles, nil
}

// ToString returns a string with a representation of the layout
func (l *Layout) ToString() string {
	s := fmt.Sprintf("Layout %s:\n"+
		"  Block XSize:     %d\n"+
		"  Block YSize:     %d\n"+
		"  Max records:     %d\n"+
		"  Grid flags:      %s\n"+
		"  Grid parameters:\n",
		l.Name, l.BlockXSize, l.BlockYSize, l.MaxRecords, strings.Join(l.GridFlags, " "))
	appendDict(l.GridParameters, &s)
	return s
}
