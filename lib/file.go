package lib

import (
	"os"
	"testing"
)

func GetFile(t *testing.T, path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = file.Close() })
	return file
}
