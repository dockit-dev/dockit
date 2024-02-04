package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	// Dir is the dockit directory name which stores all the necessary configurations.
	Dir = ".dockit"

	// filename is the file name of the dockit config.
	filename = "config.json"
)

// Config defines the dockit config.
type Config struct {
	IP             string `json:"ip"`
	CACertPath     string `json:"ca_cert_path"`
	ClientCertPath string `json:"client_cert_path"`
	ClientKeyPath  string `json:"client_key_path"`
}

// Read unmarshals the content of the given file to Config.
func Read(filename string) (Config, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("read: reading file: %w", err)
	}

	config := Config{}

	if err = json.Unmarshal(content, &config); err != nil {
		return Config{}, fmt.Errorf("read: unmarshalling content: %w", err)
	}

	return config, nil
}

// Write saves the provided config as the current dockit configuration.
func WriteCurrent(cfg Config) error {
	content, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("write: marshaling: %w", err)
	}

	path, err := fullCurrentPath()
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("write: creating file: %w", err)
	}

	_, err = outputFile.Write(content)
	if err != nil {
		return fmt.Errorf("write: writting: %w", err)
	}

	return nil
}

// Current returns the current configiration.
func Current() (Config, error) {
	path, err := fullCurrentPath()
	if err != nil {
		return Config{}, fmt.Errorf("current: %w", err)
	}

	cfg, err := Read(path)
	if err != nil {
		return Config{}, fmt.Errorf("current: reading config: %w", err)
	}

	return cfg, nil
}

func fullCurrentPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("retrieving current user: %w", err)
	}

	return filepath.Join(currentUser.HomeDir, Dir, filename), nil
}
