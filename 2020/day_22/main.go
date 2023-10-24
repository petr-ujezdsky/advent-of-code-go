package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type Player struct {
	Name   string
	Number int
	Deck   collections.Queue[int]
}

func (p *Player) CopyAndTrunc(cardsCount int) *Player {
	return &Player{
		Name:   p.Name,
		Number: p.Number,
		Deck:   collections.NewQueueFilled(p.Deck.PeekAll()[0:cardsCount]),
	}
}

func (p *Player) Equal(player *Player) bool {
	if p.Deck.Len() != player.Deck.Len() {
		return false
	}

	return slices.Equal(p.Deck.PeekAll(), player.Deck.PeekAll())
}

func (p *Player) Clone() *Player {
	return &Player{
		Name:   p.Name,
		Number: p.Number,
		Deck:   p.Deck.Clone(),
	}
}

type World struct {
	Player1, Player2 *Player
}

func (w World) Clone() World {
	return World{
		Player1: w.Player1.Clone(),
		Player2: w.Player2.Clone(),
	}
}

type History struct {
	Worlds []World
}

func (h *History) Contains(world World) bool {
	for _, w := range h.Worlds {
		if w.Player1.Equal(world.Player1) && w.Player2.Equal(world.Player2) {
			return true
		}
	}

	return false
}

func (h *History) Add(world World) {
	h.Worlds = append(h.Worlds, world)
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

func DoWithInputPart02(world World) int {
	gamesCounter := 1
	winner := playGame(world, &gamesCounter)
	fmt.Printf("Total winner is %v, Deck: %v\n", winner.Name, winner.Deck.PeekAll())

	return countScore(winner)
}

func playGame(world World, gamesCounter *int) *Player {
	gameNumber := *gamesCounter
	*gamesCounter++
	//fmt.Printf("=== Game %v ===\n\n", gameNumber)

	roundsCounter := 1
	history := &History{}

	for {
		//fmt.Printf("-- Round %v (Game %v) --\n", roundsCounter, gameNumber)
		winner, finished := playRoundRecursive(world, history, roundsCounter, gameNumber, gamesCounter)

		if finished {
			//fmt.Printf("The winner of game %v is player %v!\n", gameNumber, winner.Name)
			return winner
		}

		roundsCounter++
	}
}

func playRoundRecursive(world World, history *History, roundNumber, gameNumber int, gamesCounter *int) (*Player, bool) {
	// check history first
	if history.Contains(world) {
		// there was a game with the same configuration -> player 1 wins
		return world.Player1, true
	}

	// it is new game -> add to history
	history.Add(world.Clone())

	player1, player2 := world.Player1, world.Player2

	// draw cards
	card1, card2 := player1.Deck.Pop(), player2.Deck.Pop()

	var winner *Player

	if card1 <= player1.Deck.Len() && card2 <= player2.Deck.Len() {
		// recursive game
		world := World{
			Player1: player1.CopyAndTrunc(card1),
			Player2: player2.CopyAndTrunc(card2),
		}

		subGameWinner := playGame(world, gamesCounter)

		if subGameWinner == world.Player1 {
			winner = player1
		} else {
			winner = player2
		}
	} else {
		// simply compare cards otherwise
		if card1 > card2 {
			winner = player1
		} else {
			winner = player2
		}
	}

	//fmt.Printf("Player %v wins round %v of game %v!\n\n", winner.Number, roundNumber, gameNumber)

	// move cards to the winner's deck
	if winner == player1 {
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
		Name:   fmt.Sprintf("Player %v", i+1),
		Number: i + 1,
		Deck:   collections.Queue[int]{},
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
