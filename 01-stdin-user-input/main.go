package main

import (
	"01-stdin-user-input/doctor"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	showTheIntroToConsole()
	userInteraction()
}

func showTheIntroToConsole() {
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}

func userInteraction() {
	reader := bufio.NewReader(os.Stdin) // for reading user input

	for {
		fmt.Print("-> ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		replaceAllOccurencesFlag := -1
		userInput = strings.Replace(userInput, "\r\n", "", replaceAllOccurencesFlag) // "\r\n" for windows
		userInput = strings.Replace(userInput, "\n", "", replaceAllOccurencesFlag)   // "\n" for mac os x and linux

		if userInput == "quit" {
			break
		}

		fmt.Println(doctor.Response(userInput))
	}
}
