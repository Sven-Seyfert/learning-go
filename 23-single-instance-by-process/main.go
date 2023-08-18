package main

import (
	"23-single-instance-by-process/process"

	"fmt"
	"os"
)

func main() {
	// just a tryout of searching for a existing process
	processExists, err := process.ProcessExists("DevSpeedBar.exe")
	if err != nil {
		panic(err)
	}

	// if process exists, exit the program instance
	if processExists {
		os.Exit(0)
	}

	fmt.Println("test")
}
