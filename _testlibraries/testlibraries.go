package _testlibraries

import (
	"os"
	"path/filepath"
	"runtime"
)

func LoadMessage() (*string, error) {
	_, file, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(file), "testdata", "log_example.txt")
	messageByte, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	message := string(messageByte)
	return &message, nil
}
