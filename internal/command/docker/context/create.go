package context

import (
	"dockit/internal/config"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Create creates a new docker context and makes it a current one.
func Create() error {
	cfg, err := config.Current()
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	var (
		name        = fmt.Sprintf("dockit_%s", strings.ReplaceAll(cfg.IP, ".", ""))
		description = fmt.Sprintf("Dockit instance (%s)", cfg.IP)
		docker      = fmt.Sprintf(
			"host=tcp://%s:2376,ca=%s,cert=%s,key=%s",
			cfg.IP, cfg.CACertPath, cfg.ClientCertPath, cfg.ClientKeyPath,
		)
	)

	// Create a new docker context
	createCmd := exec.Command(
		"docker",
		"context",
		"create",
		name,
		"--description",
		description,
		"--docker",
		docker,
	)

	if err := runCmd(createCmd); err != nil {
		return err
	}

	// Use the newly created docker context
	useCmd := exec.Command(
		"docker",
		"context",
		"use",
		name,
	)

	if err := runCmd(useCmd); err != nil {
		return err
	}

	return nil
}

func runCmd(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
