package main

import (
	"fmt"
	"os"
)

func ProcessMapToCSV() string {

	return ""
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
