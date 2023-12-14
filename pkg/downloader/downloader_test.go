package downloader_test

import (
	"os"
	"sync"
	"testing"

	"github.com/yourusername/html_downloader/pkg/downloader"
)

func TestDownloadHTML(t *testing.T) {
	url := "https://example.com"
	outputDir := "test_output/"
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go downloader.DownloadHTML(url, outputDir, wg)
	wg.Wait()

	// Verify that the file has been created
	filename := downloader.ExtractFilenameFromURL(url)
	filePath := outputDir + filename + ".html"
	_, err := os.Stat(filePath)
	if err != nil {
		t.Errorf("File not found at path %s", filePath)
	}

	// Clean up test file
	err = os.Remove(filePath)
	if err != nil {
		t.Errorf("Error removing test file: %s", err)
	}
}

func TestExtractFilenameFromURL(t *testing.T) {
	url := "https://example.com/some-page.html"
	expectedFilename := "some-page.html"

	filename := downloader.ExtractFilenameFromURL(url)

	if filename != expectedFilename {
		t.Errorf("Expected filename: %s, Got: %s", expectedFilename, filename)
	}
}
