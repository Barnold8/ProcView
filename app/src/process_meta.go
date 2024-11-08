package main

import (
	"fmt"
	"time"
)

type Process struct {
	name       string
	PID        int
	time_alive time.Duration // if theres a time struct, use this
}

func ParseProcesses(str string) []Process {

	var processes []Process

	duration, err := time.ParseDuration("2h30m")

	processes = append(processes, Process{
		name:       "example_process",
		PID:        1234,
		time_alive: duration,
	})

	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return nil
	}

	return processes
}
