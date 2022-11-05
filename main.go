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
	switch conf.APIMode {
	case config.APIModeGRPC:
		fmt.Printf("GRPC interface listening on port: %d\n", conf.Port)
		grpcserver.NewGRPC(conf.Port, sector)
	case config.APIModeREST:
		fmt.Printf("REST interface listening on port: %d\n", conf.Port)
		newServer := server.New(sector)
		http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), newServer.Router())
	default:
		log.Fatal("Invalid mode")
	}
}
