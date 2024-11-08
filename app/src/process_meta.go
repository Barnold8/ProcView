package main

import "time"

type Process struct {
	name       string
	PID        int
	time_alive time.Time // if theres a time struct, use this
}
