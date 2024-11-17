package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type flatpak_struct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func flat_run() {
	// Run the flatpak remote-ls command to get the list of upgrades
	cmd := exec.Command("flatpak", "remote-ls", "--updates", "--columns=application,branch")
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
		if len(parts) == 2 {
			name := parts[0]

			newVersion := parts[1] // The new version

			updates = append(updates, flatpak_struct{
				Name:    name,
				Version: newVersion,
			})
		}
	}

	// Convert updates slice to JSON
	var jsonData []byte
	var errJSON error
	if len(updates) == 0 {
		jsonData = []byte("null")
	} else {
		jsonData, errJSON = json.MarshalIndent(updates, "", "  ")
	}
	if errJSON != nil {
		fmt.Println("Error converting to JSON:", errJSON)
		return
	}

	// Print the JSON output

	flatPath := getTMP("flatpak_updates.json")
	// Write the JSON data to a file
	err = os.WriteFile(flatPath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func flatpak_list() {
	LoadingSpinner(flat_run)
	// Get the json data so it can be manipulated later
	flatPath := getTMP("flatpak_updates.json")

	file, err := os.Open(flatPath)

	check(err)

	defer file.Close()

	var flatdata []flatpak_struct

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&flatdata)

	check(err)

	maxLength := 0
	for _, item := range flatdata {
		if len(item.Name) > maxLength {
			maxLength = len(item.Name)
		}
	}
	nameWidth := maxLength + 2 // Extra padding for formatting
	versionWidth := 10
	verticalLine := "|"
	for _, item := range flatdata {
		fmt.Printf("%s %-*s %s %-*s %s\n", verticalLine, nameWidth, item.Name, verticalLine, versionWidth, item.Version, verticalLine)
	}
}

func flat_update() {
	// Run the flatpak update command
	cmd := exec.Command("flatpak", "update", "-y")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	/* In this instance, dont need to check the output, just check errors. */
}
