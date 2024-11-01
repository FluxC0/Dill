package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type flatpak_struct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func flat_run() {
	// Run the flatpak remote-ls command to get the list of upgrades
	cmd := exec.Command("flatpak", "remote-ls", "--updates")
	output, err := cmd.CombinedOutput() // CombinedOutput to capture both stdout and stderr
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Convert output to string and split into lines
	lines := strings.Split(string(output), "\n")

	var updates []flatpak_struct

	// Parse each line for package name and version
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Example line: "org.example.Appname  1.2.3  1.2.4"
		parts := strings.Fields(line)
		if len(parts) == 3 {
			name := parts[0]
			newVersion := parts[2] // The new version

			updates = append(updates, flatpak_struct{
				Name:    name,
				Version: newVersion,
			})
		}
	}

	// Convert updates slice to JSON
	jsonData, err := json.MarshalIndent(updates, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print the JSON output
	fmt.Println(string(jsonData))
}
