package downloader

import (
	"os"
	"testing"
)

func TestDownloadHTML(t *testing.T) {
	err := os.MkdirAll("test_output", os.ModePerm)
	if err != nil {
		t.Fatalf("Error creating test_output directory: %s", err)
	}
	defer os.RemoveAll("test_output")

	tests := []struct {
		name            string
		url             string
		outputDirectory string
		wantError       bool
	}{
		{
			name:            "Successful download",
			url:             "https://www.example.com",
			outputDirectory: "test_output/",
			wantError:       false,
		},
		{
			name:            "Invalid URL",
			url:             "invalid-url", // Provide an invalid URL
			outputDirectory: "test_output/",
			wantError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DownloadHTML(tt.url, tt.outputDirectory)

			if tt.wantError && err == nil {
				t.Errorf("Expected error, but got none")
			}

			if !tt.wantError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestMain runs cleanup after all tests are done
func TestMain(m *testing.M) {
	code := m.Run()

	// Clean up test files or directories if needed
	_ = os.RemoveAll("test_output")

	os.Exit(code)
}

func TestExtractFilenameFromURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "Basic URL with filename",
			url:      "https://www.example.com/some/path/file.html",
			expected: "file.html",
		},
		{
			name:     "URL with trailing slash",
			url:      "https://www.example.com/some/path/",
			expected: "", // Modify this based on your expected behavior
		},
		{
			name:     "URL with query parameters",
			url:      "https://www.example.com/some/path/file.html?param=value",
			expected: "file.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractFilenameFromURL(tt.url)

			if result != tt.expected {
				t.Errorf("Expected: %s, Got: %s", tt.expected, result)
			}
		})
	}
}
