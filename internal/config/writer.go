package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Save(
	path string,
	cfg *GinForgeConfig,
) error {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		data,
		0644,
	)
}
