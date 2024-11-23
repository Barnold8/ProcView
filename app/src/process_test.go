package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// HELPER FUNCTIONS

func processToString(process Process) string {

	return fmt.Sprintf("%s %s", process.name, process.time_start.Format("20060102150405")+fmt.Sprintf(".%06d+000", process.time_start.Nanosecond()/1000))

}

func processSetToString(processes map[string]Process) string {

	curr_string := ""

	for key, _ := range processes {
		curr_string += processToString(processes[key]) + "\n"
	}
	return curr_string
}

// HELPER FUNCTIONS

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

func TestGatherProcesses(t *testing.T) {

	process_set1 := make(map[string]Process)
	process_set2 := make(map[string]Process)

	process_set1["test.exe"] = Process{
		name:       "test.exe",
		time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
		time_alive: time.Duration(0),
	}

	process_set2["example.exe"] = Process{
		name:       "example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["program.exe"] = Process{
		name:       "program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["test_app.exe"] = Process{
		name:       "test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["my_script.exe"] = Process{
		name:       "my_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["utility_tool.exe"] = Process{
		name:       "utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["helper.exe"] = Process{
		name:       "helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["benchmark.exe"] = Process{
		name:       "benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["diagnostic.exe"] = Process{
		name:       "diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["update.exe"] = Process{
		name:       "update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["debugger.exe"] = Process{
		name:       "debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}

	// TESTS
	tests := []struct {
		name        string
		proc_string string
		expected    map[string]Process
	}{
		{"Test 1", "test.exe                         20241111130454.507396+000", process_set1},
		{"Test 2", "sadgasqaweouhui asdaskdhk 132412489797", make(map[string]Process)},
		{"Test 3", "example.exe                      20241119121532.003452+000\n program.exe                      20241031234512.123456+000\n test_app.exe                     20240920143045.654321+000\n my_script.exe                    20241101084559.789012+000\n utility_tool.exe                 20241110101234.567890+000\n helper.exe                       20241119150022.000123+000\n benchmark.exe                    20240815093011.098765+000\n diagnostic.exe                   20241119074501.432109+000\n update.exe                       20240710123456.876543+000\n debugger.exe                     20241119183012.654321+000\n", process_set2},
	}
	// TESTS

	// Actual testing
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ParseProcesses(tc.proc_string)
			if result != nil {
				if !reflect.DeepEqual(tc.expected, result) {
					t.Errorf("Maps do not match.\n\n\n\nExpected: %#v\n\n\n\n\n\nGot: %#v", tc.expected, result)
				}
			} else {
				t.Errorf("Expected []Process but got nil")
			}
		})
	}
	// Actual testing
}

func TestUpdateProcesses(t *testing.T) {

	// Write a starting point for processes

	time_start := time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC)

	process_set1 := make(map[string]Process)
	process_set1_start := make(map[string]Process)
	process_set2_start := make(map[string]Process)
	process_set2 := make(map[string]Process)

	// SET 1
	process_set1_start["test.exe"] = Process{
		name:       "test.exe",
		time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
		time_alive: time.Duration(0),
	}

	process_set1["test.exe"] = Process{
		name:       "test.exe",
		time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("253h19m42s")
			return duration
		}(),
	}
	// SET 1

	// SET 2
	process_set2_start["example.exe"] = Process{
		name:       "example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["program.exe"] = Process{
		name:       "program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["test_app.exe"] = Process{
		name:       "test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["my_script.exe"] = Process{
		name:       "my_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["utility_tool.exe"] = Process{
		name:       "utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["helper.exe"] = Process{
		name:       "helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["benchmark.exe"] = Process{
		name:       "benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["diagnostic.exe"] = Process{
		name:       "diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["update.exe"] = Process{
		name:       "update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["debugger.exe"] = Process{
		name:       "debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}

	// ===

	process_set2["example.exe"] = Process{
		name:       "example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("444h30m20s")
			return duration
		}(),
	}
	process_set2["program.exe"] = Process{
		name:       "program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set2["test_app.exe"] = Process{
		name:       "test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("993h14m27s")
			return duration
		}(),
	}
	process_set2["my_script.exe"] = Process{
		name:       "my_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("9h0m47s")
			return duration
		}(),
	}
	process_set2["utility_tool.exe"] = Process{
		name:       "utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("226h27m22s")
			return duration
		}(),
	}
	process_set2["helper.exe"] = Process{
		name:       "helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("447h15m10s")
			return duration
		}(),
	}
	process_set2["benchmark.exe"] = Process{
		name:       "benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("1862h15m1s")
			return duration
		}(),
	}
	process_set2["diagnostic.exe"] = Process{
		name:       "diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("439h59m49s")
			return duration
		}(),
	}
	process_set2["update.exe"] = Process{
		name:       "update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2723h10m16s")
			return duration
		}(),
	}
	process_set2["debugger.exe"] = Process{
		name:       "debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("450h45m0s")
			return duration
		}(),
	}

	// SET 2

	tests := []struct {
		name      string
		processes map[string]Process
		expected  map[string]Process
	}{
		{"Test 1", process_set1_start, process_set1},
		{"Test 2", process_set2_start, process_set2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := UpdateProcesses(tc.processes, time_start, processSetToString(tc.expected))
			if result != nil {
				if !reflect.DeepEqual(tc.expected, result) {
					t.Errorf("Maps do not match.\n\n\n\nExpected: %#v\n\n\n\n\n\nGot: %#v", tc.expected, result)
				}
			} else {
				t.Errorf("Expected []Process but got nil")
			}
		})
	}

}

func CapitalizeFirstLetterTest(t *testing.T) {

	tests := []struct {
		name     string
		lowered  string
		expected string
	}{
		{"Test 1", "test", "Test"},
		{"Test 2", "foo", "foo"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := capitalizeFirstLetter(tc.lowered)

			if result != tc.expected {
				t.Errorf("String (%s) was not capitalised correctly. %s was supposed to be returned but got %s", tc.lowered, tc.expected, result)
			}
		})
	}

}

func ExtractKeysTest(t *testing.T) {

}

func ProcessMapToStringTest(t *testing.T) {

}

func ProcessMapToStringSortedByNameTest(t *testing.T) {

}

func ProcessMapToStringSortedByTimeAliveTest(t *testing.T) {

}

func ProcessMapToStringSortedByTimeStartedTest(t *testing.T) {

}
