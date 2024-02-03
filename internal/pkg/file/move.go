package file

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	errBlankSrc  = errors.New("source is blank")
	errBlankDest = errors.New("destination is blank")
)

// MoveAll moves all files from srcDir to destDir.
func MoveAll(srcDir, destDir string) error {
	if len(srcDir) == 0 {
		return fmt.Errorf("move all: %w", errBlankSrc)
	}

	if len(destDir) == 0 {
		return fmt.Errorf("move all: %w", errBlankDest)
	}

	dir, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("move all: reading src dir %s: %w", srcDir, err)
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		src := filepath.Join(srcDir, file.Name())
		dest := filepath.Join(destDir, file.Name())

		err := Move(src, dest)
		if err != nil {
			return fmt.Errorf("move all: moving file %s to %s: %w", src, dest, err)
		}
	}

	return nil
}

// Move moves the src file to the given dest file.
func Move(src, dest string) error {
	if len(src) == 0 {
		return fmt.Errorf("move: %w", errBlankSrc)
	}

	if len(dest) == 0 {
		return fmt.Errorf("move: %w", errBlankDest)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("move: opening src file %s: %w", src, err)
	}

	destFile, err := os.Create(dest)
	if err != nil {
		srcFile.Close()

		return fmt.Errorf("move: creating dest file %s: %w", dest, err)
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("move: copying file: %w", err)
	}

	if err := srcFile.Close(); err != nil {
		return fmt.Errorf("move: closing src file %s: %w", src, err)
	}

	if err = os.Remove(src); err != nil {
		return fmt.Errorf("move: removing src file %s: %w", src, err)
	}

	return nil
}
