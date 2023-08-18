package main

import (
	"08-structured-types/hotkey"
	"08-structured-types/logger"

	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")

	// ------------------------------------
	question("Do you own a dog? Type Y/N.")
	pressedKey := hotkey.GetPressedKey()
	value, err := ownsADog(pressedKey)
	if err != nil {
		logger.LogFatal(err)
	}

	user.OwnsADog = value
	// ------------------------------------

	fmt.Printf("Your name is %s and you are %d years old. Your favorite number is %.2f. You own a dog: %t\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)
}

func question(s string) {
	fmt.Println(s)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	question(s)
	prompt()

	userInput, err := reader.ReadString('\n')
	if err != nil {
		logger.LogFatal(err)
	}

	userInput = strings.ReplaceAll(userInput, "\r\n", "") // for windows
	userInput = strings.ReplaceAll(userInput, "\n", "")   // for mac and linux

	return userInput
}

func readInt(s string) int {
	userInput, err := strconv.Atoi(readString(s))
	if err != nil {
		logger.LogFatal(err)
	}

	return userInput
}

func readFloat(s string) float64 {
	userInput, err := strconv.ParseFloat(readString(s), 64)
	if err != nil {
		logger.LogFatal(err)
	}

	return userInput
}

func ownsADog(char rune) (bool, error) {
	var charMap = map[rune]bool{
		'y': true,
		'Y': true,
		'n': false,
		'N': false,
	}

	for key := range charMap {
		if char == key {
			return charMap[char], nil
		}
	}

	return false, errors.New("you didn't press y/n")
}
