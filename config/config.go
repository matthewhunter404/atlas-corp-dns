package config

import (
	"log"
	"os"
	"strconv"
)

type Fields struct {
	SectorID uint
	Port     int
	APIMode  string
}

const APIModeREST = "REST"
const APIModeGRPC = "GRPC"

func Load() Fields {
	//Load up default values
	fields := Fields{
		SectorID: 1,
		Port:     3000,
	}
	sectorIDEnv := os.Getenv("atlas_dns_sector_id")
	if sectorIDEnv != "" {
		sid, err := strconv.Atoi(sectorIDEnv)
		if err != nil {
			log.Fatalf("Invalid sector ID env variable detected, error: %v", err)
		}
		fields.SectorID = uint(sid)
	}

	portEnv := os.Getenv("atlas_dns_port")
	if portEnv != "" {
		port, err := strconv.Atoi(portEnv)
		if err != nil {
			log.Fatalf("Invalid port env variable detected, error: %v", err)
		}
		fields.Port = port
	}
	apiModeEnv := os.Getenv("api_mode")
	switch apiModeEnv {
	case APIModeGRPC:
		fields.APIMode = apiModeEnv
	case APIModeREST:
		fields.APIMode = apiModeEnv
	default:
		log.Fatal("Invalid mode")
	}
	return fields
}
