package file_test

import (
	"crypto/subtle"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockit-dev/dockit/internal/pkg/file"

	"github.com/stretchr/testify/require"
)

func TestMoveAll(t *testing.T) {
	is := require.New(t)

	tdata := []struct {
		name    string
		srcDir  string
		destDir string
		err     string
	}{
		{
			name:    "when src is blank",
			destDir: t.TempDir(),
			err:     "move all: source is blank",
		},
		{
			name:   "when dest is blank",
			srcDir: t.TempDir(),
			err:    "move all: destination is blank",
		},
		{
			name:    "happy path",
			srcDir:  t.TempDir(),
			destDir: t.TempDir(),
		},
	}

	for _, td := range tdata {
		tt := td

		t.Run(tt.name, func(t *testing.T) {
			var srcFiles []string

			if len(td.srcDir) != 0 {
				srcFiles = []string{
					createTempFile(t, td.srcDir),
					createTempFile(t, td.srcDir),
					createTempFile(t, td.srcDir),
				}
			}

			err := file.MoveAll(td.srcDir, td.destDir)
			if err != nil || len(td.err) != 0 {
				is.ErrorContains(err, td.err)
			}

			if err != nil {
				return
			}

			for _, srcFile := range srcFiles {
				destFile := filepath.Join(td.destDir, filepath.Base(srcFile))
				fileIsMoved(t, srcFile, destFile)
			}
		})
	}
}

func TestMove(t *testing.T) {
	is := require.New(t)

	tdata := []struct {
		name    string
		srcDir  string
		destDir string
		err     string
	}{
		{
			name:    "when src is blank",
			destDir: t.TempDir(),
			err:     "move: source is blank",
		},
		{
			name:   "when dest is blank",
			srcDir: t.TempDir(),
			err:    "move: destination is blank",
		},
		{
			name:    "happy path",
			srcDir:  t.TempDir(),
			destDir: t.TempDir(),
		},
	}

	for _, td := range tdata {
		tt := td

		t.Run(tt.name, func(t *testing.T) {
			var srcFile, destFile string

			if len(td.srcDir) != 0 {
				srcFile = createTempFile(t, td.srcDir)
			}

			if len(td.destDir) != 0 {
				destFile = filepath.Join(td.destDir, filepath.Base(srcFile))
			}

			err := file.Move(srcFile, destFile)
			if err != nil || len(td.err) != 0 {
				is.ErrorContains(err, td.err)
			}

			if err != nil {
				return
			}

			fileIsMoved(t, srcFile, destFile)
		})
	}
}

func createTempFile(t *testing.T, dir string) string {
	file, err := os.CreateTemp(dir, "")
	if err != nil {
		t.Fatalf("creating a temp file: %v", err)
	}

	_, err = file.Write([]byte(file.Name()))
	if err != nil {
		t.Fatalf("writing to the temp file: %v", err)
	}

	if err = file.Close(); err != nil {
		t.Fatalf("closing the temp file: %v", err)
	}

	return file.Name()
}

func fileIsMoved(t *testing.T, src, dest string) {
	content, err := os.ReadFile(dest)
	if err != nil {
		t.Errorf("unexpected error reading file: %v", err)
		return
	}

	if subtle.ConstantTimeCompare(content, []byte(src)) != 1 {
		t.Errorf("file content is wrong: expected=%s, actual=%s", src, content)
	}
}
