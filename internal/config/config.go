package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	RootDir        = ".dockit"
	ConfigFilename = "config.json"
)

type Config struct {
	IP             string `json:"ip"`
	CACertPath     string `json:"ca_cert_path"`
	ClientCertPath string `json:"client_cert_path"`
	ClientKeyPath  string `json:"client_key_path"`
}

func Read(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read: reading file: %w", err)
	}

	config := Config{}

	if err = json.Unmarshal(content, &config); err != nil {
		return Config{}, fmt.Errorf("read: unmarshalling content: %w", err)
	}

	return config, nil
}

func Write(cfg Config, path string) error {
	content, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("write: marshaling: %w", err)
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

func Current() (Config, error) {
	currentUser, err := user.Current()
	if err != nil {
		return Config{}, fmt.Errorf("current: retrieving current user: %w", err)
	}

	path := filepath.Join(currentUser.HomeDir, RootDir, ConfigFilename)

	return Read(path)
}
