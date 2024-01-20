package root

import (
	"dockit/internal/config"
	"fmt"
	"os"
	"os/exec"
)

func Run(args []string) error {
	cfg, err := config.Current()
	if err != nil {
		return fmt.Errorf("retrieving current config")
	}

	tlsArgs := []string{
		"--tlsverify",
		fmt.Sprintf("--tlscacert=%s", cfg.CACertPath),
		fmt.Sprintf("--tlscert=%s", cfg.ClientCertPath),
		fmt.Sprintf("--tlskey=%s", cfg.ClientKeyPath),
		fmt.Sprintf("-H=%s", cfg.IP),
	}
	tlsArgs = append(tlsArgs, args...)

	cmd := exec.Command("docker", tlsArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
