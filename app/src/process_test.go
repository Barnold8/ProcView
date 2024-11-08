package main

import (
	"testing"
	"time"
)

func TestGatherProcesses(t *testing.T) {

	duration, err := time.ParseDuration("2h30m")

	process_set1 := []Process{

		Process{
			name:       "f",
			PID:        1234,
			time_alive: duration,
		},
	}

	tests := []struct {
		name        string
		proc_string string
		expected    []Process
	}{
		{"Correct data", "fffff", process_set1}, // this is not actually correct yet
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ParseProcesses(tc.proc_string)
			if result != nil {

			}
		})
	}
}
