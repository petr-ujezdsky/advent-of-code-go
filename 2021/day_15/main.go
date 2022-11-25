package day_15

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type Matrix2i = utils.Matrix2i
type Vector2i = utils.Vector2i

func makeStepRecursive(pos Vector2i, m Matrix2i) (int, bool) {
	riskLevel, ok := m.GetVSafe(pos)
	if !ok {
		return math.MaxInt, false
	}

	rightScore, okRight := makeStepRecursive(Vector2i{pos.X + 1, pos.Y}, m)
	downScore, okDown := makeStepRecursive(Vector2i{pos.X, pos.Y + 1}, m)

	if okRight || okDown {
		return utils.Min(rightScore, downScore) + riskLevel, true
	}

	// at the end
	return riskLevel, true
}

func FindPathScore(m Matrix2i) int {
	score, _ := makeStepRecursive(Vector2i{0, 0}, m)
	return score - m.Columns[0][0]
}

func ParseInput(r io.Reader) (Matrix2i, error) {
	return utils.ParseToMatrix(r)
}
