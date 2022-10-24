package server

import (
	"encoding/json"
	"errors"
	"fmt"
	sector "go-atlas-corp/domain/sector"
	"net/http"
)

type locationRequest struct {
	X   *float64 `json:"x,string"`
	Y   *float64 `json:"y,string"`
	Z   *float64 `json:"z,string"`
	Vel *float64 `json:"vel,string"`
}

type locationResult struct {
	Loc float64 `json:"loc"`
}

// Validate adds a layer of validation not available in the standard Go JSON package,
// checking the locationRequest inputs to ensure none are missing. This effectively
// acts as a counterpart to DisallowUnknownFields()
func (lr locationRequest) Validate() error {
	if lr.X == nil {
		return errors.New("invalid input, X value missing")
	}
	if lr.Y == nil {
		return errors.New("invalid input, Y value missing")
	}
	if lr.Z == nil {
		return errors.New("invalid input, Z value missing")
	}
	if lr.Vel == nil {
		return errors.New("invalid input, Vel value missing")
	}
	return nil
}

// fetchLoc is a handler function that calls the sector method CalculateLocation()
func (s *server) fetchLoc(w http.ResponseWriter, r *http.Request) {
	var req locationRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body JSON"))
		return
	}

	err = req.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	coords := sector.Coordinates{
		X:   *req.X,
		Y:   *req.Y,
		Z:   *req.Z,
		Vel: *req.Vel,
	}
	location := s.Sector.CalculateLocation(coords)

	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(locationResult{
		Loc: location,
	})
	if err != nil {
		http.Error(w, "Internal Error, please contact support", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}
