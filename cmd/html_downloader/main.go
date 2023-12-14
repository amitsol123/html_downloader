package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/amitsol/html_downloader/pkg/downloader"
)

const (
	outputDirectory = "downloaded_html/"
	urlsFile        = "urls.txt"
)

func main() {
	urlList, err := readURLsFromFile(urlsFile)
	if err != nil {
		fmt.Println("Error reading URLs:", err)
		return
	}

	os.Mkdir(outputDirectory, 0755)

	var wg sync.WaitGroup
	maxWorkers := 5 // adjust this value as needed

	for _, url := range urlList {
		wg.Add(1)
		go downloader.DownloadHTML(url, outputDirectory, &wg)
	}

	wg.Wait()
	fmt.Println("All downloads completed!")
}

func readURLsFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	urlList := strings.Split(string(content), "\n")
	return urlList, nil
}
