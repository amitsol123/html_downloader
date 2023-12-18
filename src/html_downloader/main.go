package main

import (
	"fmt"
	filehandler "github.com/amitsol123/html_downloader/src/utils"
	"sync"

	"github.com/amitsol123/html_downloader/pkg/downloader"
	"github.com/amitsol123/html_downloader/src/config"
)

func main() {
	maxWorkers := config.GetDefaultMaxWorkers()
	urlFilePath := filehandler.GetURLsFilePath(config.GetRelativeURLsFile(), config.GetBaseURLsFile())
	urlList, err := filehandler.ReadURLsFromFile(urlFilePath)
	if err != nil {
		handleError("Error reading URLs", err)
		return
	}

	err = filehandler.CreateOutputDirectory(config.GetOutputDirectory())
	if err != nil {
		handleError("Error creating output directory", err)
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

			err := downloader.DownloadHTMLContent(url, config.GetOutputDirectory())
			if err != nil {
				fmt.Printf("Error downloading from %s: %s\n", url, err)
				return
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("All downloads completed!")
}

func handleError(msg string, err error) {
	fmt.Println(msg, ":", err)
}
