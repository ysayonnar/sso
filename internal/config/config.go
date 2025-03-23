package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Config struct {
	Server Server `yaml:"server"`
}

func Parse() (*Config, error) {
	const op = "config.Parse()"

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("op: %s, err: %w", op, err)
	}

	configData, err := os.ReadFile(filepath.Join(cwd, "..", "..", "configs", "config.yaml"))
	if err != nil {
		return nil, fmt.Errorf("op: %s, err: %w", op, err)
	}

	var cfg Config
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		return nil, fmt.Errorf("op: %s, err: %w", op, err)
	}

	return &cfg, nil
}
