package main

import (
	"fmt"
	"log"
	"os/exec"
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

func ParseProcesses(str string) map[string]Process {

	var split []string = strings.Split(str, "\n")
	processes := make(map[string]Process)

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
			processes[process.name] = process

		}
	}

	return processes
}

func grabProcesses() []byte {

	cmd := exec.Command("wmic.exe", "process", "get", "Caption,CreationDate")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running tasklist command: %v", err)
	}
	return output
}

func UpdateProcesses(processes map[string]Process, now time.Time, current_processes string) map[string]Process {
	updated_processes := make(map[string]Process)
	grabbed := ParseProcesses(current_processes)

	for key := range processes {
		var elapsed_time time.Duration
		_, exists := grabbed[key]

		if exists {
			elapsed_time = now.Sub(grabbed[key].time_start)
			updated_processes[key] = Process{
				grabbed[key].name,
				grabbed[key].time_start,
				elapsed_time.Abs(),
			}

		}
	}

	return updated_processes
}
