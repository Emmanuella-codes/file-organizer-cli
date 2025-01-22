package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	CustomCategories map[string][]string `json:"customCategories"`
	ExcludePatterns  []string            `json:"excludePatterns"`
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = filepath.Join(os.Getenv("HOME"), ".file-organizer.json")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				CustomCategories: make(map[string][]string),
				ExcludePatterns:  []string{},
			}, nil
		}
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) Save(configPath string) error {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
