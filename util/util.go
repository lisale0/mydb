package util

import (
	"os"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func CreateFile(filePath string) (*os.File, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func OpenFile(filePath string) (*os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ParseBool() (bool, error) {
	var err error
	return true, err
}