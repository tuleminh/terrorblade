package osenv

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadFile loads environment variables for .env file.
func LoadFile(filePath string) error {
	err := godotenv.Load(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	return nil
}

func String(key string) string {
	return os.Getenv(key)
}
