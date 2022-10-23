package server

import (
	"encoding/json"
	"fmt"
	sector "go-atlas-corp/domain/sector"
	"net/http"
)

type locationRequest struct {
	X   float64 `json:"x,string"`
	Y   float64 `json:"y,string"`
	Z   float64 `json:"z,string"`
	Vel float64 `json:"vel,string"`
}

type locationResult struct {
	Loc float64 `json:"loc"`
}

// fetchLoc is a handler function stub
func (s *Server) fetchLoc(w http.ResponseWriter, r *http.Request) {
	var req locationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body JSON"))
		return
	}
	fmt.Printf("req: %+v\n", req)

	coords := sector.Coordinates{
		X:   req.X,
		Y:   req.Y,
		Z:   req.Z,
		Vel: req.Vel,
	}
	location := s.Sector.CalculateLocation(coords)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(locationResult{
		Loc: location,
	})
	if err != nil {
		http.Error(w, "Internal Error, please contact support", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
}
