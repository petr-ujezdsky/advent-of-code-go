package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Pipe struct {
	Char           rune
	Position       utils.Vector2i
	OutputForInput [4]*PipeOutput
}

type PipeOutput struct {
	OutputDirection utils.Direction4
	Lefts, Rights   []utils.Direction8
}

type World struct {
	Pipes    matrix.Matrix[*Pipe]
	Start    *Pipe
	StartPos utils.Vector2i
}

func DoWithInputPart01(world World) int {
	longest := -1
	for dir := 0; dir < 4; dir++ {
		if length, ok := walk(world, utils.Direction4(dir)); ok {
			longest = utils.Max(longest, length)
			fmt.Printf("Start to %v is OK, length %d\n", dir, length)
		}
	}

	return longest / 2
}

func walk(world World, dir utils.Direction4) (int, bool) {
	current, ok := world.Start, true
	pos := world.StartPos
	steps := 0

	for {
		step := dir.ToStep()
		step.Y = -step.Y

		pos = pos.Add(step)
		steps++

		current, ok = world.Pipes.GetVSafe(pos)
		if !ok {
			// out of bounds
			return -1, false
		}

		if current == world.Start {
			return steps, true
		}

		outputInfo := current.OutputForInput[dir]
		if outputInfo == nil {
			// pipe does not continue
			return 0, false
		}

		dir = outputInfo.OutputDirection
	}
}

func DoWithInputPart02(world World) int {
	for dir := 0; dir < 4; dir++ {
		if _, path, lefts, rights, ok := walkPath(world, utils.Direction4(dir)); ok {
			fmt.Printf("Lefts:\n%s\n\n", areaString(world.Pipes, lefts, path))
			fmt.Printf("Rights:\n%s\n", areaString(world.Pipes, rights, path))

			if area, ok := walkAreaTiles(path, lefts, world); ok {
				return area
			}

			if area, ok := walkAreaTiles(path, rights, world); ok {
				return area
			}

			panic("Found loop but no area")
		}
	}
	panic("No area found")
}

func areaString(pipes matrix.Matrix[*Pipe], area, path map[utils.Vector2i]struct{}) string {
	return pipes.StringFmtSeparator("", func(pipe *Pipe) string {
		if _, ok := path[pipe.Position]; ok {
			switch pipe.Char {
			case '|':
				return "│"
			case '-':
				return "─"
			case 'L':
				return "└"
			case 'J':
				return "┘"
			case '7':
				return "┐"
			case 'F':
				return "┌"
			case 'S':
				return "S"
			}
		}

		if _, ok := area[pipe.Position]; ok {
			return "."
		}

		return " "
	})
}

func walkPath(world World, dir utils.Direction4) (int, map[utils.Vector2i]struct{}, map[utils.Vector2i]struct{}, map[utils.Vector2i]struct{}, bool) {
	current, ok := world.Start, true
	var outputInfo *PipeOutput
	pos := world.StartPos
	steps := 0
	path := make(map[utils.Vector2i]struct{})
	lefts := make(map[utils.Vector2i]struct{})
	rights := make(map[utils.Vector2i]struct{})

	for {
		// on path
		path[pos] = struct{}{}

		// add neighbours
		if outputInfo != nil {
			// add lefts
			for _, left := range outputInfo.Lefts {
				lefts[pos.Add(left.ToStep().InvY())] = struct{}{}
			}

			// add rights
			for _, right := range outputInfo.Rights {
				rights[pos.Add(right.ToStep().InvY())] = struct{}{}
			}
		}

		step := dir.ToStep()
		step.Y = -step.Y

		pos = pos.Add(step)
		steps++

		current, ok = world.Pipes.GetVSafe(pos)
		if !ok {
			// out of bounds
			return -1, nil, nil, nil, false
		}

		if current == world.Start {
			// remove path from lefts/rights
			for pathPos := range path {
				delete(lefts, pathPos)
				delete(rights, pathPos)
			}

			return steps, path, lefts, rights, true
		}

		outputInfo = current.OutputForInput[dir]
		if outputInfo == nil {
			// pipe does not continue
			return 0, nil, nil, nil, false
		}
		dir = outputInfo.OutputDirection
	}
}

func walkAreaTiles(path, area map[utils.Vector2i]struct{}, world World) (int, bool) {
	visited := make(map[utils.Vector2i]struct{})

	for areaTile := range area {
		if !walkAreaTile(areaTile, path, visited, world) {
			return 0, false
		}
	}

	fmt.Printf("Found area:\n%s\n\n", areaString(world.Pipes, visited, path))

	return len(visited), true
}

func walkAreaTile(current utils.Vector2i, path, visited map[utils.Vector2i]struct{}, world World) bool {
	// stepped outside of world -> end totally - we are outside the loop
	if _, ok := world.Pipes.GetVSafe(current); !ok {
		return false
	}

	// stepped on path -> end
	if _, ok := path[current]; ok {
		return true
	}

	// stepped on already visited -> end
	if _, ok := visited[current]; ok {
		return true
	}

	// looks OK, add to visited
	visited[current] = struct{}{}

	// step on all directions
	for i := 0; i < 4; i++ {
		next := current.Add(utils.Direction4(i).ToStep())
		if !walkAreaTile(next, path, visited, world) {
			return false
		}
	}

	// all inspected, continue
	return true
}

func ParseInput(r io.Reader) World {
	// parse to matrix
	parseItem := func(char rune, x, y int) *Pipe {
		return &Pipe{
			Char:           char,
			Position:       utils.Vector2i{X: x, Y: y},
			OutputForInput: [4]*PipeOutput{},
		}
	}

	pipes := parsers.ParseToMatrixIndexed(r, parseItem)

	// join pipes
	var start *Pipe
	var startPos utils.Vector2i
	for x := 0; x < pipes.Width; x++ {
		for y := 0; y < pipes.Height; y++ {
			pos := utils.Vector2i{X: x, Y: y}

			pipe := pipes.GetV(pos)

			switch pipe.Char {
			case '|':
				pipe.OutputForInput[utils.Up] = &PipeOutput{
					OutputDirection: utils.Up,
					Lefts:           []utils.Direction8{utils.West},
					Rights:          []utils.Direction8{utils.East},
				}

				pipe.OutputForInput[utils.Down] = &PipeOutput{
					OutputDirection: utils.Down,
					Lefts:           []utils.Direction8{utils.East},
					Rights:          []utils.Direction8{utils.West},
				}
			case '-':
				pipe.OutputForInput[utils.Right] = &PipeOutput{
					OutputDirection: utils.Right,
					Lefts:           []utils.Direction8{utils.North},
					Rights:          []utils.Direction8{utils.South},
				}

				pipe.OutputForInput[utils.Left] = &PipeOutput{
					OutputDirection: utils.Left,
					Lefts:           []utils.Direction8{utils.South},
					Rights:          []utils.Direction8{utils.North},
				}
			case 'L':
				pipe.OutputForInput[utils.Down] = &PipeOutput{
					OutputDirection: utils.Right,
					Lefts:           []utils.Direction8{},
					Rights:          []utils.Direction8{utils.West, utils.SouthWest, utils.South},
				}

				pipe.OutputForInput[utils.Left] = &PipeOutput{
					OutputDirection: utils.Up,
					Lefts:           []utils.Direction8{utils.West, utils.SouthWest, utils.South},
					Rights:          []utils.Direction8{},
				}
			case 'J':
				pipe.OutputForInput[utils.Down] = &PipeOutput{
					OutputDirection: utils.Left,
					Lefts:           []utils.Direction8{utils.East, utils.SouthEast, utils.South},
					Rights:          []utils.Direction8{},
				}

				pipe.OutputForInput[utils.Right] = &PipeOutput{
					OutputDirection: utils.Up,
					Lefts:           []utils.Direction8{},
					Rights:          []utils.Direction8{utils.East, utils.SouthEast, utils.South},
				}
			case '7':
				pipe.OutputForInput[utils.Up] = &PipeOutput{
					OutputDirection: utils.Left,
					Lefts:           []utils.Direction8{},
					Rights:          []utils.Direction8{utils.East, utils.NorthEast, utils.North},
				}

				pipe.OutputForInput[utils.Right] = &PipeOutput{
					OutputDirection: utils.Down,
					Lefts:           []utils.Direction8{utils.East, utils.NorthEast, utils.North},
					Rights:          []utils.Direction8{},
				}
			case 'F':
				pipe.OutputForInput[utils.Up] = &PipeOutput{
					OutputDirection: utils.Right,
					Lefts:           []utils.Direction8{utils.West, utils.NorthWest, utils.North},
					Rights:          []utils.Direction8{},
				}

				pipe.OutputForInput[utils.Left] = &PipeOutput{
					OutputDirection: utils.Down,
					Lefts:           []utils.Direction8{},
					Rights:          []utils.Direction8{utils.West, utils.NorthWest, utils.North},
				}
			case '.':
			case 'S':
				start = pipe
				startPos = pos
			}
		}
	}

	return World{
		Pipes:    pipes,
		Start:    start,
		StartPos: startPos,
	}
}
