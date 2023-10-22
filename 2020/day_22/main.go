package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Player struct {
	Name string
	Deck collections.Queue[int]
}

type World struct {
	Player1, Player2 *Player
}

func DoWithInputPart01(world World) int {
	player1, player2 := world.Player1, world.Player2
	roundsCount := 0

	for {
		winner, finished := playRound(player1, player2)
		roundsCount++

		if finished {
			fmt.Printf("Winner in %v rounds: %v, Deck: %v\n", roundsCount, winner.Name, winner.Deck.PeekAll())
			return countScore(winner)
		}
	}
}

func DoWithInputPart02(world World) int {
	return 0
}

func playRound(player1, player2 *Player) (*Player, bool) {
	card1, card2 := player1.Deck.Pop(), player2.Deck.Pop()

	if card1 > card2 {
		player1.Deck.Push(card1)
		player1.Deck.Push(card2)
	} else {
		player2.Deck.Push(card2)
		player2.Deck.Push(card1)
	}

	if player1.Deck.Empty() {
		return player2, true
	}

	if player2.Deck.Empty() {
		return player1, true
	}

	return nil, false
}

func countScore(player *Player) int {
	score := 0

	for i, val := range player.Deck.PeekAll() {
		score += val * (player.Deck.Len() - i)
	}

	return score
}

func parseGroup(lines []string, i int) Player {
	player := Player{
		Name: fmt.Sprintf("Player %v", i+1),
		Deck: collections.Queue[int]{},
	}

	for _, line := range lines[1:] {
		player.Deck.Push(utils.ParseInt(line))
	}

	return player
}

func ParseInput(r io.Reader) World {
	players := parsers.ParseToGroups(r, parseGroup)

	return World{
		Player1: &players[0],
		Player2: &players[1],
	}
}
