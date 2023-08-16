package main

import (
	"22-rock-paper-scissors-with-channels/game"
)

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()

	game.ClearScreen()
	game.PrintIntro()

	for {
		game.RoundChan <- 1 // increment
		<-game.RoundChan    // wait for the response of the channel

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1 // decrement
			<-game.RoundChan     // wait for the response of the channel
		}
	}

	game.PrintSummary()
}
