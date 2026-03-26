package rps

import (
	"math/rand"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int
	ComputerChoice string
	RoundResult    string
}

func PlayRound(playervalue int) Round {
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	// lets see what computer chose.
	var result Round
	switch computerValue {
	case ROCK:
		computerChoice = "Computer Chose Rock"
	case PAPER:
		computerChoice = "Computer Chose Paper"
	case SCISSORS:
		computerChoice = "Computer Chose Scissors"
	default:
		return result
	}

	// check who won this game
	if playervalue == computerValue {
		roundResult = "Its a draw"
		winner = DRAW

	} else if playervalue == (computerValue+1)%3 {
		roundResult = "Player wins"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins"
		winner = COMPUTERWINS
	}
	result = Round{winner, computerChoice, roundResult}
	return result
}
