package targzip

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	errBlankZipPath = errors.New("zip path is blank")
	errBlankDest    = errors.New("dest is blank")
)

// Extract unzips the provided archive to dest.
func Extract(zipPath, dest string) error {
	if len(zipPath) == 0 {
		return fmt.Errorf("extract: %w", errBlankZipPath)
	}

	if len(dest) == 0 {
		return fmt.Errorf("extract: %w", errBlankDest)
	}

	zipFile, err := os.Open(zipPath)
	if err != nil {
		return fmt.Errorf("extract: opening zip file %s: %w", zipPath, err)
	}

	gzipReader, err := gzip.NewReader(zipFile)
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
