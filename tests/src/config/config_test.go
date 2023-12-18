package config

import (
	"github.com/amitsol123/html_downloader/src/config"
	"os"
	"testing"
)

func TestGetDefaultMaxWorkers(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedResult int
	}{
		{
			name:           "No args provided",
			args:           []string{"executable"},
			expectedResult: config.GetDefaultMaxWorkers(), // Expect the default value when no args are provided
		},
		{
			name:           "Valid arg provided",
			args:           []string{"executable", "10"}, // Simulate passing a valid argument
			expectedResult: 10,                           // Expect the argument to be parsed and returned
		},
		{
			name:           "Invalid arg provided",
			args:           []string{"executable", "invalid"}, // Simulate passing an invalid argument
			expectedResult: config.GetDefaultMaxWorkers(),     // Expect the default value due to the error
		},
		{
			name:           "Too big arg provided",
			args:           []string{"executable", "999999999999999999999999999"}, // Simulate passing argument too big
			expectedResult: config.GetDefaultMaxWorkers(),                         // Expect the default value when argument is too big
		},
		{
			name:           "Negative arg provided",
			args:           []string{"executable", "-5"},  // Simulate passing negative number of workers
			expectedResult: config.GetDefaultMaxWorkers(), // Expect the default value when number of workers is negative
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Args = test.args // Mock os.Args for the test case

			result := config.GetDefaultMaxWorkers()

			if result != test.expectedResult {
				t.Errorf("Expected result: %d, Got: %d", test.expectedResult, result)
			}
		})
	}
}
