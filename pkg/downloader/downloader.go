package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// DownloadHTML downloads HTML content from a given URL and saves it to a file.
func DownloadHTML(url, outputDirectory string) error {
	// Perform GET request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(resp.Body)

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create the output file
	filename := ExtractFilenameFromURL(url)
	filePath := outputDirectory + filename + ".html"
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy response body to file: %w", err)
	}

	return nil // No error occurred
}

// ExtractFilenameFromURL Function to extract the filename from the URL
func ExtractFilenameFromURL(url string) string {
	// Split the URL by '/'
	parts := strings.Split(url, "/")

	// Get the last segment, which could represent the filename or directory
	lastSegment := parts[len(parts)-1]

	// Remove query parameters if present in the last segment
	filenameWithQuery := strings.Split(lastSegment, "?")[0]

	// Return the extracted filename or directory
	return filenameWithQuery
}
