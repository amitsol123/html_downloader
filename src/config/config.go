package config

import (
	"os"
	"strconv"
)

const (
	OutputDirectory   = "downloaded_html/"
	RelativeURLsFile  = "../../ListOfAsciiSiteUrls.txt"
	BaseURLsFile      = "ListOfAsciiSiteUrls.txt"
	DefaultMaxWorkers = 5 // Maximum number of workers
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

func GetOutputDirectory() string {
	return OutputDirectory
}

func GetRelativeURLsFile() string {
	return RelativeURLsFile
}

func GetBaseURLsFile() string {
	return BaseURLsFile
}
