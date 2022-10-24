package main

import (
	"fmt"
	"go-atlas-corp/domain/sector"
	"go-atlas-corp/server"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Server starting...")
	sectorID, err := strconv.Atoi(os.Getenv("atlas_dns_sector_id"))
	if err != nil {
		log.Fatalf("Invalid sector ID env variable detected, error: %v", err)
	}
	os.Getenv("atlas_dns_port")
	port, err := strconv.Atoi(os.Getenv("atlas_dns_port"))
	if err != nil {
		log.Fatalf("Invalid port env variable detected, error: %v", err)
	}

	sector := sector.New(uint(sectorID))
	newServer := server.New(sector)
	http.ListenAndServe(fmt.Sprintf(":%d", port), newServer.Router())
}
