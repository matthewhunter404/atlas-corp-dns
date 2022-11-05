package grpcserver

import (
	"context"
	"go-atlas-corp/domain/sector"
	dns "go-atlas-corp/grpcserver/proto"
	"strconv"
)

type sectorGRPC struct {
	sector sector.Sector
	dns.UnimplementedSectorServiceServer
}

func (s *sectorGRPC) CalculateLocation(ctx context.Context, req *dns.LocationRequest) (*dns.LocationResult, error) {
	x, err := strconv.ParseFloat(req.X, 32)
	if err != nil {
		return nil, err
	}
	y, err := strconv.ParseFloat(req.Y, 32)
	if err != nil {
		return nil, err
	}
	z, err := strconv.ParseFloat(req.Z, 32)
	if err != nil {
		return nil, err
	}
	vel, err := strconv.ParseFloat(req.Vel, 32)
	if err != nil {
		return nil, err
	}

	coords := sector.Coordinates{
		X:   x,
		Y:   y,
		Z:   z,
		Vel: vel,
	}
	result := s.sector.CalculateLocation(coords)
	return &dns.LocationResult{Loc: result}, nil
}
