package utils

import (
	"errors"
	"os"
)

// DoesFileExist returns true if a file exists or false if not.
func DoesFileExist(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	defer file.Close()

	return true, nil
}

// DoesFolderExist returns true if a file exists or false if not.
func DoesFolderExist(folderPath string) (bool, error) {
	_, err := os.Stat(folderPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
