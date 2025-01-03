package main

import (
	"fmt"
	"reflect"
	"sort"
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

	process_set1["Test.exe"] = Process{
		name:       "Test.exe",
		time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Example.exe"] = Process{
		name:       "Example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Program.exe"] = Process{
		name:       "Program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Test_app.exe"] = Process{
		name:       "Test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["My_script.exe"] = Process{
		name:       "My_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Utility_tool.exe"] = Process{
		name:       "Utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Helper.exe"] = Process{
		name:       "Helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Benchmark.exe"] = Process{
		name:       "Benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Diagnostic.exe"] = Process{
		name:       "Diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Update.exe"] = Process{
		name:       "Update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2["Debugger.exe"] = Process{
		name:       "Debugger.exe",
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
				t.Errorf("Expected map[string]Process but got nil")
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
	process_set1_start["Test.exe"] = Process{
		name:       "Test.exe",
		time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
		time_alive: time.Duration(0),
	}

	process_set1["Test.exe"] = Process{
		name:       "Test.exe",
		time_start: time.Date(2024, 11, 11, 13, 4, 54, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("253h19m42s")
			return duration
		}(),
	}
	// SET 1

	// SET 2
	process_set2_start["Example.exe"] = Process{
		name:       "Example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Program.exe"] = Process{
		name:       "Program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Test_app.exe"] = Process{
		name:       "Test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["My_script.exe"] = Process{
		name:       "My_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Utility_tool.exe"] = Process{
		name:       "Utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Helper.exe"] = Process{
		name:       "Helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Benchmark.exe"] = Process{
		name:       "Benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Diagnostic.exe"] = Process{
		name:       "Diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Update.exe"] = Process{
		name:       "Update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: time.Duration(0),
	}
	process_set2_start["Debugger.exe"] = Process{
		name:       "Debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: time.Duration(0),
	}

	// ===

	process_set2["Example.exe"] = Process{
		name:       "Example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("444h30m20s")
			return duration
		}(),
	}
	process_set2["Program.exe"] = Process{
		name:       "Program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set2["Test_app.exe"] = Process{
		name:       "Test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("993h14m27s")
			return duration
		}(),
	}
	process_set2["My_script.exe"] = Process{
		name:       "My_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("9h0m47s")
			return duration
		}(),
	}
	process_set2["Utility_tool.exe"] = Process{
		name:       "Utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("226h27m22s")
			return duration
		}(),
	}
	process_set2["Helper.exe"] = Process{
		name:       "Helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("447h15m10s")
			return duration
		}(),
	}
	process_set2["Benchmark.exe"] = Process{
		name:       "Benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("1862h15m1s")
			return duration
		}(),
	}
	process_set2["Diagnostic.exe"] = Process{
		name:       "Diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("439h59m49s")
			return duration
		}(),
	}
	process_set2["Update.exe"] = Process{
		name:       "Update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2723h10m16s")
			return duration
		}(),
	}
	process_set2["Debugger.exe"] = Process{
		name:       "Debugger.exe",
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
				t.Errorf("Expected map[string]Process but got nil")
			}
		})
	}

}

func TestCapitalizeFirstLetter(t *testing.T) {

	tests := []struct {
		name     string
		lowered  string
		expected string
	}{
		{"Test 1", "test", "Test"},
		{"Test 2", "foo", "Foo"},
		{"Test 3", "bar", "Bar"},
		{"Test 4", "Bar", "Bar"},
		{"Test 5", "b", "B"},
		{"Test 6", "B", "B"},
		{"Test 7", "", ""},
		{"Test 8", ",", ","},
		{"Test 9", "./';a", "./';a"},
		{"Test 10", "this is a long sentence for testing, i hope you like it", "This is a long sentence for testing, i hope you like it"},
		{"Test 11", "abcdefg", "Abcdefg"},
		{"Test 12", "steven", "Steven"},
		{"Test 13", "apple", "Apple"},
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

func TestExtractKeys(t *testing.T) {

	keySet1 := map[string]Process{
		"A": {}, "B": {}, "C": {}, "D": {}, "E": {}, "F": {},
		"G": {}, "H": {}, "I": {}, "J": {}, "K": {}, "L": {},
	}
	keys1 := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

	keySet2 := map[string]Process{
		"foo": {}, "bar": {}, "temp": {}, "test": {},
	}
	keys2 := []string{"foo", "bar", "temp", "test"}

	keySet3 := map[string]Process{}
	keys3 := []string{}

	keySet4 := map[string]Process{
		"One": {}, "Two": {}, "Three": {}, "four": {},
	}
	keys4 := []string{"One", "Two", "Three", "four"}

	tests := []struct {
		name     string
		keys     map[string]Process
		expected []string
	}{
		{"Test 1", keySet1, keys1},
		{"Test 2", keySet2, keys2},
		{"Test 3", keySet3, keys3},
		{"Test 4", keySet4, keys4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := extractKeys(tc.keys)

			sort.Strings(result)
			sort.Strings(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestProcessMapToStringSortedByName(t *testing.T) {

	process_set1 := make(map[string]Process)
	process_set2 := make(map[string]Process)
	process_set3 := make(map[string]Process)

	process_set1["Example.exe"] = Process{
		name:       "Example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("444h30m20.00000s")
			return duration
		}(),
	}
	process_set1["Program.exe"] = Process{
		name:       "Program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set1["Test_app.exe"] = Process{
		name:       "Test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("993h14m27s")
			return duration
		}(),
	}
	process_set1["My_script.exe"] = Process{
		name:       "My_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("9h0m47s")
			return duration
		}(),
	}
	process_set1["Utility_tool.exe"] = Process{
		name:       "Utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("226h27m22s")
			return duration
		}(),
	}
	process_set1["Helper.exe"] = Process{
		name:       "Helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("447h15m10s")
			return duration
		}(),
	}
	process_set1["Benchmark.exe"] = Process{
		name:       "Benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("1862h15m1s")
			return duration
		}(),
	}
	process_set1["Diagnostic.exe"] = Process{
		name:       "Diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("439h59m49s")
			return duration
		}(),
	}
	process_set1["Update.exe"] = Process{
		name:       "Update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2723h10m16s")
			return duration
		}(),
	}
	process_set1["Debugger.exe"] = Process{
		name:       "Debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("450h45m0s")
			return duration
		}(),
	}

	process_set2["App1.exe"] = Process{
		name:       "App1.exe",
		time_start: time.Date(2024, 11, 24, 14, 30, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2h15m")
			return duration
		}(),
	}

	process_set2["Service.exe"] = Process{
		name:       "Service.exe",
		time_start: time.Date(2024, 11, 23, 9, 0, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("48h30m")
			return duration
		}(),
	}

	process_set2["Tool.exe"] = Process{
		name:       "Tool.exe",
		time_start: time.Date(2024, 11, 24, 6, 15, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("30m")
			return duration
		}(),
	}

	process_set1_string1 := "Utility_tool.exe, 2024-11-10 10:12:34, 9 days 10 hours 27 minutes 22 seconds,\nUpdate.exe, 2024-07-10 12:34:56, 113 days 11 hours 10 minutes 16 seconds,\nTest_app.exe, 2024-09-20 14:30:45, 41 days 9 hours 14 minutes 27 seconds,\nProgram.exe, 2024-10-31 23:45:12, 0 seconds,\nMy_script.exe, 2024-11-01 08:45:59, 9 hours 47 seconds,\nHelper.exe, 2024-11-19 15:00:22, 18 days 15 hours 15 minutes 10 seconds,\nExample.exe, 2024-11-19 12:15:32, 18 days 12 hours 30 minutes 20 seconds,\nDiagnostic.exe, 2024-11-19 07:45:01, 18 days 7 hours 59 minutes 49 seconds,\nDebugger.exe, 2024-11-19 18:30:12, 18 days 18 hours 45 minutes,\nBenchmark.exe, 2024-08-15 09:30:11, 77 days 14 hours 15 minutes 1 seconds,\n"
	process_set1_string2 := "Benchmark.exe, 2024-08-15 09:30:11, 77 days 14 hours 15 minutes 1 seconds,\nDebugger.exe, 2024-11-19 18:30:12, 18 days 18 hours 45 minutes,\nDiagnostic.exe, 2024-11-19 07:45:01, 18 days 7 hours 59 minutes 49 seconds,\nExample.exe, 2024-11-19 12:15:32, 18 days 12 hours 30 minutes 20 seconds,\nHelper.exe, 2024-11-19 15:00:22, 18 days 15 hours 15 minutes 10 seconds,\nMy_script.exe, 2024-11-01 08:45:59, 9 hours 47 seconds,\nProgram.exe, 2024-10-31 23:45:12, 0 seconds,\nTest_app.exe, 2024-09-20 14:30:45, 41 days 9 hours 14 minutes 27 seconds,\nUpdate.exe, 2024-07-10 12:34:56, 113 days 11 hours 10 minutes 16 seconds,\nUtility_tool.exe, 2024-11-10 10:12:34, 9 days 10 hours 27 minutes 22 seconds,\n"
	process_set2_string1 := "Tool.exe, 2024-11-24 06:15:00, 30 minutes,\nService.exe, 2024-11-23 09:00:00, 2 days 30 minutes,\nApp1.exe, 2024-11-24 14:30:00, 2 hours 15 minutes,\n"
	process_set2_string2 := "App1.exe, 2024-11-24 14:30:00, 2 hours 15 minutes,\nService.exe, 2024-11-23 09:00:00, 2 days 30 minutes,\nTool.exe, 2024-11-24 06:15:00, 30 minutes,\n"
	process_set3_string1 := ""
	process_set3_string2 := ""

	tests := []struct {
		name      string
		processes map[string]Process
		expected  string
	}{
		{"Test 1", process_set1, process_set1_string1},
		{"Test 1 inverse", process_set1, process_set1_string2},
		{"Test 2", process_set2, process_set2_string1},
		{"Test 2 inverse", process_set2, process_set2_string2},
		{"Test 3", process_set3, process_set3_string1},
		{"Test 3 inverse", process_set3, process_set3_string2},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if i%2 == 0 {
				result := ProcessMapToStringSortedByName(tc.processes, true)
				if tc.expected != result {
					t.Errorf("\nExpected\n\n%s \n\n\nbut got result \n\n%s\n\n", tc.expected, result)
				}
			} else {
				result := ProcessMapToStringSortedByName(tc.processes, false)
				if tc.expected != result {
					t.Errorf("\nExpected\n\n%s \n\n\nbut got result \n\n%s\n\n", tc.expected, result)
				}
			}

		})
	}

}
func TestProcessMapToStringSortedByTimeAlive(t *testing.T) {

	process_set1 := make(map[string]Process)
	process_set2 := make(map[string]Process)
	process_set3 := make(map[string]Process)

	process_set1["Example.exe"] = Process{
		name:       "Example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("444h30m20s")
			return duration
		}(),
	}
	process_set1["Program.exe"] = Process{
		name:       "Program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set1["Test_app.exe"] = Process{
		name:       "Test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("993h14m27s")
			return duration
		}(),
	}
	process_set1["My_script.exe"] = Process{
		name:       "My_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("9h0m47s")
			return duration
		}(),
	}
	process_set1["Utility_tool.exe"] = Process{
		name:       "Utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("226h27m22s")
			return duration
		}(),
	}
	process_set1["Helper.exe"] = Process{
		name:       "Helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("447h15m10s")
			return duration
		}(),
	}
	process_set1["Benchmark.exe"] = Process{
		name:       "Benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("1862h15m1s")
			return duration
		}(),
	}
	process_set1["Diagnostic.exe"] = Process{
		name:       "Diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("439h59m49s")
			return duration
		}(),
	}
	process_set1["Update.exe"] = Process{
		name:       "Update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2723h10m16s")
			return duration
		}(),
	}
	process_set1["Debugger.exe"] = Process{
		name:       "Debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("450h45m0s")
			return duration
		}(),
	}

	process_set2["App1.exe"] = Process{
		name:       "App1.exe",
		time_start: time.Date(2024, 11, 24, 14, 30, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2h15m")
			return duration
		}(),
	}

	process_set2["Service.exe"] = Process{
		name:       "Service.exe",
		time_start: time.Date(2024, 11, 23, 9, 0, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("48h30m")
			return duration
		}(),
	}

	process_set2["Tool.exe"] = Process{
		name:       "Tool.exe",
		time_start: time.Date(2024, 11, 24, 6, 15, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("30m")
			return duration
		}(),
	}

	process_set1_string1 := "Program.exe, 2024-10-31 23:45:12, 0 seconds,\nMy_script.exe, 2024-11-01 08:45:59, 9 hours 47 seconds,\nUtility_tool.exe, 2024-11-10 10:12:34, 9 days 10 hours 27 minutes 22 seconds,\nDiagnostic.exe, 2024-11-19 07:45:01, 18 days 7 hours 59 minutes 49 seconds,\nExample.exe, 2024-11-19 12:15:32, 18 days 12 hours 30 minutes 20 seconds,\nHelper.exe, 2024-11-19 15:00:22, 18 days 15 hours 15 minutes 10 seconds,\nDebugger.exe, 2024-11-19 18:30:12, 18 days 18 hours 45 minutes,\nTest_app.exe, 2024-09-20 14:30:45, 41 days 9 hours 14 minutes 27 seconds,\nBenchmark.exe, 2024-08-15 09:30:11, 77 days 14 hours 15 minutes 1 seconds,\nUpdate.exe, 2024-07-10 12:34:56, 113 days 11 hours 10 minutes 16 seconds,\n"
	process_set1_string2 := "Update.exe, 2024-07-10 12:34:56, 113 days 11 hours 10 minutes 16 seconds,\nBenchmark.exe, 2024-08-15 09:30:11, 77 days 14 hours 15 minutes 1 seconds,\nTest_app.exe, 2024-09-20 14:30:45, 41 days 9 hours 14 minutes 27 seconds,\nDebugger.exe, 2024-11-19 18:30:12, 18 days 18 hours 45 minutes,\nHelper.exe, 2024-11-19 15:00:22, 18 days 15 hours 15 minutes 10 seconds,\nExample.exe, 2024-11-19 12:15:32, 18 days 12 hours 30 minutes 20 seconds,\nDiagnostic.exe, 2024-11-19 07:45:01, 18 days 7 hours 59 minutes 49 seconds,\nUtility_tool.exe, 2024-11-10 10:12:34, 9 days 10 hours 27 minutes 22 seconds,\nMy_script.exe, 2024-11-01 08:45:59, 9 hours 47 seconds,\nProgram.exe, 2024-10-31 23:45:12, 0 seconds,\n"
	process_set2_string1 := "Tool.exe, 2024-11-24 06:15:00, 30 minutes,\nApp1.exe, 2024-11-24 14:30:00, 2 hours 15 minutes,\nService.exe, 2024-11-23 09:00:00, 2 days 30 minutes,\n"
	process_set2_string2 := "Service.exe, 2024-11-23 09:00:00, 2 days 30 minutes,\nApp1.exe, 2024-11-24 14:30:00, 2 hours 15 minutes,\nTool.exe, 2024-11-24 06:15:00, 30 minutes,\n"
	process_set3_string1 := ""
	process_set3_string2 := ""

	tests := []struct {
		name      string
		processes map[string]Process
		expected  string
	}{
		{"Test 1", process_set1, process_set1_string1},
		{"Test 1 inverse", process_set1, process_set1_string2},
		{"Test 2", process_set2, process_set2_string1},
		{"Test 2 inverse", process_set2, process_set2_string2},
		{"Test 3", process_set3, process_set3_string1},
		{"Test 3 inverse", process_set3, process_set3_string2},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if i%2 == 0 {
				result := ProcessMapToStringSortedByTimeAlive(tc.processes, false)
				if tc.expected != result {
					t.Errorf("\nExpected\n\n%s \n\n\nbut got result \n\n%s\n\n", tc.expected, result)
				}
			} else {
				result := ProcessMapToStringSortedByTimeAlive(tc.processes, true)
				if tc.expected != result {
					t.Errorf("\nExpected\n\n%s but got result \n\n%s\n\n", tc.expected, result)
				}
			}

		})
	}

}

func TestProcessMapToStringSortedByTimeStarted(t *testing.T) {

	process_set1 := make(map[string]Process)
	process_set2 := make(map[string]Process)
	process_set3 := make(map[string]Process)

	process_set1["Example.exe"] = Process{
		name:       "Example.exe",
		time_start: time.Date(2024, 11, 19, 12, 15, 32, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("444h30m20s")
			return duration
		}(),
	}
	process_set1["Program.exe"] = Process{
		name:       "Program.exe",
		time_start: time.Date(2024, 10, 31, 23, 45, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set1["Test_app.exe"] = Process{
		name:       "Test_app.exe",
		time_start: time.Date(2024, 9, 20, 14, 30, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("993h14m27s")
			return duration
		}(),
	}
	process_set1["My_script.exe"] = Process{
		name:       "My_script.exe",
		time_start: time.Date(2024, 11, 1, 8, 45, 59, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("9h0m47s")
			return duration
		}(),
	}
	process_set1["Utility_tool.exe"] = Process{
		name:       "Utility_tool.exe",
		time_start: time.Date(2024, 11, 10, 10, 12, 34, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("226h27m22s")
			return duration
		}(),
	}
	process_set1["Helper.exe"] = Process{
		name:       "Helper.exe",
		time_start: time.Date(2024, 11, 19, 15, 0, 22, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("447h15m10s")
			return duration
		}(),
	}
	process_set1["Benchmark.exe"] = Process{
		name:       "Benchmark.exe",
		time_start: time.Date(2024, 8, 15, 9, 30, 11, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("1862h15m1s")
			return duration
		}(),
	}
	process_set1["Diagnostic.exe"] = Process{
		name:       "Diagnostic.exe",
		time_start: time.Date(2024, 11, 19, 7, 45, 1, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("439h59m49s")
			return duration
		}(),
	}
	process_set1["Update.exe"] = Process{
		name:       "Update.exe",
		time_start: time.Date(2024, 7, 10, 12, 34, 56, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2723h10m16s")
			return duration
		}(),
	}
	process_set1["Debugger.exe"] = Process{
		name:       "Debugger.exe",
		time_start: time.Date(2024, 11, 19, 18, 30, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("450h45m0s")
			return duration
		}(),
	}

	process_set2["App1.exe"] = Process{
		name:       "App1.exe",
		time_start: time.Date(2024, 11, 24, 14, 30, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("2h15m")
			return duration
		}(),
	}

	process_set2["Service.exe"] = Process{
		name:       "Service.exe",
		time_start: time.Date(2024, 11, 23, 9, 0, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("48h30m")
			return duration
		}(),
	}

	process_set2["Tool.exe"] = Process{
		name:       "Tool.exe",
		time_start: time.Date(2024, 11, 24, 6, 15, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("30m")
			return duration
		}(),
	}

	process_set1_string1 := "Debugger.exe, 2024-11-19 18:30:12, 18 days 18 hours 45 minutes,\nHelper.exe, 2024-11-19 15:00:22, 18 days 15 hours 15 minutes 10 seconds,\nExample.exe, 2024-11-19 12:15:32, 18 days 12 hours 30 minutes 20 seconds,\nDiagnostic.exe, 2024-11-19 07:45:01, 18 days 7 hours 59 minutes 49 seconds,\nUtility_tool.exe, 2024-11-10 10:12:34, 9 days 10 hours 27 minutes 22 seconds,\nMy_script.exe, 2024-11-01 08:45:59, 9 hours 47 seconds,\nProgram.exe, 2024-10-31 23:45:12, 0 seconds,\nTest_app.exe, 2024-09-20 14:30:45, 41 days 9 hours 14 minutes 27 seconds,\nBenchmark.exe, 2024-08-15 09:30:11, 77 days 14 hours 15 minutes 1 seconds,\nUpdate.exe, 2024-07-10 12:34:56, 113 days 11 hours 10 minutes 16 seconds,\n"
	process_set1_string2 := "Update.exe, 2024-07-10 12:34:56, 113 days 11 hours 10 minutes 16 seconds,\nBenchmark.exe, 2024-08-15 09:30:11, 77 days 14 hours 15 minutes 1 seconds,\nTest_app.exe, 2024-09-20 14:30:45, 41 days 9 hours 14 minutes 27 seconds,\nProgram.exe, 2024-10-31 23:45:12, 0 seconds,\nMy_script.exe, 2024-11-01 08:45:59, 9 hours 47 seconds,\nUtility_tool.exe, 2024-11-10 10:12:34, 9 days 10 hours 27 minutes 22 seconds,\nDiagnostic.exe, 2024-11-19 07:45:01, 18 days 7 hours 59 minutes 49 seconds,\nExample.exe, 2024-11-19 12:15:32, 18 days 12 hours 30 minutes 20 seconds,\nHelper.exe, 2024-11-19 15:00:22, 18 days 15 hours 15 minutes 10 seconds,\nDebugger.exe, 2024-11-19 18:30:12, 18 days 18 hours 45 minutes,\n"
	process_set2_string1 := "App1.exe, 2024-11-24 14:30:00, 2 hours 15 minutes,\nTool.exe, 2024-11-24 06:15:00, 30 minutes,\nService.exe, 2024-11-23 09:00:00, 2 days 30 minutes,\n"
	process_set2_string2 := "Service.exe, 2024-11-23 09:00:00, 2 days 30 minutes,\nTool.exe, 2024-11-24 06:15:00, 30 minutes,\nApp1.exe, 2024-11-24 14:30:00, 2 hours 15 minutes,\n"
	process_set3_string1 := ""
	process_set3_string2 := ""

	tests := []struct {
		name      string
		processes map[string]Process
		expected  string
	}{
		{"Test 1", process_set1, process_set1_string1},
		{"Test 1 inverse", process_set1, process_set1_string2},
		{"Test 2", process_set2, process_set2_string1},
		{"Test 2 inverse", process_set2, process_set2_string2},
		{"Test 3", process_set3, process_set3_string1},
		{"Test 3 inverse", process_set3, process_set3_string2},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if i%2 == 0 {
				result := ProcessMapToStringSortedByTimeStarted(tc.processes, true)
				if tc.expected != result {
					t.Errorf("\nExpected\n\n%s \n\n\nbut got result \n\n%s\n\n", tc.expected, result)
				}
			} else {
				result := ProcessMapToStringSortedByTimeStarted(tc.processes, false)
				if tc.expected != result {
					t.Errorf("\nExpected\n\n%s \n\n\nbut got result \n\n%s\n\n", tc.expected, result)
				}
			}

		})
	}

}

func TestParseTime(t *testing.T) {
	invalidDateTime := time.Time{}

	tests := []struct {
		name        string
		timeString  string
		expected    time.Time
		expectError bool
	}{
		{"Test 1   Valid", "20241201231957.323450+000", time.Date(2024, 12, 01, 23, 19, 57, 0, time.UTC), false},
		{"Test 1   Non Valid", "This should NOT work", invalidDateTime, true},
		{"Test 2   Valid", "20231130220030.123456+000", time.Date(2023, 11, 30, 22, 0, 30, 0, time.UTC), false},
		{"Test 2   Non Valid", "23X30704121212.345678+000", invalidDateTime, true},
		{"Test 3   Valid", "20240101000000.000000+000", time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC), false},
		{"Test 3   Non Valid", "2023-11-30T220030+000", invalidDateTime, true},
		{"Test 4   Valid", "20241225123045.654321+000", time.Date(2024, 12, 25, 12, 30, 45, 0, time.UTC), false},
		{"Test 4   Non Valid", "2023/11/30/220030+000", invalidDateTime, true},
		{"Test 5   Valid", "20240315101515.999999+000", time.Date(2024, 03, 15, 10, 15, 15, 0, time.UTC), false},
		{"Test 5   Non Valid", "20230704121212:345678+000", invalidDateTime, true},
		{"Test 6   Valid", "20230704121212.345678+000", time.Date(2023, 07, 04, 12, 12, 12, 0, time.UTC), false},
		{"Test 6   Non Valid", "20230704121212.ABCDEF+000", invalidDateTime, true},
		{"Test 7   Valid", "20221115074525.123456+000", time.Date(2022, 11, 15, 07, 45, 25, 0, time.UTC), false},
		{"Test 7   Non Valid", "202307041212.3456781212+000", invalidDateTime, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := parseTime(tc.timeString)

			if tc.expectError == true && err == nil {
				t.Errorf("\n\nExpected an error exception but error was %s\nIt's possible that the time was parsed 'correctly' when it shouldn't, the result was %s\n", err, result)
			} else if result != tc.expected {
				t.Errorf("\n\nExpected parsed time: %s\n\nBut got %s\n", tc.expected, result)
			}

		})
	}
}

func TestReformatDate(t *testing.T) {
	tests := []struct {
		name        string
		dateString  string
		expected    string
		expectError bool
	}{
		{"Test 1   Valid", "2024-12-01 16:22:30 +0000 UTC", "2024-12-01 16:22:30", false},
		{"Test 2   Non-Valid", "20AB-12-01 16:22:30 +0000 UTC", "", true},
		{"Test 3   Valid", "2022-03-17 14:52:10 +0000 UTC", "2022-03-17 14:52:10", false},
		{"Test 4   Non-Valid", "2024-13-01 25:60:00 +0000 UTC", "", true},
		{"Test 5   Valid", "2004-07-08 23:17:59 +0000 UTC", "2004-07-08 23:17:59", false},
		{"Test 6   Non-Valid", "2021-07-45 99:99:99 +0000 UTC", "", true},
		{"Test 7   Valid", "2018-01-25 08:04:01 +0000 UTC", "2018-01-25 08:04:01", false},
		{"Test 8    Non-Valid", "2020-02-30 XX:XX:XX +0000 UTC", "", true},
		{"Test 9   Valid", "2007-05-14 16:11:50 +0000 UTC", "2007-05-14 16:11:50", false},
		{"Test 10   Non-Valid", "2019-11-99 14:AB:CD +0000 UTC", "", true},
		{"Test 11   Valid", "2016-12-31 03:22:40 +0000 UTC", "2016-12-31 03:22:40", false},
		{"Test 12   Non-Valid", "2018-06-15 32:61:01 +0000 UTC", "", true},
		{"Test 13   Valid", "2019-10-22 10:56:12 +0000 UTC", "2019-10-22 10:56:12", false},
		{"Test 14   Non-Valid", "2017-14-01 27:15:45 +0000 UTC", "", true},
		{"Test 15   Valid", "2014-06-09 12:00:00 +0000 UTC", "2014-06-09 12:00:00", false},
		{"Test 16   Non-Valid", "2016-05-01 99:99:99 +0000 UTC", "", true},
		{"Test 17   Valid", "2010-04-18 21:45:20 +0000 UTC", "2010-04-18 21:45:20", false},
		{"Test 18   Non-Valid", "2015-03-00 07:70:70 +0000 UTC", "", true},
		{"Test 19   Valid", "2012-08-06 05:30:10 +0000 UTC", "2012-08-06 05:30:10", false},
		{"Test 20   Non-Valid", "2014-01-01 99:99:ZZ +0000 UTC", "", true},
		{"Test 21   Valid", "2011-02-28 18:12:59 +0000 UTC", "2011-02-28 18:12:59", false},
		{"Test 22   Non-Valid", "2013-02-29 60:00:00 +0000 UTC", "", true},
		{"Test 23   Valid", "2010-10-03 07:22:35 +0000 UTC", "2010-10-03 07:22:35", false},
		{"Test 24   Non-Valid", "2012-12-01 25:00:00 +0000 UTC", "", true},
		{"Test 25  Valid", "2009-12-12 09:50:01 +0000 UTC", "2009-12-12 09:50:01", false},
		{"Test 26   Non-Valid", "2011-10-99 14:00:00 +0000 UTC", "", true},
		{"Test 27   Valid", "2008-11-19 15:42:33 +0000 UTC", "2008-11-19 15:42:33", false},
		{"Test 28   Non-Valid", "2010-09-31 45:15:30 +0000 UTC", "", true},
		{"Test 29   Valid", "2001-12-01 16:22:30 +0000 UTC", "2001-12-01 16:22:30", false},
		{"Test 30   Non-Valid", "2009-08-00 12:12:XY +0000 UTC", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := reformatDate(tc.dateString)
			if tc.expectError == true && err == nil {
				t.Errorf("\n\nExpected an error exception but error was %s\nIt's possible that the date was parsed 'correctly' when it shouldn't, the result was %s\n", err, result)
			} else if result != tc.expected {
				t.Errorf("\n\nExpected parsed date: %s\n\nBut got %s\n", tc.expected, result)
			}

		})
	}
}

func TestReformatDuration(t *testing.T) {
	tests := []struct {
		name           string
		durationString string
		expected       string
	}{
		{"Test 1", "69h54m23.0496006s", "2 days 21 hours 54 minutes 23 seconds"},
		{"Test 2", "123h5m50s", "5 days 3 hours 5 minutes 50 seconds"},
		{"Test 3", "12h0m0s", "12 hours"},
		{"Test 4", "0h30m15s", "30 minutes 15 seconds"},
		{"Test 5", "48h0m0s", "2 days"},
		{"Test 6", "96h45m22s", "4 days 45 minutes 22 seconds"},
		{"Test 7", "500h5m10s", "20 days 20 hours 5 minutes 10 seconds"},
		{"Test 8", "24h0m0s", "1 days"},
		{"Test 9", "9h15m35s", "9 hours 15 minutes 35 seconds"},
		{"Test 10", "0h0m45s", "45 seconds"},
		{"Test 11", "78h12m59s", "3 days 6 hours 12 minutes 59 seconds"},
		{"Test 12", "10h10m5s", "10 hours 10 minutes 5 seconds"},
		{"Test 13", "360h0m0s", "15 days"},
		{"Test 14", "0h5m5s", "5 minutes 5 seconds"},
		{"Test 15", "200h59m59s", "8 days 8 hours 59 minutes 59 seconds"},
		{"Test 16", "60h0m0s", "2 days 12 hours"},
		{"Test 17", "0h0m0s", "0 seconds"},
		{"Test 18", "999h0m0s", "41 days 15 hours"},
		{"Test 19", "72h1m20s", "3 days 1 minutes 20 seconds"},
		{"Test 20", "0h1m5s", "1 minutes 5 seconds"},
		{"Test 21", "250h0m0s", "10 days 10 hours"},
		{"Test 22", "0h10m10s", "10 minutes 10 seconds"},
		{"Test 23", "57h48m12s", "2 days 9 hours 48 minutes 12 seconds"},
		{"Test 24", "333h12m0s", "13 days 21 hours 12 minutes"},
		{"Test 25", "480h30m0s", "20 days 30 minutes"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := reformatDuration(tc.durationString)
			if result != tc.expected {
				t.Errorf("\n\nThe result given by reformatDuration was %s\n\nwhen\n\n%s was expected\n\n", result, tc.expected)
			}

		})
	}

}

func TestParseProcesses(t *testing.T) {

	process_string_prelim := "Caption                             CreationDate\n System Idle Process                 20241127220850.650061+000\n System                              20241127220850.650061+000\nSecure System                       20241127220755.656859+000\nRegistry                            20241127220755.763257+000\n"
	process_string1 := process_string_prelim + "smss.exe                 20241127220850.650061+000\ncsrss.exe                           20241127220854.462581+000\nwininit.exe                         20241127220855.811044+000\nservices.exe                        20241127220855.857921+000\n"
	process_string2 := process_string_prelim + "explorer.exe                        20230814104512.342100+000\nlsass.exe                           20211225181245.987654+000\nsvchost.exe                         20231009130918.294718+000\nnotepad.exe                         20240917231425.837654+000\n"
	process_string3 := process_string_prelim + "cmd.exe                             20210706134502.142233+000\npowershell.exe                      20221220175510.545566+000\n"
	process_string4 := process_string_prelim + "chrome.exe                          20210203143000.129084+000\n"
	process_string5 := process_string_prelim + "outlook.exe                         20230414101230.783210+000\nteams.exe                           20220125124545.942183+000\nonenote.exe                         20240807194510.537281+000\nedge.exe                            20230330150225.781092+000\n"

	process_set1 := make(map[string]Process)
	process_set2 := make(map[string]Process)
	process_set3 := make(map[string]Process)
	process_set4 := make(map[string]Process)
	process_set5 := make(map[string]Process)

	// SET 1
	process_set1["Smss.exe"] = Process{ // 20241127220850.650061+000
		name:       "Smss.exe",
		time_start: time.Date(2024, 11, 27, 22, 8, 50, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set1["Csrss.exe"] = Process{ // 20241127220854.462581
		name:       "Csrss.exe",
		time_start: time.Date(2024, 11, 27, 22, 8, 54, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set1["Wininit.exe"] = Process{ // 20241127220855.811044+000
		name:       "Wininit.exe",
		time_start: time.Date(2024, 11, 27, 22, 8, 55, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set1["Services.exe"] = Process{ //20241127220855.857921+000
		name:       "Services.exe",
		time_start: time.Date(2024, 11, 27, 22, 8, 55, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	// SET 1

	// SET 2
	process_set2["Explorer.exe"] = Process{ // 20230814104512.342100+000
		name:       "Explorer.exe",
		time_start: time.Date(2023, 8, 14, 10, 45, 12, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set2["Lsass.exe"] = Process{ // 20211225181245.987654+000
		name:       "Lsass.exe",
		time_start: time.Date(2021, 12, 25, 18, 12, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set2["Svchost.exe"] = Process{ // 20231009130918.294718+000
		name:       "Svchost.exe",
		time_start: time.Date(2023, 10, 9, 13, 9, 18, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set2["Notepad.exe"] = Process{ // 20240917231425.837654+000
		name:       "Notepad.exe",
		time_start: time.Date(2024, 9, 17, 23, 14, 25, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}

	// SET 2

	// SET 3
	process_set3["Cmd.exe"] = Process{ // 20210706134502.142233+000
		name:       "Cmd.exe",
		time_start: time.Date(2021, 7, 6, 13, 45, 2, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set3["Powershell.exe"] = Process{ // 20221220175510.545566+000
		name:       "Powershell.exe",
		time_start: time.Date(2022, 12, 20, 17, 55, 10, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	// SET 3

	// SET 4
	process_set4["Chrome.exe"] = Process{ // 20210203143000.129084+000
		name:       "Chrome.exe",
		time_start: time.Date(2021, 2, 3, 14, 30, 0, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	// SET 4

	// SET 5
	process_set5["Outlook.exe"] = Process{ // 20230414101230.783210+000
		name:       "Outlook.exe",
		time_start: time.Date(2023, 4, 14, 10, 12, 30, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}

	process_set5["Teams.exe"] = Process{ // 20220125124545.942183+000
		name:       "Teams.exe",
		time_start: time.Date(2022, 1, 25, 12, 45, 45, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set5["Onenote.exe"] = Process{ // 20240807194510.537281+000
		name:       "Onenote.exe",
		time_start: time.Date(2024, 8, 7, 19, 45, 10, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	process_set5["Edge.exe"] = Process{ // 20230330150225.781092+000
		name:       "Edge.exe",
		time_start: time.Date(2023, 3, 30, 15, 2, 25, 0, time.UTC),
		time_alive: func() time.Duration {
			duration, _ := time.ParseDuration("0s")
			return duration
		}(),
	}
	// SET 5

	tests := []struct {
		name          string
		processString string
		expected      map[string]Process
	}{
		{"Test 1", process_string1, process_set1},
		{"Test 2", process_string2, process_set2},
		{"Test 3", process_string3, process_set3},
		{"Test 4", process_string4, process_set4},
		{"Test 5", process_string5, process_set5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := ParseProcesses(tc.processString)
			if reflect.DeepEqual(result, tc.expected) == false {
				result_keys := extractKeys(result)
				expected_keys := extractKeys(tc.expected)
				sort.Strings(result_keys)
				sort.Strings(expected_keys)

				if reflect.DeepEqual(result_keys, expected_keys) == false {
					t.Errorf("The resulting map does not match the expected map based on the keys extracted from both maps not matching")

					t.Errorf("\n\nResult keys: \n")
					for _, key := range result_keys {
						t.Errorf("%s\n", key)
					}
					t.Errorf("\n\nExpected keys: \n")
					for _, key := range expected_keys {
						t.Errorf("%s\n", key)
					}

				} else {
					for _, key := range result_keys {

						t.Errorf("\n\nKEY: %s\n", key)
						t.Errorf("Got: %s\tExpected: %s\n", result[key].name, tc.expected[key].name)
						t.Errorf("Got: %s\tExpected: %s\n", result[key].time_start, tc.expected[key].time_start)
						t.Errorf("Got: %s\tExpected: %s\n", result[key].time_alive, tc.expected[key].time_alive)

					}

				}
			}

		})
	}
}
