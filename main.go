package main

import (
	"flag"
	"fmt"
  "os"
)
type Config struct {
  Package_Managers  string[]  `json:"package-managers"`
  
}
func parse() {
	isDangerous := flag.Bool("dangerous", false, "turns on danger mode if set, which bypasses all user checks. Tread lightly when using this flag.")
	flag.Parse()
  file, err := os.Open("~/.config/dill/config.json")
  if err != nil {
    fmt.Println("failed to open config. Exit.")
    return
  }
  defer file.Close()
	main_loop(*isDangerous)
}

func main_loop(isDangerous bool, config any) {
	fmt.Println("Welcome to Dill! The meta-package-manager.")
	if isDangerous {
		fmt.Println("Danger mode enabled! Here be dragons...")
	}
}

func main() {
	parse()
}
