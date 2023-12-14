package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// DownloadHTML downloads HTML content from a given URL and saves it to a file.
func DownloadHTML(url, outputDirectory string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %s\n", url, err)
		return
	}

	// Extract the filename from the URL (you might need a more robust method)
	filename := extractFilenameFromURL(url)

	filePath := outputDirectory + filename + ".html"
	err = ioutil.WriteFile(filePath, html, 0644)
	if err != nil {
		fmt.Printf("Error writing file for %s: %s\n", url, err)
		return
	}

	fmt.Printf("Downloaded %s\n", url)
}

// Function to extract the filename from the URL
func extractFilenameFromURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
