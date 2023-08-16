package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	name := readString("What is your name?")
	age := readInt("How old are you?")

	// string interpolation variants
	fmt.Println("Your name is", name, "and you are", age, "years old.")              // variant 1 - default but inefficient
	fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old.", name, age)) // variant 2 - not necessaryly because of variant 3
	fmt.Printf("Your name is %s and you are %d years old.\n", name, age)             // variant 3 - best way
	// fmt.Println("Your name is " + name + " and you are " + age + " years old.")   // variant 4 - invalid because of string and int concat
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
		logFatal(err)
	}

	userInput = strings.ReplaceAll(userInput, "\r\n", "") // for windows
	userInput = strings.ReplaceAll(userInput, "\n", "")   // for mac and linux

	return userInput
}

func readInt(s string) int {
	userInput, err := strconv.Atoi(readString(s))
	if err != nil {
		logFatal(err)
	}

	return userInput
}

func logFatal(err error) {
	pc, _, line, _ := runtime.Caller(1)

	log.Fatalf("[ERROR] %s (func: %s, line: %d)", err, runtime.FuncForPC(pc).Name(), line)
}
