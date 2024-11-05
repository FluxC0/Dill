package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func containsCharacterSet(s, charset string) bool {
	for _, char := range charset {
		if !strings.ContainsRune(s, char) {
			return false
		}
	}
	return true
}

func LoadingSpinner(task func()) {
	done := make(chan struct{})

	go func() {
		// Define a set of Unicode spinner characters
		spinChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠶", "⠦", "⠇", "⠏"}
		for {
			select {
			case <-done:
				return
			default:
				for _, char := range spinChars {
					fmt.Printf("\r%s Loading...", char)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	// Execute the provided task in a separate goroutine
	go func() {
		task()
		close(done) // Signal the loader to stop once the task is done
	}()

	// Wait until the task is complete
	<-done                             // Block until the task is complete
	fmt.Println("\rDone!            ") // Clear line after done
}

func getConfigPath(target string) string {
	targetUser := os.Getenv("SUDO_USER")
	var configHome string
	if targetUser == "" {
		fmt.Println("ERROR: SUDO_USER empty. this could be because you are using an unsupported authenticator. Please use sudo.")
		panic("Could not get config.")
	} else {
		configHome = filepath.Join("/home/", targetUser, "/.config/")
	}

	if target == "" {
		return filepath.Join(configHome, "dill") // make this so you can just retrieve the directory, for writing.
	} else {
		return filepath.Join(configHome, "dill", target)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func UnmarshalJSON(filepath string, storageStruct any) any {
	file, err := os.Open(filepath)
	check(err)
	defer file.Close() // Ensure the file is closed after reading

	decoder := json.NewDecoder(file)
	err = decoder.Decode(storageStruct)
	check(err)
	fmt.Println(storageStruct)
	return storageStruct // Return the unmarshaled struct
}

func remove_pac(s []Pac_Out, i int) []Pac_Out {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}
	return currentUser.Username == "root"
}

func confirm_choice() {
	fmt.Printf("\n ")
	fmt.Println("Would you like to make these changes to your installation? Y/N")
	var input string
	fmt.Scanln(&input)
	if input == "Y" {
		fmt.Println("Okay, continuing.")
	} else if input == "N" {
		fmt.Println("Okay, exiting gracefully.")
		os.Exit(0)

	} else {
		fmt.Println("Invalid choice, try again.")
		confirm_choice()
	}
}
