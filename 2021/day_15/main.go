package day_15

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type Matrix2i = utils.Matrix2i
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

func FindPathScore(m Matrix2i) int {
	currentScore := 0
	bestScore := math.MaxInt
	makeStepRecursive(Vector2i{0, 0}, m, currentScore, &bestScore)
	printStats(bestScore)
	return bestScore - m.Columns[0][0]
}

func calcBestScores(scoreTillEnd int, pos Vector2i, m, scores Matrix2i) {
	currentScore, ok := scores.GetVSafe(pos)

	if ok && scoreTillEnd < currentScore {
		// store better score
		scores.SetV(pos, scoreTillEnd)

		// propagate further
		riskLevel := m.Columns[pos.X][pos.Y]

		// left
		calcBestScores(scoreTillEnd+riskLevel, Vector2i{pos.X - 1, pos.Y}, m, scores)
		// up
		calcBestScores(scoreTillEnd+riskLevel, Vector2i{pos.X, pos.Y - 1}, m, scores)
		// right
		calcBestScores(scoreTillEnd+riskLevel, Vector2i{pos.X + 1, pos.Y}, m, scores)
		// down
		calcBestScores(scoreTillEnd+riskLevel, Vector2i{pos.X, pos.Y + 1}, m, scores)
	}
}

func CalcBestScore(m Matrix2i) (int, Matrix2i) {
	bestScores := utils.NewMatrix2[int](m.Width, m.Height)
	bestScores.SetAll(math.MaxInt)

	endPos := Vector2i{m.Width - 1, m.Height - 1}
	calcBestScores(0, endPos, m, bestScores)

	bestScore := bestScores.Columns[0][0]

	return bestScore, bestScores
}

func ParseInput(r io.Reader) (Matrix2i, error) {
	return utils.ParseToMatrix(r)
}
