package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/amitsol123/html_downloader/pkg/downloader"
)

const (
	outputDirectory   = "downloaded_html/"
	relativeURLsFile  = "../../ListOfAsciiSiteUrls.txt"
	baseURLsFile      = "ListOfAsciiSiteUrls.txt"
	defaultMaxWorkers = 5 // Maximum number of workers
)

func main() {
	maxWorkers := getDefaultMaxWorkers()
	urlList, err := readURLsFromFile(getURLsFilePath())
	if err != nil {
		fmt.Println("Error reading URLs:", err)
		return
	}

	err = createOutputDirectory()
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	workerChannel := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for _, url := range urlList {
		workerChannel <- struct{}{}
		wg.Add(1)

		go func(url string) {
			defer func() {
				<-workerChannel
				wg.Done()
			}()

			err := downloader.DownloadHTML(url, outputDirectory)
			if err != nil {
				fmt.Printf("Error downloading from %s: %s\n", url, err)
				return
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("All downloads completed!")
}

func getDefaultMaxWorkers() int {
	if len(os.Args) > 1 {
		workers, err := strconv.Atoi(os.Args[1])
		if err == nil && workers > 0 {
			return workers
		}
		fmt.Println("Invalid number of workers. Using default.")
	}
	return defaultMaxWorkers
}

func readURLsFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return []string{}, nil // Return an empty string array if content is empty
	}

	return strings.Split(string(content), "\n"), nil
}

func getURLsFilePath() string {
	if _, err := os.Stat(relativeURLsFile); err == nil {
		return relativeURLsFile // Use relative path if file exists
	}
	// File doesn't exist using the relative path, try using an absolute path

	if _, err := os.Stat(baseURLsFile); err == nil {
		return baseURLsFile // Use absolute path if file exists
	}
	// If file not found in both relative and absolute paths, handle it accordingly
	return "" // Or return an error, log a message, etc.
}

func createOutputDirectory() error {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		return os.Mkdir(outputDirectory, 0755)
	} else if err != nil {
		return err
	}
	return nil
}
