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
			var temp []string = strings.Fields(element)

			t, err := parseTime(temp[1])
			if err != nil {

			}

			fmt.Println("Time: ", t.String())
		}
	}

	return processes
}
