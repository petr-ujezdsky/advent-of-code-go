package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Player struct {
	Deck collections.Queue[int]
}

type World struct {
	Player1, Player2 Player
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func parseGroup(lines []string, _ int) Player {
	player := Player{Deck: collections.NewQueue[int]()}

	for _, line := range lines[1:] {
		player.Deck.Push(utils.ParseInt(line))
	}

	return player
}

func ParseInput(r io.Reader) World {
	players := parsers.ParseToGroups(r, parseGroup)

	return World{
		Player1: players[0],
		Player2: players[1],
	}
}
