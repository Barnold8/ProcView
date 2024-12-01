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

	fmt.Errorf("Test not yet implemented")
}

func TestReformatDate(t *testing.T) {

	fmt.Errorf("Test not yet implemented")
}

func TestReformatDuration(t *testing.T) {

	fmt.Errorf("Test not yet implemented")
}

func TestParseProcesses(t *testing.T) {

	fmt.Errorf("Test not yet implemented")
}

// TODO: Change the parse durations to strings in the test examples in the sorting algorithms. for example 25h2m3s should be 1 day 1 hourr 2 minutes 3 seconds
