package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeParser(t *testing.T) {

}

func TestGatherProcesses(t *testing.T) {

	// PROCESS SETS
	process_set1 := []Process{
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

	// Actual testing - Data present
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ParseProcesses(tc.proc_string)
			if result != nil {

				for i, expectedProc := range tc.expected {

					if result[i].name != expectedProc.name {
						t.Errorf("Expected name %s, got %s", expectedProc.name, result[i].name)
					}
					if !reflect.DeepEqual(result[i].time_alive, expectedProc.time_alive) {
						t.Errorf("Expected time_alive %v, got %v", expectedProc.time_alive, result[i].time_alive)
					}

				}

			} else {
				t.Errorf("Expected name")
			}
		})
	}
	// Actual testing - Data present

	// Actual testing - No data present

	// Actual testing - No data present

}
