package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Config struct {
	Package_Managers []string `json:"packagemanagers"`
	Authenticator    string   `json:"authenticator"`
	concurrency_tmp  string   `json:"concurrency"`
	Concurrency      bool
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

	config.Concurrency, _ = strconv.ParseBool(config.concurrency_tmp)

	main_loop(*isDangerous, config)
}

func main_loop(isDangerous bool, config Config) {
	fmt.Println("Welcome to Dill! The meta-package-manager.")
	if isDangerous {
		fmt.Println("Danger mode enabled! Here be dragons...")
	}
	fmt.Println("checking package managers...")
	managers := config.Package_Managers
	if slices.Contains(managers, "pacman") {
		fmt.Println("pacman detected")
		pac_run()
	} else if slices.Contains(managers, "apt") {
		fmt.Println("apt detected")
	} else if slices.Contains(managers, "dnf") {
		fmt.Println("dnf detected")
	} else {
		fmt.Println("no package managers found in config.json. exiting...")
		os.Exit(1)
	}

	horizontalLine := "_"
	bottomLeft := "⎣"
	flatpak_list()
	// Print the bottom line
	fmt.Printf("%s%s\n", bottomLeft, horizontalLine)
}

func main() {
	parse()
}
