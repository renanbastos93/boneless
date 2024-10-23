package internal

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteApp(t *testing.T) {
	tmpDir := t.TempDir()
	appName := "test-app"
	appFolderPath := filepath.Join(tmpDir, "internal", appName)

	t.Log("tmpDir: ", tmpDir)

	// Create a temporary app folder
	if err := os.MkdirAll(appFolderPath, 0755); err != nil {
		t.Fatalf("Failed to create app folder: %v", err)
	}

	// Change the working directory to the temporary directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Call DeleteApp
	DeleteApp(appName)

	// Verify the app folder has been deleted
	if _, err := os.Stat(appFolderPath); !os.IsNotExist(err) {
		t.Fatalf("App folder was not deleted")
	}
}

func TestDeleteApp_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	appName := "test-app"

	// Change the working directory to the temporary directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Save the current stdout
	originalStdout := os.Stdout

	// Create a buffer to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	DeleteApp(appName)

	// Restore the original stdout and close the writer
	w.Close()
	os.Stdout = originalStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != "App folder not found\n" {
		t.Fatalf("Unexpected output: %s", buf.String())
	}
}

func TestDeleteApp_ErrorDeleting(t *testing.T) {
	tmpDir := t.TempDir()

	// Change the working directory to the temporary directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Call DeleteApp and expect a panic
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic but did not occur")
		}
	}()
	DeleteApp("\x00")
}
