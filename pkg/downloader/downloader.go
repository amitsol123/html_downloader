package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// DownloadHTMLContent DownloadHTML downloads HTML content from a given URL and saves it to a file.
func DownloadHTMLContent(url, outputDirectory string) error {
	// Perform GET request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer closeBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	filename := ExtractFilenameFromURL(url)
	filePath := outputDirectory + filename + ".html"
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer closeFile(file)

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy response body to file: %w", err)
	}

	return nil
}

// ExtractFilenameFromURL Function to extract the filename from the URL
func ExtractFilenameFromURL(url string) string {
	parts := strings.Split(url, "/")
	lastSegment := parts[len(parts)-1]

	filenameWithQuery := strings.Split(lastSegment, "?")[0]

	return filenameWithQuery
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		fmt.Println("Error closing body:", err)
	}
}

// closeFile safely closes the file.
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}
