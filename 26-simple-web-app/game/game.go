package game

import (
	"math/rand"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2

	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := "" //nolint:wastedassign
	winner := 0       //nolint:wastedassign

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	switch playerValue {
	case computerValue:
		roundResult = "It's a draw"
		winner = DRAW
	case (computerValue + 1) % 3: //nolint:gomnd
		roundResult = "Player wins!"
		winner = PLAYERWINS
	default:
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
	}

	return Round{
		Winner:         winner,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
}
