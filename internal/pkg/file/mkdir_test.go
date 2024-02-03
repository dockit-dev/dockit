package file_test

import (
	"dockit/internal/pkg/file"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMkdir(t *testing.T) {
	is := require.New(t)

	tdata := []struct {
		name string
		path string
		err  string
	}{
		{
			name: "when path is blank",
			err:  "mkdir: path is blank",
		},
		{
			name: "happy path",
			path: ".dockit_test",
		},
	}

	for _, td := range tdata {
		tt := td

		t.Run(tt.name, func(t *testing.T) {
			path, err := file.Mkdir(td.path)
			if err != nil || len(td.err) != 0 {
				is.Error(err)
				is.ErrorContains(err, td.err)
			}

			if err != nil {
				return
			}

			dirStat, err := os.Stat(path)
			is.NoError(err)
			is.True(dirStat.IsDir())
		})
	}
}
