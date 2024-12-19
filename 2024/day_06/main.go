package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Matrix matrix.Matrix[rune]
	Start  utils.Vector2i
}

type PathNode struct {
	Position  utils.Vector2i
	Direction utils.Direction4
	Next      *PathNode
}

var RotationsOrder = []int{0, 1, 2, 3, -1}

func walk(matrix matrix.Matrix[rune], startPos utils.Vector2i, startDir utils.Direction4) (map[utils.Vector2i]collections.BitSet8, *PathNode, bool) {
	dir := startDir
	pos := startPos
	node := &PathNode{
		Position:  utils.Vector2i{},
		Direction: 0,
		Next:      nil,
	}
	startNode := node

	visited := make(map[utils.Vector2i]collections.BitSet8)

	for {
		dirs, ok := visited[pos]
		if ok && dirs.Contains(int(dir)) {
			// found loop
			return visited, startNode, true
		}

		dirs.Push(int(dir))
		visited[pos] = dirs

		node.Next = &PathNode{
			Position:  pos,
			Direction: dir,
			Next:      nil,
		}

		node = node.Next
		for _, rotations := range RotationsOrder {
			if rotations == -1 {
				panic("Unable to find rotation")
			}

			newDir := dir.Rotate(-rotations)
			newPos := pos.Add(newDir.ToStep())

			char, ok := matrix.GetVSafe(newPos)
			if !ok {
				return visited, startNode, false
			}

			if char == '.' {
				if dir != newDir {
					dirs.Push(int(newDir))
					visited[pos] = dirs
				}
				dir = newDir
				pos = newPos
				break
			}

			if char != '#' && char != 'O' {
				panic("Unknown char " + string(char))
			}
		}

	}
}

func DoWithInputPart01(world World) int {
	visited, _, _ := walk(world.Matrix, world.Start, utils.Down)

	return len(visited)
}

func createMatrixFormatter(visited map[utils.Vector2i]collections.BitSet8, start utils.Vector2i) func(value rune, x int, y int) string {
	return func(value rune, x, y int) string {
		pos := utils.Vector2i{X: x, Y: y}

		if pos == start {
			return "^"
		}

		if dirs, ok := visited[pos]; ok {
			vertical := dirs.Contains(int(utils.Up)) || dirs.Contains(int(utils.Down))
			horizontal := dirs.Contains(int(utils.Left)) || dirs.Contains(int(utils.Right))

			if vertical && horizontal {
				return "+"
			}

			if vertical {
				return "|"
			}

			if horizontal {
				return "-"
			}

			panic("Empty dirs")
		}

		return string(value)
	}
}

func DoWithInputPart02(world World) int {
	visitedOrig, pathNode, _ := walk(world.Matrix, world.Start, utils.Down)
	fmt.Println(matrix.StringFmtSeparatorIndexed(world.Matrix, true, "", createMatrixFormatter(visitedOrig, world.Start)))

	obstructions := make(map[utils.Vector2i]struct{})

	// to the guard position
	pathNode = pathNode.Next

	loopedCount := 0
	for pathNode.Next != nil {
		_, ok := obstructions[pathNode.Next.Position]
		if !ok {
			world.Matrix.SetV(pathNode.Next.Position, 'O')
			//visited, _, looped := walk(world.Matrix, pathNode.Position, pathNode.Direction)
			//visited, _, looped := walk(world.Matrix, world.Start, utils.Down)

			_, _, looped := walk(world.Matrix, pathNode.Position, pathNode.Direction)
			//_, _, looped := walk(world.Matrix, world.Start, utils.Down)

			if looped {
				loopedCount++

				//fmt.Println(matrix.StringFmtSeparatorIndexed(world.Matrix, true, "", createMatrixFormatter(visited, world.Start)))
				//break
			}

			world.Matrix.SetV(pathNode.Next.Position, '.')

			obstructions[pathNode.Next.Position] = struct{}{}
		}
		pathNode = pathNode.Next
	}

	return loopedCount
}

func ParseInput(r io.Reader) World {
	start := utils.Vector2i{}

	mapper := func(char rune, x, y int) rune {
		if char == '^' {
			start = utils.Vector2i{X: x, Y: y}
			return '.'
		}

		return char
	}

	return World{
		Matrix: parsers.ParseToMatrixIndexed(r, mapper),
		Start:  start,
	}
}
