package game

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func PlayRound(playerValue int) Round {
	computerValue := random(3) //nolint:gomnd
	computerChoice := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	var winMessage = []string{
		"Good job!",
		"Nice work!",
		"You should buy a lottery ticket.",
	}

	var loseMessage = []string{
		"Too bad!",
		"Try again",
		"This is just not your day.",
	}

	var drawMessage = []string{
		"Great minds think alike.",
		"Uh oh. Try again.",
		"Noboday wins, but you can try again.",
	}

	messageInt := random(3) //nolint:gomnd
	roundResult := ""       //nolint:wastedassign
	message := ""           //nolint:wastedassign

	switch playerValue {
	case computerValue:
		roundResult = "🎰 It's a draw!"
		message = fmt.Sprintf("📣 %s", drawMessage[messageInt])
	case (computerValue + 1) % 3: //nolint:gomnd
		roundResult = "🦸‍♂️ Player wins!"
		message = fmt.Sprintf("📣 %s", winMessage[messageInt])
	default:
		roundResult = "💻 Computer wins!"
		message = fmt.Sprintf("📣 %s", loseMessage[messageInt])
	}

	return Round{
		Message:        message,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}
}

func random(number int64) int {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(number))
	check(err)

	return int(bigInt.Int64())
}
