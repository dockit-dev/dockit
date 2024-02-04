package configure

import (
	"dockit/internal/config"
	"dockit/internal/pkg/file"
	"dockit/internal/pkg/targzip"
	"fmt"
	"path/filepath"
)

const (
	caCertFileName     = "ca_cert.pem"
	clientCertFileName = "client_cert.pem"
	clientKeyFileName  = "client_key.pem"
)

func Run(configPath string) error {
	// Create a directory for dockit certificates and config: /username/.dockit
	dirPath, err := file.Mkdir(config.Dir)
	if err != nil {
		return fmt.Errorf("creating dockit root folder: %w", err)
	}

	// Unarchive certficates and config
	if err := targzip.Extract(configPath, dirPath); err != nil {
		return fmt.Errorf("extracting config archive: %w", err)
	}

	// Read the config.json to get IP address of the dockit instance
	cfg, err := config.Current()
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	// Create a folder to store dockit instance certificates
	ipPath := filepath.Join(config.Dir, cfg.IP)

	ipDirPath, err := file.Mkdir(ipPath)
	if err != nil {
		return fmt.Errorf("creating ip folder: %w", err)
	}

	// Move certificates and config to the IP folder
	if err = file.MoveAll(dirPath, ipDirPath); err != nil {
		return fmt.Errorf("moving files from %s to %s: %w", dirPath, ipDirPath, err)
	}

	// Create a config with currently used dockit instance certificates and IP address
	newCfg := config.Config{
		IP:             cfg.IP,
		CACertPath:     filepath.Join(ipDirPath, caCertFileName),
		ClientCertPath: filepath.Join(ipDirPath, clientCertFileName),
		ClientKeyPath:  filepath.Join(ipDirPath, clientKeyFileName),
	}

	if err := config.WriteCurrent(newCfg); err != nil {
		return fmt.Errorf("writting config: %w", err)
	}

	return nil
}
