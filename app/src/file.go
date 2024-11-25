package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
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

func addProgramToStartup(executablePath string) error {

	sourceFile, err := os.Open(executablePath)
	destinationPath := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\ProcView.exe", getCurrentUser()) // Replace with the desired destination path

	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	err = destinationFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close destination file: %v", err)
	}

	return nil

}

func getCurrentUser() string {

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return "NIL"
	}

	sections := strings.Split(currentUser.Username, "\\")

	return sections[len(sections)-1]

}
