package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Project     string   `yaml:"project"`
	Type        string   `yaml:"type"`
	Sources     []string `yaml:"sources"`
	IncludeDirs []string `yaml:"include_dirs"`
	Libraries   []string `yaml:"libraries"`
	Compiler    string   `yaml:"compiler"`
	CFlags      string   `yaml:"cflags"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
