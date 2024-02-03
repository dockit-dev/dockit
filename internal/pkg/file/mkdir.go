package file

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const defaultPerm = 0755

var errBlankPath = errors.New("path is blank")

// Mkdir creates a folder in the current user root folder.
func Mkdir(path string) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("mkdir: %w", errBlankPath)
	}

	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("mkdir: retrieving current user: %w", err)
	}

	dirPath := filepath.Join(currentUser.HomeDir, path)

	err = os.MkdirAll(dirPath, defaultPerm)
	if err != nil {
		return "", fmt.Errorf("mkdir: creating folder: %w", err)
	}

	return dirPath, nil
}
