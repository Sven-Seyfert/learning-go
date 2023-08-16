package main

import (
	"06-hammer-bitcoin-game/game"
	"fmt"
)

func main() {
	gameRestart := true

	for gameRestart {
		game.Play()
		gameRestart = game.GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}
