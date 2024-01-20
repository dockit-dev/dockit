package configure

import (
	"dockit/internal/config"
	"dockit/internal/pkg/file"
	"dockit/internal/pkg/homedir"
	"dockit/internal/pkg/targzip"
	"fmt"
	"os"
	"path/filepath"
)

const (
	caCertFileName     = "ca_cert.pem"
	clientCertFileName = "client_cert.pem"
	clientKeyFileName  = "client_key.pem"
)

func Run(configPath string) error {
	cfgArchive, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("opening config file %s: %w", configPath, err)
	}

	// Create a directory for dockit certificates and config: /username/.dockit
	rootDirPath, err := homedir.Mkdir(config.RootDir)
	if err != nil {
		return fmt.Errorf("creating dockit root folder: %w", err)
	}

	// Unarchive certficates and config
	if err := targzip.Extract(cfgArchive, rootDirPath); err != nil {
		return fmt.Errorf("extracting config archive: %w", err)
	}

	// Read the config.json to get IP address of the dockit instance
	cfgPath := filepath.Join(rootDirPath, config.ConfigFilename)

	cfg, err := config.Read(cfgPath)
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	// Create a folder to store dockit instance certificates
	ipPath := filepath.Join(config.RootDir, cfg.IP)

	ipDirPath, err := homedir.Mkdir(ipPath)
	if err != nil {
		return fmt.Errorf("creating ip folder: %w", err)
	}

	// Move certificates and config to the IP folder
	if err = file.MoveAll(rootDirPath, ipDirPath); err != nil {
		return fmt.Errorf("moving files from %s to %s: %w", rootDirPath, ipDirPath, err)
	}

	// Create a config with currently used dockit instance certificates and IP address
	newCfg := config.Config{
		IP:             cfg.IP,
		CACertPath:     filepath.Join(ipDirPath, caCertFileName),
		ClientCertPath: filepath.Join(ipDirPath, clientCertFileName),
		ClientKeyPath:  filepath.Join(ipDirPath, clientKeyFileName),
	}

	if err := config.Write(newCfg, cfgPath); err != nil {
		return fmt.Errorf("writting config to %s: %w", cfgPath, err)
	}

	return nil
}
