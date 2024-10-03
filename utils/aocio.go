package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func GetInput(t *testing.T, file string) string {
	goldenfile := filepath.Join("testdata", file)
	want, err := os.ReadFile(goldenfile)
	if err != nil {
		t.Fatal("error reading golden file:", err)
	}

	return string(want)
}
