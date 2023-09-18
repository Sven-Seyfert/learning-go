package main

import "fmt"

func catchit() {
	if err := recover(); err != nil {
		fmt.Println("error message:", err)
	} else {
		fmt.Println("All good.")
	}
}

func main() {
	defer catchit()

	errorMessage := "Arrrrgh!"
	panic(errorMessage)
}
