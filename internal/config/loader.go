package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*GinForgeConfig, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg GinForgeConfig

	if err := yaml.Unmarshal(
		data,
		&cfg,
	); err != nil {
		return nil, err
	}

	return &cfg, nil
}
