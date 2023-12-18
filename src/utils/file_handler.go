package filehandler

import (
	"os"
	"strings"
)

func ReadURLsFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return []string{}, nil // Return an empty string array if content is empty
	}

	return strings.Split(string(content), "\n"), nil
}

func GetURLsFilePath(relativeURLsFile, baseURLsFile string) string {
	if _, err := os.Stat(relativeURLsFile); err == nil {
		return relativeURLsFile // Use relative path if file exists
	}

	if _, err := os.Stat(baseURLsFile); err == nil {
		return baseURLsFile // Use absolute path if file exists
	}

	return "" // If file not found in both relative and absolute paths
}

func CreateOutputDirectory(outputDirectory string) error {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		return os.Mkdir(outputDirectory, 0755)
	} else if err != nil {
		return err
	}
	return nil
}
