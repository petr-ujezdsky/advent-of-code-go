package day_15

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
)

type Matrix2i = matrix.MatrixInt
type Vector2i = utils.Vector2i

var iterNr = 0
var prunedCount = 0

func printStats(bestScore int) {
	fmt.Printf("Iter #%d, pruned %d (%f%%), best %d\n", iterNr, prunedCount, float64(100*prunedCount)/float64(iterNr), bestScore)
}

func makeStepRecursive(pos Vector2i, m Matrix2i, currentScore int, bestScore *int) bool {
	riskLevel, ok := m.GetVSafe(pos)
	if !ok {
		return false
	}

	iterNr++
	if iterNr%70_000_000 == 0 {
		//if iterNr%100 == 0 {
		printStats(*bestScore)
	}

	currentScore += riskLevel

	// prune on worse score
	if currentScore > *bestScore {
		prunedCount++
		return false
	}

	okRight := makeStepRecursive(Vector2i{pos.X + 1, pos.Y}, m, currentScore, bestScore)
	okDown := makeStepRecursive(Vector2i{pos.X, pos.Y + 1}, m, currentScore, bestScore)

	if okRight || okDown {
		return true
	}

	endPos := Vector2i{m.Width - 1, m.Height - 1}
	if pos == endPos {
		// at the end
		*bestScore = currentScore
		return true
	}

	return false
}

var dirs = []Vector2i{
	// left
	{-1, 0},
	// up
	{0, -1},
	// right
	{1, 0},
	// down
	{0, 1},
}

func FindPathScore(m Matrix2i) int {
	currentScore := 0
	bestScore := math.MaxInt
	makeStepRecursive(Vector2i{0, 0}, m, currentScore, &bestScore)
	printStats(bestScore)
	return bestScore - m.Columns[0][0]
}

func calcBestScores(scoreTillEnd int, pos Vector2i, m, scores Matrix2i) {
	// store better score
	scores.SetV(pos, scoreTillEnd)

	// propagate further
	riskLevel := m.GetV(pos)
	score := scoreTillEnd + riskLevel

	for _, dir := range dirs {
		nextPos := pos.Add(dir)

		nextScore, ok := scores.GetVSafe(nextPos)
		if ok && score < nextScore {
			calcBestScores(score, nextPos, m, scores)
		}
	}
}

func CalcBestScore(m Matrix2i) (int, Matrix2i) {
	bestScores := matrix.NewMatrixNumber[int](m.Width, m.Height)
	bestScores.SetAll(math.MaxInt)

	endPos := Vector2i{m.Width - 1, m.Height - 1}
	calcBestScores(0, endPos, m, bestScores)

	bestScore := bestScores.Columns[0][0]

	return bestScore, bestScores
}

func h(endPos Vector2i) func(Vector2i) int {
	return func(pos Vector2i) int {
		// manhattan distance
		return utils.Abs(pos.X-endPos.X) + utils.Abs(pos.Y-endPos.Y)
	}
}

func d(m Matrix2i) func(Vector2i, Vector2i) int {
	return func(nodeFrom, nodeTo Vector2i) int {
		// distance or price is the risk level of target node
		return m.GetV(nodeTo)
	}
}

func neighbours(m Matrix2i) func(origin Vector2i) []Vector2i {
	return func(origin Vector2i) []Vector2i {
		var neighbours []Vector2i
		for _, dir := range dirs {
			nextPos := origin.Add(dir)

			// check validity
			_, ok := m.GetVSafe(nextPos)
			if ok {
				neighbours = append(neighbours, nextPos)
			}
		}

		return neighbours
	}
}

func CalcBestScoreAStar(m Matrix2i) ([]Vector2i, map[Vector2i]int, int, bool) {
	endPos := Vector2i{m.Width - 1, m.Height - 1}
	return alg.AStar(Vector2i{0, 0}, endPos, h(endPos), d(m), neighbours(m))
}

func EnlargeWorld(m Matrix2i) Matrix2i {
	factor := 5
	enlarged := matrix.NewMatrixInt(m.Width*factor, m.Height*factor)

	for x, col := range enlarged.Columns {
		for y := range col {
			xx := x % m.Width
			yy := y % m.Height

			xi := x / m.Width
			yi := y / m.Height

			valueToAdd := xi + yi

			enlarged.Columns[x][y] = (m.Columns[xx][yy]+valueToAdd-1)%9 + 1
		}
	}

	return enlarged
}

func ParseInput(r io.Reader) Matrix2i {
	return parsers.ParseToMatrixInt(r)
}
