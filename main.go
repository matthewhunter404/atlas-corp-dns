package main

import (
	"fmt"
	"go-atlas-corp/config"
	"go-atlas-corp/domain/sector"
	"go-atlas-corp/grpcserver"
	"go-atlas-corp/server"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")
	conf := config.Load()
	sector := sector.New(conf.SectorID)
	newServer := server.New(sector)
	switch conf.APIMode {
	case config.APIModeGPRC:
		grpcserver.NewGRPC(conf.Port, sector)
	case config.APIModeREST:
		http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), newServer.Router())
	default:
		log.Fatal("Invalid mode")
	}
}
