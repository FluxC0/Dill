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
	configHome := getTMP("pacman_dry_run_output.json")
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
	check(err)

	// Write JSON to file
	err = os.WriteFile(configHome, jsonData, 0644)
	check(err)
}

func pacman_list() {
	LoadingSpinner(pac_run) // pac_run takes the output from a pacman -Syu and turns it into a .json file, that is **TEMPORARILY** put in .config (this will probably be the most permanent thing in the project.)
	/* I just want to make a short note here. I spent a solid hour attempting to debug this LoadingSpinner function, thinking something was wrong with it, but no.
	* Turns out, it's just good old go making unused variables into errors. so i spent ALL this time trying to debug my program, but turns out, it never compiled in the first place.
	* Just leaving this here as a warning. Here were dragons. */
	fmt.Println("test")
	pacPath := getTMP("pacman_dry_run_output.json")
	var pacout []Pac_Out
	file, err := os.Open(pacPath)
	check(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pacout)
	check(err)

	maxLength := 0
	for i, item := range pacout {
		if pacout[i].Version == "downloading..." {
			remove_pac(pacout, i)
		}
		if len(item.Package_name) > maxLength {
			maxLength = len(item.Package_name)
		}
	}

	// Define widths based on max length
	nameWidth := maxLength + 2 // Extra padding for formatting
	versionWidth := 10

	verticalLine := "│"
	fmt.Println("Pacman Packages ")

	// Print each item in pacout
	for _, item := range pacout {
		fmt.Printf("%s %-*s %s %-*s %s\n", verticalLine, nameWidth, item.Package_name, verticalLine, versionWidth, item.Version, verticalLine)
	}
}

func pac_update() {
	cmd := exec.Command("pacman", "-Syu", "--noconfirm")
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "!! Could not update pacman !! %s", err)
	}
}
