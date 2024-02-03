package targzip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Extract unzips content of the prodived tar file to dest.
func Extract(reader io.Reader, dest string) error {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("extract: gzip reader: %w", err)
	}

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("extract: tar next: %w", err)
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		if strings.HasPrefix(header.Name, "._") {
			continue
		}

		fname := filepath.Join(dest, header.Name)

		outFile, err := os.Create(fname)
		if err != nil {
			return fmt.Errorf("extract: creating file %s: %w", fname, err)
		}

		if _, err := io.Copy(outFile, tarReader); err != nil {
			return fmt.Errorf("extract: copy file %s: %w", fname, err)
		}

		outFile.Close()
	}

	return nil
}
