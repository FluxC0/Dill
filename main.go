package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Package_Managers []string `json:"packagemanagers"`
	Authenticator    string   `json:"authenticator"`
	concurrency_tmp  string   `json:concurrency`
	Concurrency      bool
}

func getConfigPath() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		// If XDG_CONFIG_HOME is not set, use a default path
		configHome = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return filepath.Join(configHome, "dill", "config.json")
}

func parse() {
	isDangerous := flag.Bool("dangerous", false, "turns on danger mode if set, which bypasses all user checks. Tread lightly when using this flag.")
	flag.Parse()
	configPath := getConfigPath()
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("failed to open config. Exit.", err)
		return
	}
	defer file.Close()

	var config Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("failed to decode config. Exit.", err)
		return
	}
	fmt.Println(config.Concurrency)
	config.Concurrency, _ = strconv.ParseBool(config.concurrency_tmp)
	fmt.Printf("%T \n", config.Concurrency)

	main_loop(*isDangerous, config)
}

func main_loop(isDangerous bool, config Config) {
	fmt.Println("Welcome to Dill! The meta-package-manager.")
	if isDangerous {
		fmt.Println("Danger mode enabled! Here be dragons...")
	}
	fmt.Println("checking package managers...")
	var managers [1]string
	managers[0] = "pacman"
	pac_run()
}

func main() {
	parse()
}
