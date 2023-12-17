package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestGetDefaultMaxWorkers(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"programName"}, 5}, // Mimicking no arguments
		{[]string{"programName", "8"}, 8},
		{[]string{"programName", "invalid"}, 5},
		{[]string{"programName", "-5"}, 5},                                 // Negative number argument
		{[]string{"programName", "9999999999999999999999999999999999"}, 5}, // Too large number
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			os.Args = test.args
			result := getDefaultMaxWorkers()
			if result != test.expected {
				t.Errorf("Expected %d workers, but got %d", test.expected, result)
			}
		})
	}
}

func TestReadURLsFromFile(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		content     string
		expected    []string
		expectedErr error
	}{
		{
			description: "Empty content",
			filename:    "empty.txt",
			content:     "",
			expected:    []string{},
			expectedErr: nil,
		},
		{
			description: "Single URL",
			filename:    "single_url.txt",
			content:     "https://example.com",
			expected:    []string{"https://example.com"},
			expectedErr: nil,
		},
		{
			description: "Multiple URLs",
			filename:    "multiple_urls.txt",
			content:     "https://example.com\nhttps://google.com\nhttps://stackoverflow.com",
			expected: []string{
				"https://example.com",
				"https://google.com",
				"https://stackoverflow.com",
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			tmpfile := createTempFile(t, test.filename, test.content)
			defer os.Remove(tmpfile.Name()) // Clean up the temp file
			defer tmpfile.Close()

			result, err := readURLsFromFile(tmpfile.Name())

			assertResult(t, result, test.expected)
			assertError(t, err, test.expectedErr)
		})
	}
}

// Utility functions for creating/deleting a test file
func createTempFile(t *testing.T, filename, content string) *os.File {
	tmpfile, err := ioutil.TempFile("", filename)
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Error writing to temp file: %s", err)
	}
	return tmpfile
}

func assertResult(t *testing.T, result, expected []string) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func assertError(t *testing.T, err, expectedErr error) {
	if err != expectedErr {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}
