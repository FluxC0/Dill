package main

import (
	"flag"
	"fmt"
	"os"
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
	UnmarshalJSON(configPath, &config)

	config.Concurrency, _ = strconv.ParseBool(config.concurrency_tmp)

	main_loop(*isDangerous, config)
}

func main_loop(isDangerous bool, config Config) {
	fmt.Println("Welcome to Dill! The meta-package-manager.")
	if isDangerous {
		fmt.Println("Danger mode enabled! Here be dragons...")
	}
	fmt.Println("checking package managers...")
	var managers [1]string

	LoadingSpinner(pac_run) // pac_run takes the output from a pacman -Syu and turns it into a .json file, that is **TEMPORARILY**
	pacPath := getConfigPath("pacman_dry_run_output.json")
	var pacout []Pac_Out
	UnmarshalJSON(pacPath, &pacout)

	horizontalLine := "─"
	verticalLine := "│"
	topLeft := "╔"
	topRight := "╗"
	bottomLeft := "╚"
	bottomRight := "╝"
}

func main() {
	parse()
}
