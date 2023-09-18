package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux.")
	case "darwin":
		fmt.Print("commercial ")
		fallthrough
	case "openbsd", "freebsd":
		fmt.Println("BSD-derivative.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
