package grpcserver

import (
	"fmt"
	"go-atlas-corp/domain/sector"
	dns "go-atlas-corp/grpcserver/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewGRPC(port int, sector sector.Sector) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	dns.RegisterSectorServiceServer(s, &sectorGPRC{sector, dns.UnimplementedSectorServiceServer{}})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
