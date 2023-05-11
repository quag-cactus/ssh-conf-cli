package utils

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func DefineDefaultConfigPath() string {

	var configPath string
	switch runtime.GOOS {
	case "windows": // windows
		configPath = filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "config")
	case "darwin": // mac
		configPath = filepath.Join(os.Getenv("HOME"), ".ssh", "config")
	case "linux": // linux
		configPath = filepath.Join(os.Getenv("HOME"), ".ssh", "config")
	}

	return configPath

}

func CreateBackupConfigFile(filePath string) (string, error) {

	inputFs, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	backupFilePath := filePath + ".old"
	bkupFs, err := os.Create(backupFilePath)
	if err != nil {
		return "", err
	}
	defer bkupFs.Close()
	defer inputFs.Close()

	io.Copy(bkupFs, inputFs)

	return backupFilePath, nil

}
