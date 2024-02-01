package config

import (
	"chess/pkg/zaplogger"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server ServerConfig     `yaml:"server"`
	Logger zaplogger.Config `yaml:"zaplogger"`
}

func Load(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	cfg := Config{}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	return &cfg, nil
}

type ServerConfig struct {
	Port int `yaml:"port"`
}
