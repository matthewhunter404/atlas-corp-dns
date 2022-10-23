package sector_test

import (
	"go-atlas-corp/domain/sector"
	"testing"
)

func TestCalculateLocation(t *testing.T) {
	testSector := sector.Sector{
		ID: 1,
	}
	testInput := sector.Coordinates{
		X:   123.12,
		Y:   456.56,
		Z:   789.89,
		Vel: 20.0,
	}
	want := 1389.57
	got := testSector.CalculateLocation(testInput)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
