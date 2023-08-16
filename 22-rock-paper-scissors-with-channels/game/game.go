package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	for {
		// use select to process input in channels
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 0

		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			g.DisplayChan <- ""
		}
	}
}

func (g *Game) ClearScreen() {
	if runtime.GOOS == "windows" {
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

func (g *Game) PrintIntro() {
	g.DisplayChan <- `Rock, Paper & Scissors
----------------------
Game is played for three rounds, and best of three wins the game. Good luck!`
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {

	g.DisplayChan <- fmt.Sprintf(`
Round %d
-------`, g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Print("Please enter rock, paper, or scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)

	playerValue := -1

	// variant 1, with if else ..
	// if playerChoice == "rock" {
	// 	playerValue = ROCK
	// } else if playerChoice == "paper" {
	// 	playerValue = PAPER
	// } else if playerChoice == "scissors" {
	// 	playerValue = SCISSORS
	// }

	// variant 2, with map ..
	value := map[string]int{
		"rock":     ROCK,
		"paper":    PAPER,
		"scissors": SCISSORS,
	}

	playerValue = value[playerChoice]

	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)

	// variant 3, with switch statement
	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
		<-g.DisplayChan
	default:
	}

	g.DisplayChan <- ""
	<-g.DisplayChan

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:
			g.DisplayChan <- "Invalid choice!"
			<-g.DisplayChan
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	// *** changed here
	g.DisplayChan <- fmt.Sprintf(`
Final Score
-----------
Player: %d/3, Computer %d/3`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChan

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins game!"
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Computer wins game!"
		<-g.DisplayChan
	}

	g.DisplayChan <- ""
	<-g.DisplayChan
}
