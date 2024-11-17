package main

import (
	"testing"
	"time"
)

func TestTimeParser(t *testing.T) { // TODO: add fail cases where the function handles an error correctly

	tests := []struct {
		name        string
		time_string string
		expected    time.Time
	}{
		{"Time 1", "20241113120608.272105+000", time.Date(2024, time.Month(11), 13, 12, 6, 8, 0, time.UTC)},
		{"Time 2", "20241112003439.986502+000", time.Date(2024, time.Month(11), 12, 0, 34, 39, 0, time.UTC)},
		{"Time 3", "20241113120605.955232+000", time.Date(2024, time.Month(11), 13, 12, 6, 5, 0, time.UTC)},
		{"Time 4", "20241113123910.361191+000", time.Date(2024, time.Month(11), 13, 12, 39, 10, 0, time.UTC)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseTime(tc.time_string)

			if (result == time.Time{} && err != nil) { // case where the function ends and returns an error with empty time when it shouldnt
				t.Errorf("Function expected to return a non error value (%s) but got  an error (%s), instead", tc.expected, err)
			}

			if result != tc.expected {

				t.Errorf("Expected %s but got %s instead", tc.expected.String(), result.String())

			}

		})
	}

}

func TestGatherProcesses(t *testing.T) { // TODO: change this to reflect hashmap

	// PROCESS SETS
	process_set1 := []Process{ //TODO: add more test cases
		Process{
			name:       "test.exe",
			time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
			time_alive: time.Duration(0),
		},
	}
	// PROCESS SETS

	// TESTS
	tests := []struct {
		name        string
		proc_string string
		expected    []Process
	}{
		{"Correct data", "test.exe                         20241111130454.507396+000", process_set1},
	}
	// TESTS

	// Actual testing
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
	// Actual testing
}

// if i ever want to make a function to sort the processes, make a test here
