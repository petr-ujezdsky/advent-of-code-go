package day_21

import "github.com/petr-ujezdsky/advent-of-code-go/utils"

type Player struct {
	Position int
	Score    int
}

func Play(p1, p2 Player, endScore int) (int, int, Player, Player) {
	players := []*Player{&p1, &p2}

	turn := 0
	for p1.Score < endScore && p2.Score < endScore {
		currentPlayer := players[turn%len(players)]

		moves := utils.SumNtoM(turn*3+1, turn*3+3)
		position := (currentPlayer.Position+moves-1)%10 + 1
		score := currentPlayer.Score + position

		currentPlayer.Position = position
		currentPlayer.Score = score

		turn++
	}

	dieRolls := turn * 3

	return dieRolls * utils.Min(p1.Score, p2.Score), dieRolls, p1, p2
}
