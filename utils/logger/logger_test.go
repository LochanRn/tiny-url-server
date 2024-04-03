package logger

import (
	"os"
	"testing"
)

func TestIsFileExists(t *testing.T) {
	tempFile := "testfile.txt"
	f, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer func() {
		f.Close()
		os.Remove(tempFile)
	}()

	exists := IsFileExists(tempFile)
	if !exists {
		t.Errorf("Expected file to exist, but it doesn't.")
	}

	nonExistentFile := "nonexistentfile.txt"
	exists = IsFileExists(nonExistentFile)
	if exists {
		t.Errorf("Expected file not to exist, but it does.")
	}
}
