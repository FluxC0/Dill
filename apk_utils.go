// this is the file the Alpine Linux package manager, APK, not to be confused with the Android File, also known as APK.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type APK_struct struct {
	PriorVS string `json:prior`
	PostVS  string `json:post`
}

func alpine_run() {
	cmd := exec.Command("apk", "version")

	output, err := cmd.CombinedOutput()
	check(err)
	lines := strings.Split(string(output), "\n")
	lines = append(lines[:0], lines[0+1:]...)

	var apk_ups []APK_struct

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 3 {
			priorvs := parts[0]
			postvs := parts[2]

			apk_ups = append(apk_ups, APK_struct{
				PriorVS: priorvs,
				PostVS:  postvs,
			})
		}
	}

	var jsonData []byte
	var errJSON error
	if len(apk_ups) == 0 {
		jsonData = []byte("null")
	} else {
		jsonData, errJSON = json.MarshalIndent(apk_ups, "", "  ")
	}
	check(errJSON)

	apkPath := getTMP("apk_updates.json")
	err = os.WriteFile(apkPath, jsonData, 0644)
	check(err)
}

func alpine_list() {
	LoadingSpinner(alpine_run)
	apkPath := getTMP("apk_updates.json")

	file, err := os.Open(apkPath)
	check(err)
	defer file.Close()
	var apkdata []APK_struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&apkdata)
	check(err)
	versionWidth := 0
	maxLength := 0
	for _, item := range apkdata {
		if len(item.PriorVS) > maxLength {
			maxLength = len(item.PriorVS)
		}
		if len(item.PostVS) > versionWidth {
			versionWidth = len(item.PostVS)
		}
	}
	nameWidth := maxLength + 2
	verticalLine := "|"
	for _, item := range apkdata {
		fmt.Printf(" %s %-*s %s %-*s %s\n", verticalLine, nameWidth, item.PriorVS, verticalLine, versionWidth, item.PostVS, verticalLine)
	}
}

func alpine_update() {
	cmd := exec.Command("apk", "upgrade", "--no-interactive")
	_, err := cmd.CombinedOutput()
	check(err)
}
