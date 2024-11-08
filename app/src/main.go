package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("wmic process get Caption, CreationDate")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running tasklist command: %v", err)
	}
	for _, element := range output {
		fmt.Printf("Running processes:\n%s", element)
	}

}
