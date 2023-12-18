package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Matrix2i = matrix.MatrixInt

type World struct {
	Tiles Matrix2i
}

func DoWithInputPart01(world World) int {
	tiles := world.Tiles
	_, _, totalHeatLoss, ok := FindMinHeatLossPath(tiles)
	if !ok {
		panic("No path found")
	}

	return totalHeatLoss
}

func h(endPos utils.Vector2i) func(utils.Vector2i) int {
	return func(pos utils.Vector2i) int {
		// manhattan distance
		return utils.Abs(pos.X-endPos.X) + utils.Abs(pos.Y-endPos.Y)
	}
}

func d(m Matrix2i) func(utils.Vector2i, utils.Vector2i) int {
	return func(nodeFrom, nodeTo utils.Vector2i) int {
		// step heat loss is the heat loss of target node
		return m.GetV(nodeTo)
	}
}

func findForbiddenPositions(pathIterator *alg.PathIterator[utils.Vector2i]) []utils.Vector2i {
	current, ok := pathIterator.Next()
	if !ok {
		panic("First tile should be the current tile")
	}

	previous, ok := pathIterator.Next()
	if !ok {
		// has no previous tile
		return nil
	}

	previous2, ok := pathIterator.Next()
	if !ok {
		// has only 2 previous tiles
		return []utils.Vector2i{previous}
	}

	if current.X == previous.X && previous.X == previous2.X || current.Y == previous.Y && previous.Y == previous2.Y {
		// all 3 are in a row
		next := current.Add(current.Subtract(previous))
		return []utils.Vector2i{previous, next}
	}

	return []utils.Vector2i{previous}
}

func neighbours(m Matrix2i) func(origin utils.Vector2i, pathIterator *alg.PathIterator[utils.Vector2i]) []utils.Vector2i {
	return func(origin utils.Vector2i, pathIterator *alg.PathIterator[utils.Vector2i]) []utils.Vector2i {
		var neighbours []utils.Vector2i

		// find forbidden positions based on the path
		forbiddenPositions := findForbiddenPositions(pathIterator)

		for _, dir := range utils.Direction4Steps {
			nextPos := origin.Add(dir)

			// check world validity
			if _, ok := m.GetVSafe(nextPos); !ok {
				continue
			}

			// check for forbidden positions
			forbidden := false
			for _, forbiddenPosition := range forbiddenPositions {
				if nextPos == forbiddenPosition {
					forbidden = true
					break
				}
			}

			if forbidden {
				continue
			}

			// everything is OK
			neighbours = append(neighbours, nextPos)
		}

		return neighbours
	}
}

func FindMinHeatLossPath(m Matrix2i) ([]utils.Vector2i, map[utils.Vector2i]int, int, bool) {
	endPos := utils.Vector2i{X: m.Width - 1, Y: m.Height - 1}
	return alg.AStar(utils.Vector2i{}, endPos, h(endPos), d(m), neighbours(m))
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	return World{Tiles: parsers.ParseToMatrixInt(r)}
}
