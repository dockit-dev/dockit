package homedir

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const defaultPerm = 0755

func Mkdir(path string) (string, error) {
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
