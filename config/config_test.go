package config

import "testing"

func TestCanReadConfig(t *testing.T) {
    config, err := ReadConfig("config-test")
    if err != nil {
        t.Errorf("Failed to read config")
    }
	if config.ServerUrl == "" {
		t.Errorf("ServerUrl is not read")
	}
}
