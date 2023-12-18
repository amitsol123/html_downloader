package main

//
//import (
//	"fmt"
//	"sync"
//	"testing"
//
//	"github.com/amitsol123/html_downloader/pkg/downloader"
//	"github.com/amitsol123/html_downloader/src/config"
//	filehandler "github.com/amitsol123/html_downloader/src/utils"
//)
//
//func TestMain(t *testing.T) {
//	cfg := config.NewConfig()
//
//	// Create your downloader instance
//	dl := downloader.NewDownloader()
//	oldMaxWorkers := config.DefaultMaxWorkers
//	config.DefaultMaxWorkers = 3
//	defer func() {
//		config.DefaultMaxWorkers = oldMaxWorkers
//	}()
//
//	oldURLsFilePath := config.RelativeURLsFilePath
//	config.RelativeURLsFilePath = "test_relative_urls.txt"
//	defer func() {
//		config.RelativeURLsFilePath = oldURLsFilePath
//	}()
//
//	oldOutputDirectory := config.OutputDirectory
//	config.OutputDirectory = "test_output/"
//	defer func() {
//		config.OutputDirectory = oldOutputDirectory
//	}()
//
//	tests := []struct {
//		name string
//		// Add other input parameters if necessary
//	}{
//		{
//			name: "Test Case 1",
//			// Add test parameters here
//		},
//		{
//			name: "Test Case 2",
//			// Add test parameters here
//		},
//		// Add more test cases as needed
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			Run()
//			// Add assertions or validations based on your test case
//		})
//	}
//}
//
//func Run() {
//	maxWorkers := config.GetDefaultMaxWorkers()
//	urlFilePath := filehandler.GetURLsFilePath(config.GetRelativeURLsFile(), config.GetBaseURLsFile())
//	urlList, err := filehandler.ReadURLsFromFile(urlFilePath)
//	if err != nil {
//		handleError("Error reading URLs", err)
//		return
//	}
//
//	err = filehandler.CreateOutputDirectory(config.GetOutputDirectory())
//	if err != nil {
//		handleError("Error creating output directory", err)
//		return
//	}
//
//	workerChannel := make(chan struct{}, maxWorkers)
//	var wg sync.WaitGroup
//
//	for _, url := range urlList {
//		workerChannel <- struct{}{}
//		wg.Add(1)
//
//		go func(url string) {
//			defer func() {
//				<-workerChannel
//				wg.Done()
//			}()
//
//			err := downloader.DownloadHTMLContent(url, config.GetOutputDirectory())
//			if err != nil {
//				fmt.Printf("Error downloading from %s: %s\n", url, err)
//				return
//			}
//		}(url)
//	}
//
//	wg.Wait()
//	fmt.Println("All downloads completed!")
//}
//
//func handleError(msg string, err error) {
//	fmt.Println(msg, ":", err)
//}
