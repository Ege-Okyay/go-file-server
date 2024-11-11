package services

import (
	"io"
	"os"
	"path/filepath"

	"github.com/Ege-Okyay/go-file-server/config"
	"github.com/google/uuid"
)

func SaveFile(file io.Reader, extension string) (string, error) {
	filename := uuid.New().String() + extension
	path := filepath.Join(config.UploadPath, filename)

	outFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		return "", err
	}

	return filename, nil
}
