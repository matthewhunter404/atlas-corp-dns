package server

import (
	"bytes"
	"encoding/json"
	"go-atlas-corp/domain/sector"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchLoc_Success(t *testing.T) {
	x := 123.12
	y := 456.56
	z := 789.89
	vel := 20.0
	requestBody, err := json.Marshal(locationRequest{
		X:   &x,
		Y:   &y,
		Z:   &z,
		Vel: &vel,
	})
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("POST", "/location", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()
	sector := sector.New(uint(1))

	newServer := server{
		Sector: sector,
	}

	newServer.Router().ServeHTTP(w, req)

	resp := w.Result()

	var result locationResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 1389.57, result.Loc)
}

func TestFetchLoc_BadRequest(t *testing.T) {
	testRequestBody := []byte("\"x\": 123.12,\"y\": 456.56,\"z\": 789.89,\"vel\": 20.0")
	req := httptest.NewRequest("POST", "/location", bytes.NewReader(testRequestBody))
	w := httptest.NewRecorder()
	sector := sector.New(uint(1))

	newServer := server{
		Sector: sector,
	}

	newServer.Router().ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
