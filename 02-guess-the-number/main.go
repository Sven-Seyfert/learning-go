package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	randomNumbers := myNumbers{rand.Intn(8) + 2, rand.Intn(8) + 2, rand.Intn(8) + 2}
	answer := randomNumbers.firstNumber*randomNumbers.secondNumber - randomNumbers.substraction

	displayWelcomeInstructions()
	playTheGame(randomNumbers.firstNumber, randomNumbers.secondNumber, randomNumbers.substraction, answer)
}

type myNumbers struct {
	firstNumber  int
	secondNumber int
	substraction int
}

func displayWelcomeInstructions() {
	introMessage := introMessage()
	fmt.Println(introMessage, prompt)
}

func introMessage() string {
	return `Guess the Number Game
---------------------
Think of a number between 1 and 10`
}

func playTheGame(firstNumber, secondNumber, substraction, answer int) {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// game instructions
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", substraction, prompt)
	reader.ReadString('\n')

	// the answer
	fmt.Println("The answer is", answer)
}
