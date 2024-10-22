package internal

import (
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
	err := os.MkdirAll(appFolderPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create app folder: %v", err)
	}

	// Change the working directory to the temporary directory
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Call DeleteApp
	DeleteApp(appName)

	// Verify the app folder has been deleted
	if _, err = os.Stat(appFolderPath); !os.IsNotExist(err) {
		t.Fatalf("App folder was not deleted")
	}
}

func TestDeleteApp_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	appName := "test-app"
	appFolderPath := filepath.Join(tmpDir, appName)

	// Change the working directory to the temporary directory
	err := os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Call DeleteApp and expect a panic
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic but did not occur")
		}
	}()
	DeleteApp(appName)

	// Verify the app folder does not exist
	if _, err = os.Stat(appFolderPath); !os.IsNotExist(err) {
		t.Fatalf("App folder should not exist")
	}
}

func TestDeleteApp_ErrorDeleting(t *testing.T) {
	tmpDir := t.TempDir()
	appName := "test-app"
	appFolderPath := filepath.Join(tmpDir, appName)

	// Create a temporary app folder
	err := os.Mkdir(appFolderPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create app folder: %v", err)
	}

	// Change the working directory to the temporary directory
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Make the app folder read-only to simulate a deletion error
	err = os.Chmod(appFolderPath, 0444)
	if err != nil {
		t.Fatalf("Failed to change folder permissions: %v", err)
	}

	// Call DeleteApp and expect a panic
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic but did not occur")
		}
	}()
	DeleteApp(appName)

	// Clean up by making the app folder writable again
	err = os.Chmod(appFolderPath, 0755)
	if err != nil {
		t.Fatalf("Failed to change folder permissions: %v", err)
	}
}
