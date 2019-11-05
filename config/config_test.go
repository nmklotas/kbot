package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanReadConfig(t *testing.T) {
	assert := assert.New(t)
	config, err := ReadConfig("config-test")
	if err != nil {
		t.Errorf("Failed to read config")
	}

	assert.NotEqual(config.ServerUrl, "")
}
