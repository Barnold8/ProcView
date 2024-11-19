package main

import (
	"reflect"
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

}

// if i ever want to make a function to sort the processes, make a test here
