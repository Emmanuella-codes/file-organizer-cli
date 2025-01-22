package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestConfigLoadAndSave(t *testing.T) {
	// create a temp config file
	tempDir, err := os.MkdirTemp("", "file-organizer-config-test")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "config.json")

	// create a test config
	testConfig := &Config{
		CustomCategories: map[string][]string{
			"Code": {".go", ".py", ".js"},
			"Data": {".csv", ".json"},
		},
		ExcludePatterns: []string{".git/*", "node_modules/*"},
	}

	// save config
	if err := testConfig.Save(configPath); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// load config
	loadedConfig, err := LoadConfig(configPath)
	if err != nil {
			t.Fatalf("Failed to load config: %v", err)
	}

	// compare configs
	if !reflect.DeepEqual(testConfig, loadedConfig) {
			t.Errorf("Loaded config doesn't match saved config")
	}
}

func TestConfigLoadNonexistent(t *testing.T) {
	// Try to load non-existent config
	config, err := LoadConfig("nonexistent.json")
	if err != nil {
			t.Fatalf("Expected no error for non-existent config, got: %v", err)
	}

	// Verify default values
	if len(config.CustomCategories) != 0 {
			t.Error("Expected empty custom categories for default config")
	}
	if len(config.ExcludePatterns) != 0 {
			t.Error("Expected empty exclude patterns for default config")
	}
}
