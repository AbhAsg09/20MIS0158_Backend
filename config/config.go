package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

// Config represents the structure of the config.yaml file
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`

	YouTube struct {
		APIKeys              []string `yaml:"api_keys"`
		SearchQuery          string   `yaml:"search_query"`
		FetchIntervalSeconds int      `yaml:"fetch_interval_seconds"`
	} `yaml:"youtube"`
}

// LoadConfig loads configuration from the config.yaml file
func LoadConfig(path string) (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// GetFetchInterval returns the fetch interval as a time.Duration
func (cfg *Config) GetFetchInterval() time.Duration {
	return time.Duration(cfg.YouTube.FetchIntervalSeconds) * time.Second
}
