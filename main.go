package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
)

type Config struct {
	Package_Managers []string `json:"packagemanagers"`
}

type Pac_Out struct {
	Package_name string `json:"package"`
	Version      string `json:"version"`
}

// LoadingSpinner starts a loading spinner while the provided task runs.

func parse() {
	isDangerous := flag.Bool("dangerous", false, "turns on danger mode if set, which bypasses all user checks. Tread lightly when using this flag.")
	flag.Parse()
	configPath := getConfigPath("config.json")
	file, err := os.Open(configPath)
	check(err)

	defer file.Close()
	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	check(err)

	main_loop(*isDangerous, config)
}

func main_loop(isDangerous bool, config Config) {
	fmt.Println("Welcome to Dill! The meta-package-manager.")
	if isDangerous {
		fmt.Println("Danger mode enabled! Here be dragons...")
	}
	if !isRoot() {
		fmt.Println("ERROR: Dill must be run as root.")
		panic("Permission Denied.")
	}
	if fileInfo, _ := os.Stdin.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {

		println("ERROR: do NOT run Dill in a non-interactive terminal.")
		panic("Dill is designed to be run interactively.")
	} else {
		fmt.Println("Running interactively, continuing...")
	}
	fmt.Println("checking package managers...")
	managers := config.Package_Managers
	if slices.Contains(managers, "pacman") {
		fmt.Println("pacman detected")
		pacman_list()
	} else if slices.Contains(managers, "apt") {
		fmt.Println("apt detected")
	} else if slices.Contains(managers, "dnf") {
		fmt.Println("dnf detected")
	} else if slices.Contains(managers, "flatpak") {
		fmt.Println("flatpak detected")
		flatpak_list()
	} else {
		fmt.Println("no package managers found in config.json. exiting...")
		os.Exit(1)
	}

	horizontalLine := "_"
	bottomLeft := "⎣"

	// Print the bottom line
	fmt.Printf("%s%s\n", bottomLeft, horizontalLine)
	confirm_choice()
	if slices.Contains(managers, "pacman") {
		LoadingSpinner(pac_update)
	}
}

func main() {
	parse()
}
