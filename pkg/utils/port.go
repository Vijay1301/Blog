package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigPath = "pkg/config/"
)

func PortConfig(env string) (*Config, error) {
	var config *Config

	filename := env + ".json"

	configFile := filepath.Join(ConfigPath, filename)
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return config, nil
}
