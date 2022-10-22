package main

import (
	"fmt"
	sector "go-atlas-corp/domain/sector"
	"go-atlas-corp/server"
	"net/http"
)

const SectorID = 1

func main() {
	fmt.Println("Server starting...")
	sector := sector.Sector{
		ID: uint(SectorID),
	}
	newServer := server.Server{
		Sector: sector,
	}
	http.ListenAndServe(fmt.Sprintf(":%d", 3000), server.NewRouter(&newServer))
}
