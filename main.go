package main

import (
	"encoding/json"
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
	// var managers [1]string

	pacman_list()
	horizontalLine := "_"
	bottomLeft := "⎣"

	// Print the bottom line
	fmt.Printf("%s%s\n", bottomLeft, horizontalLine)
	flat_run()
}

func main() {
	parse()
}
