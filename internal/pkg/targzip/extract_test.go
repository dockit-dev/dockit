package targzip_test

import (
	"dockit/internal/pkg/targzip"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {
	is := require.New(t)

	tdata := []struct {
		name    string
		zipPath string
		dest    string
		files   []string
		err     string
	}{
		{
			name: "when zip path is blank",
			err:  "extract: zip path is blank",
		},
		{
			name:    "when dest is blank",
			zipPath: "./fixtures/test_2_files.tar.gz",
			err:     "extract: dest is blank",
		},
		{
			name:    "happy path",
			zipPath: "./fixtures/test_2_files.tar.gz",
			dest:    t.TempDir(),
			files:   []string{"file1", "file2"},
		},
	}

	for _, td := range tdata {
		tt := td

		t.Run(tt.name, func(t *testing.T) {
			err := targzip.Extract(td.zipPath, td.dest)
			if err != nil || len(td.err) != 0 {
				is.Error(err)
				is.ErrorContains(err, td.err)
			}

			if err != nil {
				return
			}

			entries, err := os.ReadDir(td.dest)
			is.NoError(err)

			var actualFiles []string
			for _, entry := range entries {
				actualFiles = append(actualFiles, filepath.Base(entry.Name()))
			}
			sort.Strings(actualFiles)

			is.Equal(td.files, actualFiles)
		})
	}
}
