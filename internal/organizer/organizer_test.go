package organizer

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestFiles(t *testing.T, baseDir string) {
	// create test files with different extensions
	files := map[string]string{
		"document.pdf": "Documents",
		"image.jpg":    "Images",
    "video.mp4":    "Videos",
    "music.mp3":    "Audio",
    "archive.zip":  "Archives",
    "random.xyz":   "Others",
	}

	for filename, _ := range files {
		path := filepath.Join(baseDir, filename)
		if err := os.WriteFile(path, []byte("test content"), 0644); err != nil {
				t.Fatalf("Failed to create test file %s: %v", filename, err)
		}
	}
}

func TestOrganizer(t *testing.T) {
	// create temp directory
	testDir, err := os.MkdirTemp("", "file-organizer-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	// create test files
	createTestFiles(t, testDir)

	// create and run organizer
	org := New(testDir)
	if err := org.Organize(); err != nil {
		t.Fatalf("Failed to organize files: %v", err)
	}

	// verify that files were moved to correct categories
	expectedLocations := map[string]string{
		"document.pdf": "Documents",
		"image.jpg":    "Images",
		"video.mp4":    "Videos",
		"music.mp3":    "Audio",
		"archive.zip":  "Archives",
		"random.xyz":   "Others",
	}

	for filename, category := range expectedLocations {
		expectedPath := filepath.Join(testDir, category, filename)
		if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
			t.Errorf("Expected file %s to be in %s category", filename, category)
		}
	}
}

func TestOrganizerWithCustomCategories(t *testing.T) {
	testDir, err := os.MkdirTemp("", "file-organizer-custom-test")
    if err != nil {
        t.Fatalf("Failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(testDir)

    // create a test file
    testFile := filepath.Join(testDir, "test.xyz")
    if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }

    // create organizer with custom categories
    org := New(testDir)
    org.Categories = []FileCategory{
        {
            Name:       "CustomCat",
            Extensions: []string{".xyz"},
        },
    }

    // run organizer
    if err := org.Organize(); err != nil {
        t.Fatalf("Failed to organize files: %v", err)
    }

    // verify file was moved to custom category
    expectedPath := filepath.Join(testDir, "CustomCat", "test.xyz")
    if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
        t.Errorf("Expected file test.xyz to be in CustomCat category")
    }
}
