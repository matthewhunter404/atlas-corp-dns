package config_test

import (
	"go-atlas-corp/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad_EnvVariables(t *testing.T) {
	t.Setenv("atlas_dns_sector_id", "45")
	t.Setenv("atlas_dns_port", "8080")
	actualFields := config.Load()

	assert.Equal(t, uint(45), actualFields.SectorID)
	assert.Equal(t, 8080, actualFields.Port)
}

func TestLoad_DefaultValues(t *testing.T) {
	actualFields := config.Load()

	assert.Equal(t, uint(1), actualFields.SectorID)
	assert.Equal(t, 3000, actualFields.Port)
}
