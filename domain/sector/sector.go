package Sector

type Sector struct {
	ID uint
}

type Coordinates struct {
	X   float64
	Y   float64
	Z   float64
	Vel float64
}

func (s Sector) CalculateLocation(input Coordinates) float64 {
	sectorID := float64(s.ID)
	return input.X*sectorID + input.Y*sectorID + input.Z*sectorID + input.Vel
}
