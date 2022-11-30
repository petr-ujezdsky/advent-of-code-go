package day_21

import "github.com/petr-ujezdsky/advent-of-code-go/utils"

type Player struct {
	Position, Score int
}

type Couple struct {
	CurrentPlayer, OtherPlayer Player
}

func NewCouple(currentPlayer, otherPlayer Player) Couple {
	return Couple{
		CurrentPlayer: currentPlayer,
		OtherPlayer:   otherPlayer,
	}
}

var dieRollOutcomes = dieRollOutcomesMap()

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

func dieRollOutcomesMapRecursive(outcomes map[int]int, height, sum int) {
	if height == 0 {
		outcomes[sum]++
		return
	}

	// rolled 1
	dieRollOutcomesMapRecursive(outcomes, height-1, sum+1)

	// rolled 2
	dieRollOutcomesMapRecursive(outcomes, height-1, sum+2)

	// rolled 3
	dieRollOutcomesMapRecursive(outcomes, height-1, sum+3)
}

// dieRollOutcomesMap creates map with
// key - sum from 3 rolls
// value - combinations count to get this sum
func dieRollOutcomesMap() map[int]int {
	// there will be 7 possible sums - from 3 (1+1+1) to 9 (3+3+3)
	outcomes := make(map[int]int, 7)
	dieRollOutcomesMapRecursive(outcomes, 3, 0)
	return outcomes
}

func PlayRecursive(currentPlayer, otherPlayer Player, depth int, cache map[Couple][]int) (int, int) {
	// check cache
	couple := NewCouple(currentPlayer, otherPlayer)
	counts, ok := cache[couple]
	if ok {
		// cache hit!
		return counts[0], counts[1]
	}

	totalCurrentWins, totalOtherWins := 0, 0

	for moves, count := range dieRollOutcomes {
		position := (currentPlayer.Position+moves-1)%10 + 1
		score := currentPlayer.Score + position

		// current player is winner
		if score >= 21 {
			totalCurrentWins += count
			continue
		}

		// update current player
		updatedPlayer := Player{
			Position: position,
			Score:    score,
		}

		// next turn
		otherWins, currentWins := PlayRecursive(otherPlayer, updatedPlayer, depth+1, cache)
		totalCurrentWins += count * currentWins
		totalOtherWins += count * otherWins
	}

	// cache results
	cache[couple] = []int{totalCurrentWins, totalOtherWins}

	return totalCurrentWins, totalOtherWins
}

func PlayFaster(p1, p2 Player) (int, int, int) {
	cache := make(map[Couple][]int)
	p1wins, p2wins := PlayRecursive(p1, p2, 0, cache)

	return p1wins, p2wins, utils.Max(p1wins, p2wins)
}
