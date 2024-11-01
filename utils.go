package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		// If XDG_CONFIG_HOME is not set, use a default path
		configHome = filepath.Join(os.Getenv("HOME"), ".config")
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
