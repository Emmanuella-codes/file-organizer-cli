package main

import (
	"log"
	"file-organizer-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
