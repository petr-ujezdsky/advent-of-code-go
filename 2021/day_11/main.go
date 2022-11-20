package day_11

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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

func raiseNeighboursLevels(energyLevels utils.Matrix2i, pos utils.Vector2i) int {
	flashesCount := 0

	for _, neighbourDir := range neighbourDirs {
		neighbourPos := pos.Add(neighbourDir)
		neighbourLevel, ok := energyLevels.GetSafe(neighbourPos.X, neighbourPos.Y)

		// neighbour exists and has not flashed this round
		if ok && neighbourLevel != 0 {
			// raise level
			energyLevels.Set(neighbourPos.X, neighbourPos.Y, neighbourLevel+1)

			// recursively inspect another flashes
			flashesCount += inspectFlash(energyLevels, neighbourPos)
		}
	}

	return flashesCount
}

func inspectFlash(energyLevels utils.Matrix2i, pos utils.Vector2i) int {
	energyLevel := energyLevels.Get(pos.X, pos.Y)

	// flashing
	if energyLevel > 9 {
		// reset
		energyLevels.Set(pos.X, pos.Y, 0)

		// raise levels in neighbours
		return raiseNeighboursLevels(energyLevels, pos) + 1
	}

	return 0
}

func step(energyLevels utils.Matrix2i) int {
	// increment levels by 1
	for x := 0; x < energyLevels.Width; x++ {
		for y := 0; y < energyLevels.Height; y++ {
			energyLevels.Set(x, y, energyLevels.Get(x, y)+1)
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

func CountFlashes(energyLevels utils.Matrix2i, stepsCount int) int {
	flashesCount := 0

	for i := 0; i < stepsCount; i++ {
		flashesCount += step(energyLevels)
	}

	return flashesCount
}

func ParseInput(r io.Reader) (utils.Matrix2i, error) {
	return utils.ParseToMatrix(r)
}
