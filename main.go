package main

import (
	"fmt"
	"go-atlas-corp/config"
	"go-atlas-corp/domain/sector"
	"go-atlas-corp/server"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")
	config := config.Load()
	sector := sector.New(config.SectorID)
	newServer := server.New(sector)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), newServer.Router())
}
