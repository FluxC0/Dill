package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PackageUpdate struct {
	Package string `json:"package"`
	Version string `json:"version"`
}

func pac_run() {
	// Run the pacman -Syu command with --print and --print-format
	cmd := exec.Command("sudo", "pacman", "-Syu", "--print", "--print-format", "%n %v")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr // Capture any errors

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	var updates []PackageUpdate

	// Split the output into lines
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line into package name and version
		parts := strings.Fields(line)
		if len(parts) == 2 {
			updates = append(updates, PackageUpdate{Package: parts[0], Version: parts[1]})
		}
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(updates, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Write JSON to file
	err = os.WriteFile("pacman_dry_run_output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Conversion to JSON completed successfully!")
}
