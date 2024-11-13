package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("wmic.exe", "process", "get", "Caption,CreationDate")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running tasklist command: %v", err)
	}

	unsorted := ParseProcesses(string(output))

	sorted := removeDuplicateProcesses(unsorted)

	fmt.Println(sorted)

}
