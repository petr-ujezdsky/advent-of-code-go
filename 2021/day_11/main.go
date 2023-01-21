package day_11

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

var neighbourDirs = []utils.Vector2i{
	// up
	{0, -1},
	// right
	{1, 0},
	// down
	{0, 1},
	// left
	{-1, 0},

	// diagonals
	{-1, -1},
	{1, -1},
	{1, 1},
	{-1, 1},
}

func raiseNeighboursLevels(energyLevels matrix.MatrixInt, pos utils.Vector2i) int {
	flashesCount := 0

	for _, neighbourDir := range neighbourDirs {
		neighbourPos := pos.Add(neighbourDir)
		neighbourLevel, ok := energyLevels.GetSafe(neighbourPos.X, neighbourPos.Y)

		// neighbour exists and has not flashed this round
		if ok && neighbourLevel != 0 {
			// raise level
			energyLevels.Columns[neighbourPos.X][neighbourPos.Y] = neighbourLevel + 1

			// recursively inspect another flashes
			flashesCount += inspectFlash(energyLevels, neighbourPos)
		}
	}

	return flashesCount
}

func inspectFlash(energyLevels matrix.MatrixInt, pos utils.Vector2i) int {
	energyLevel := energyLevels.Columns[pos.X][pos.Y]

	// flashing
	if energyLevel > 9 {
		// reset
		energyLevels.Columns[pos.X][pos.Y] = 0

		// raise levels in neighbours
		return raiseNeighboursLevels(energyLevels, pos) + 1
	}

	return 0
}

func step(energyLevels matrix.MatrixInt) int {
	// increment levels by 1
	for x := 0; x < energyLevels.Width; x++ {
		for y := 0; y < energyLevels.Height; y++ {
			energyLevels.Columns[x][y] = energyLevels.Columns[x][y] + 1
		}
	}

	// detect flashes
	flashesCount := 0
	for x := 0; x < energyLevels.Width; x++ {
		for y := 0; y < energyLevels.Height; y++ {
			flashesCount += inspectFlash(energyLevels, utils.Vector2i{x, y})
		}
	}

	return flashesCount
}

func CountFlashes(energyLevels matrix.MatrixInt, stepsCount int) (int, int) {
	flashesCount := 0
	allFlashedStepNumber := -1

	for i := 0; true; i++ {
		// move one step and count flashes
		stepFlashesCount := step(energyLevels)

		// all flashed at once
		if stepFlashesCount == energyLevels.Width*energyLevels.Height {
			allFlashedStepNumber = i + 1
		}

		// aggregate flashes count for first stepsCount
		if i < stepsCount {
			flashesCount += stepFlashesCount
		}

		// all detected
		if i >= stepsCount && allFlashedStepNumber > 0 {
			break
		}
	}

	return flashesCount, allFlashedStepNumber
}

func ParseInput(r io.Reader) matrix.MatrixInt {
	return parsers.ParseToMatrixInt(r)
}
