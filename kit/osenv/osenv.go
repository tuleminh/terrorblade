package osenv

import (
	"os"
	"path/filepath"
	"strconv"

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

// String returns cfg value as a string.
func String(key string, defVal ...string) string {
	var v string
	if len(defVal) > 0 {
		v = defVal[0]
	}
	s := GetEnv(key)
	if s != "" {
		v = s
	}
	return v
}

// Bool parses and returns cfg as boolean.
func Bool(key string, defVal ...bool) bool {
	var v bool
	if len(defVal) > 0 {
		v = defVal[0]
	}
	b, err := strconv.ParseBool(GetEnv(key))
	if err == nil {
		v = b
	}
	return v
}

// Int parses and returns cfg as integer.
func Int(key string, defVal ...int) int {
	var v int
	if len(defVal) > 0 {
		v = defVal[0]
	}
	n, err := strconv.Atoi(GetEnv(key))
	if err == nil {
		v = n
	}
	return v
}

// GetEnv returns value of environment variables as string.
func GetEnv(key string) string {
	return os.Getenv(key)
}
