package sectormock

import (
	"go-atlas-corp/domain/sector"

	"github.com/stretchr/testify/mock"
)

type MockSector struct {
	mock.Mock
}

func (m *MockSector) CalculateLocation(input sector.Coordinates) float64 {
	args := m.Called(input)
	return args.Get(0).(float64)

}
