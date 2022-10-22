package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationRequest struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Z   string `json:"z"`
	Vel string `json:"vel"`
}

type locationResult struct {
	Loc string `json:"loc"`
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	responseJson, err := json.Marshal(locationResult{
		Loc: "123",
	})
	if err != nil {
		http.Error(w, "Internal Error, please contact support", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
}
