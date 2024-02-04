package config_test

import (
	"dockit/internal/config"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	is := require.New(t)

	correctConfigPath := createTempFile(
		t,
		`{
			"ip": "127.0.0.1",
			"ca_cert_path": "ca_cert_path",
			"client_cert_path": "client_cert_path",
			"client_key_path": "client_key_path"
		}`,
	)

	malformedConfigPath := createTempFile(t, `{ip": "127.0.0.1"}`)

	tdata := []struct {
		name   string
		path   string
		expCfg config.Config
		err    string
	}{
		{
			name: "when filename is blank",
			err:  "read: reading file: open : no such file or directory",
		},
		{
			name: "when file is malformed",
			path: malformedConfigPath,
			err:  "read: unmarshalling content: invalid character 'i' looking for beginning of object key string",
		},
		{
			name: "happy path",
			path: correctConfigPath,
			expCfg: config.Config{
				IP:             "127.0.0.1",
				CACertPath:     "ca_cert_path",
				ClientCertPath: "client_cert_path",
				ClientKeyPath:  "client_key_path",
			},
		},
	}

	for _, td := range tdata {
		tt := td

		t.Run(tt.name, func(t *testing.T) {
			cfg, err := config.Read(td.path)
			if err != nil || len(td.err) != 0 {
				is.Error(err)
				is.ErrorContains(err, td.err)
			}

			if err != nil {
				return
			}

			is.Equal(cfg, td.expCfg)
		})
	}
}

func TestCurrent(t *testing.T) {
	is := require.New(t)

	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("retrieving current user: %v", err)
	}

	err = os.Mkdir(filepath.Join(currentUser.HomeDir, ".dockit"), 0755)
	if err != nil {
		t.Fatalf("creating config file: %v", err)
	}

	cfgFile, err := os.Create(filepath.Join(currentUser.HomeDir, ".dockit", "config.json"))
	if err != nil {
		t.Fatalf("creating config file: %v", err)
	}

	defer func() {
		if err := os.Remove(cfgFile.Name()); err != nil {
			log.Printf("removing file: %v", err)
		}
	}()

	_, err = cfgFile.WriteString(`{"ip": "127.0.0.1"}`)
	if err != nil {
		t.Fatalf("writing to config file: %v", err)
	}

	if err = cfgFile.Close(); err != nil {
		t.Fatalf("closing config file: %v", err)
	}

	cfg, err := config.Current()
	is.NoError(err)
	is.Equal(cfg, config.Config{IP: "127.0.0.1"})
}

func TestWriteCurrent(t *testing.T) {
	is := require.New(t)

	cfg := config.Config{
		IP:             "127.0.0.1",
		CACertPath:     "ca_cert_path",
		ClientCertPath: "client_cert_path",
		ClientKeyPath:  "client_key_path",
	}

	err := config.WriteCurrent(cfg)
	is.NoError(err)

	currentUser, err := user.Current()
	is.NoError(err)

	content, err := os.ReadFile(filepath.Join(currentUser.HomeDir, ".dockit", "config.json"))
	is.NoError(err)
	is.Equal(
		string(content),
		"{\"ip\":\"127.0.0.1\",\"ca_cert_path\":\"ca_cert_path\",\"client_cert_path\":\"client_cert_path\",\"client_key_path\":\"client_key_path\"}",
	)
}

func createTempFile(t *testing.T, content string) string {
	file, err := os.CreateTemp(t.TempDir(), "")
	if err != nil {
		t.Fatalf("creating a temp file: %v", err)
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		t.Fatalf("writing to the temp file: %v", err)
	}

	if err = file.Close(); err != nil {
		t.Fatalf("closing the temp file: %v", err)
	}

	return file.Name()
}
