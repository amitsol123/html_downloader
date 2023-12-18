package config

import (
	"os"
	"strconv"
)

const (
	OutputDirectory      = "downloaded_html/"
	RelativeURLsFilePath = "../../ListOfAsciiSiteUrls.txt"
	BaseURLsFilePath     = "ListOfAsciiSiteUrls.txt"
	DefaultMaxWorkers    = 5 // Maximum number of workers
)

func GetDefaultMaxWorkers() int {
	if len(os.Args) > 1 {
		workers, err := strconv.Atoi(os.Args[1])
		if err == nil && workers > 0 {
			return workers
		}
	}
	return DefaultMaxWorkers
}

func GetRelativeURLsFile() string {
	return RelativeURLsFilePath
}

func GetBaseURLsFile() string {
	return BaseURLsFilePath
}

func GetOutputDirectory() string {
	return OutputDirectory
}
