package server

import (
	"bytes"
	"encoding/json"
	"go-atlas-corp/domain/sector"
	"net/http"
	"net/http/httptest"
	"testing"

	sectormock "go-atlas-corp/domain/sector/mocks"

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

	mockSector := new(sectormock.MockSector)
	mockSector.On("CalculateLocation", sector.Coordinates{X: 123.12, Y: 456.56, Z: 789.89, Vel: 20.0}).Return(float64(1389.57))
	newServer := New(mockSector)

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

	mockSector := new(sectormock.MockSector)
	newServer := server{
		Sector: mockSector,
	}

	newServer.Router().ServeHTTP(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestValidate_Invalid(t *testing.T) {
	req := locationRequest{
		X:   nil,
		Y:   nil,
		Z:   nil,
		Vel: nil,
	}
	err := req.Validate()
	assert.NotNil(t, err)
}

func TestValidate_Valid(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0
	vel := 0.0
	req := locationRequest{
		X:   &x,
		Y:   &y,
		Z:   &z,
		Vel: &vel,
	}
	err := req.Validate()
	assert.Nil(t, err)
}
