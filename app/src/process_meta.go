package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	name       string
	time_start time.Time     // This was needed for testing...
	time_alive time.Duration // if theres a time struct, use this
}

func parseTime(timeStr string) (time.Time, error) {

	var process_time time.Time
	var errors []error
	re := regexp.MustCompile(`(\d{4})(\d{2})(\d{2})(\d{2})(\d{2})(\d{2})\.\d+\+(\d{3})`)
	match := re.FindStringSubmatch(timeStr)

	if match != nil {
		year, yearErr := strconv.Atoi(match[1])
		month, monthErr := strconv.Atoi(match[2])
		day, dayErr := strconv.Atoi(match[3])
		hour, hourErr := strconv.Atoi(match[4])
		min, minErr := strconv.Atoi(match[5])
		second, secondErr := strconv.Atoi(match[6])

		errors = append(errors, yearErr, monthErr, dayErr, hourErr, minErr, secondErr)

		for _, err := range errors {
			if err != nil {
				return time.Time{}, err
			}
		}

		process_time = time.Date(
			year,
			time.Month(month),
			day,
			hour,
			min,
			second,
			0,
			time.UTC,
		)

	} else {
		return time.Time{}, fmt.Errorf("no date/time was found in the date string provided (%s)", timeStr)
	}

	return process_time, nil
}

func ParseProcesses(str string) []Process {

	var processes []Process
	var split []string = strings.Split(str, "\n")

	for _, element := range split {

		element = strings.TrimSpace(element)
		if strings.Contains(element, "exe") {

			var process Process
			var processed_string []string = strings.Fields(element)

			_time, err := parseTime(processed_string[1])
			if err != nil {
				return nil
			}

			process.name = processed_string[0]
			process.time_start = _time
			processes = append(processes, process)
		}
	}

	return processes
}

func removeDuplicateProcesses(unsortedProcesses []Process) []Process {

	processes := make([]Process, 0, len(unsortedProcesses)) // allocate memory for n processes that already exist so we dont go over the limit
	names := make(map[string]struct{})                      // make a map to look for keys only, we use a struct to fill the data field

	for _, process := range unsortedProcesses { // for the processes in our unsorted slice

		if _, exists := names[process.name]; !exists { // if the name doesnt exist
			processes = append(processes, process) // add it to the new array
			names[process.name] = struct{}{}       // add the name to the list of names
		}

	}

	return processes
}
