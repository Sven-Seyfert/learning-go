package main

import (
	"20-debugging-console-applications/game"

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
