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

func TestFetchLoc(t *testing.T) {
	requestBody, err := json.Marshal(locationRequest{
		X:   123.12,
		Y:   456.56,
		Z:   789.89,
		Vel: 20.0,
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
