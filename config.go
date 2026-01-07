package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Device   string                 `yaml:"device"`
	Deadzone int32                  `yaml:"deadzone"`
	Buttons  map[string]interface{} `yaml:"buttons"`
	Axes     map[string]AxisMap     `yaml:"axes"`
}

type AxisMap struct {
	Negative string `yaml:"negative"`
	Positive string `yaml:"positive"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

