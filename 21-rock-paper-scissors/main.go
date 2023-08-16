package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())

	playerChoice := ""
	playerValue := -1

	playerScore := 0
	computerScore := 0

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println()
		fmt.Println("Round", i)
		fmt.Println("-------")

		fmt.Print("Please enter rock, paper, scissors -> ")

		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)

		switch playerChoice {
		case "rock":
			playerValue = ROCK
		case "paper":
			playerValue = PAPER
		case "scissors":
			playerValue = SCISSORS
		default:
			// reset playerValue to -1 if invalid choice
			playerValue = -1
		}

		fmt.Println()
		fmt.Println("Player chose", strings.ToUpper(playerChoice))

		computerValue := rand.Intn(3)

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
		case PAPER:
			fmt.Println("Computer chose PAPER")
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
		default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")

			// decrement i to repeat the round
			i--
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
			case PAPER:
				if computerValue == SCISSORS {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
			case SCISSORS:
				if computerValue == ROCK {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
			default:
				fmt.Println("Invalid choice!")

				// decrement i to repeat the round
				i--
			}
		}
	}

	fmt.Println()
	fmt.Println("Final score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3", playerScore, computerScore)
	fmt.Println()

	if playerScore > computerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	}
}

func clearScreen() {
	// windows
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		return
	}

	// linux or mac
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func computerWins(score int) int {
	fmt.Println("Computer wins!")
	return score + 1
}

func playerWins(score int) int {
	fmt.Println("Player wins!")
	return score + 1
}
