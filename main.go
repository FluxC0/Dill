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
	concurrency_tmp  string   `json:concurrency`
	Concurrency      bool
}

func parse() {
	isDangerous := flag.Bool("dangerous", false, "turns on danger mode if set, which bypasses all user checks. Tread lightly when using this flag.")
	flag.Parse()

	file, err := os.Open("/home/kengel/.config/dill/config.json")
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
}

func main() {
	parse()
}
