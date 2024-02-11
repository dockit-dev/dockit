package configure_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockit-dev/dockit/internal/command/configure"
	"github.com/dockit-dev/dockit/internal/config"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	is := require.New(t)

	homeDir, err := os.UserHomeDir()
	is.NoError(err)

	dockitDir := filepath.Join(homeDir, config.Dir)

	t.Cleanup(func() {
		if err := os.RemoveAll(dockitDir); err != nil {
			log.Printf("removing dockit dir: %v", err)
		}
	})

	is.NoError(
		configure.Run("./fixtures/dockit-config.tar.gz"),
	)

	currentConfig, err := os.ReadFile(filepath.Join(dockitDir, "config.json"))
	is.NoError(err)
	is.Equal(
		"{\"ip\":\"49.13.13.232\",\"ca_cert_path\":\"/root/.dockit/49.13.13.232/ca_cert.pem\",\"client_cert_path\":\"/root/.dockit/49.13.13.232/client_cert.pem\",\"client_key_path\":\"/root/.dockit/49.13.13.232/client_key.pem\"}",
		string(currentConfig),
	)

	entries, err := os.ReadDir(filepath.Join(dockitDir, "49.13.13.232"))
	is.NoError(err)

	var files []string
	for _, entry := range entries {
		files = append(files, filepath.Base(entry.Name()))
	}

	is.Equal([]string([]string{"ca_cert.pem", "client_cert.pem", "client_key.pem", "config.json"}), files)

	ipConfig, err := os.ReadFile(filepath.Join(dockitDir, "49.13.13.232", "config.json"))
	is.NoError(err)
	is.Equal(
		"{\"ip\": \"49.13.13.232\"}\n",
		string(ipConfig),
	)

	for _, pem := range []string{"ca_cert.pem", "client_cert.pem", "client_key.pem"} {
		cert, err := os.ReadFile(filepath.Join(dockitDir, "49.13.13.232", pem))
		is.NoError(err)
		is.NotEmpty(cert)
	}
}
