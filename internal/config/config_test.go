package config_test

import (
	"os"
	"testing"

	"github.com/dockit-dev/dockit/internal/config"

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
