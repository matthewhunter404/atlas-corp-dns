package main

import (
	"fmt"
	"go-atlas-corp/domain/sector"
	"go-atlas-corp/server"
	"net/http"
)

const SectorID = 1

func main() {
	fmt.Println("Server starting...")
	sector := sector.New(uint(SectorID))
	newServer := server.New(sector)
	http.ListenAndServe(fmt.Sprintf(":%d", 3000), newServer.Router())
}
