package main

import (
	"fmt"
	"os"
	"strings"
)

func ProcessMapToCSV(processes map[string]Process) string {

	var builder strings.Builder

	builder.WriteString("Process Name, Time Created, RunTime\n")

	for _, value := range processes {
		builder.WriteString(fmt.Sprintf("%s,%s,%s\n", value.name, value.time_start, value.time_alive))
	}

	return builder.String()
}

func SaveToFile(filename string, contents string) {

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Write data to the file
	_, err = file.WriteString(contents)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	file.Close()

}
