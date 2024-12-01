package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Process struct {
	name       string
	time_start time.Time     // This was needed for testing...
	time_alive time.Duration // if theres a time struct, use this
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
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

func reformatDate(input string) (string, error) {

	layout := "2006-01-02 15:04:05 -0700 UTC"

	t, err := time.Parse(layout, input)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02 15:04:05"), nil
}

func reformatDuration(input string) string {

	fmt.Printf("INCOMING TIME STRING: %s \n\n\n\n", input)

	if !strings.Contains(input, "s") {
		re := regexp.MustCompile(`(\d+(\.\d+)?)s`)
		input = re.ReplaceAllStringFunc(input, func(s string) string {

			parts := strings.Split(s, ".")
			return parts[0] + "s"
		})
	}

	fmt.Printf("PARSED TIME: %s \n\n\n\n", input)

	duration, err := time.ParseDuration(input)
	if err != nil {
		return fmt.Sprintf("Invalid time format, got %s\n with error %s", input, err)
	}

	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	var result []string
	if days > 0 {
		result = append(result, fmt.Sprintf("%d days", days))
	}
	if hours > 0 {
		result = append(result, fmt.Sprintf("%d hours", hours))
	}
	if minutes > 0 {
		result = append(result, fmt.Sprintf("%d minutes", minutes))
	}
	if seconds > 0 {
		result = append(result, fmt.Sprintf("%d seconds", seconds))
	}

	if len(result) == 0 {
		result = append(result, "0 seconds")
	}

	return strings.Join(result, " ")
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

			process.name = capitalizeFirstLetter(processed_string[0])
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

func extractKeys(processes map[string]Process) []string {

	keys := make([]string, 0, len(processes))

	for key := range processes {
		keys = append(keys, key)
	}

	return keys
}

// Process to string functions

func ProcessMapToStringSortedByName(processes map[string]Process, inverse bool) string {

	var builder strings.Builder
	keys := extractKeys(processes)
	sort.Strings(keys)

	if inverse {
		for i := len(keys) - 1; i >= 0; i-- {

			timeStart, timeErr := reformatDate(processes[keys[i]].time_start.String())
			timeAlive := reformatDuration(processes[keys[i]].time_alive.String())

			if timeErr != nil {
				timeStart = "Error!"
			}

			builder.WriteString(fmt.Sprintf("%s, %s, %s,\n", processes[keys[i]].name, timeStart, timeAlive))
		}

	} else {
		for _, key := range keys {

			timeStart, timeErr := reformatDate(processes[key].time_start.String())
			timeAlive := reformatDuration(processes[key].time_alive.String())

			if timeErr != nil {
				timeStart = "Error!"
			}

			builder.WriteString(fmt.Sprintf("%s, %s, %s,\n", processes[key].name, timeStart, timeAlive))
		}
	}

	return builder.String()
}

func ProcessMapToStringSortedByTimeStarted(processes map[string]Process, inverse bool) string {

	var builder strings.Builder
	keys := extractKeys(processes)

	sort.Slice(keys, func(i, j int) bool {
		return processes[keys[i]].time_start.Before(processes[keys[j]].time_start)
	})

	if inverse {
		for i := len(keys) - 1; i >= 0; i-- {

			fmt.Println()

			timeStart, timeErr := reformatDate(processes[keys[i]].time_start.String())
			timeAlive := reformatDuration(processes[keys[i]].time_alive.String())

			if timeErr != nil {
				timeStart = "Error!"
			}
			fmt.Println(timeStart)
			builder.WriteString(fmt.Sprintf("%s, %s, %s,\n", processes[keys[i]].name, timeStart, timeAlive))
		}

	} else {
		for _, key := range keys {

			fmt.Println(processes[key].time_start.String())

			timeStart, timeErr := reformatDate(processes[key].time_start.String())
			timeAlive := reformatDuration(processes[key].time_alive.String())

			if timeErr != nil {
				timeStart = "Error!"
			}

			builder.WriteString(fmt.Sprintf("%s, %s, %s,\n", processes[key].name, timeStart, timeAlive))
		}
	}

	return builder.String()
}

func ProcessMapToStringSortedByTimeAlive(processes map[string]Process, inverse bool) string {

	var builder strings.Builder
	keys := extractKeys(processes)

	sort.SliceStable(keys, func(i, j int) bool {
		return processes[keys[i]].time_alive < processes[keys[j]].time_alive
	})

	if inverse {
		for i := len(keys) - 1; i >= 0; i-- {

			timeStart, timeErr := reformatDate(processes[keys[i]].time_start.String())
			timeAlive := reformatDuration(processes[keys[i]].time_alive.String())

			if timeErr != nil {
				timeStart = "Error!"
			}

			builder.WriteString(fmt.Sprintf("%s, %s, %s,\n", processes[keys[i]].name, timeStart, timeAlive))
		}

	} else {
		for _, key := range keys {

			timeStart, timeErr := reformatDate(processes[key].time_start.String())
			timeAlive := reformatDuration(processes[key].time_alive.String())

			if timeErr != nil {
				timeStart = "Error!"
			}

			builder.WriteString(fmt.Sprintf("%s, %s, %s,\n", processes[key].name, timeStart, timeAlive))
		}
	}

	return builder.String()
}

// Process to string functions
