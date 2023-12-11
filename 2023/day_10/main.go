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
	Char  rune
	Next  [4]*Pipe
	Next2 [4]utils.Direction4

	NextA, NextB *Pipe
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

		dir = current.Next2[dir]
		if dir == -1 {
			// pipe does not continue
			return 0, false
		}

	}
}

func DoWithInputPart02(world World) int {
	for dir := 0; dir < 4; dir++ {
		if _, path, lefts, rights, ok := walkPath(world, utils.Direction4(dir)); ok {
			fmt.Printf("Lefts:\n%s\n\n", areaString(world.Pipes, lefts))
			fmt.Printf("Rights:\n%s\n", areaString(world.Pipes, rights))

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

func areaString(pipes matrix.Matrix[*Pipe], area map[utils.Vector2i]struct{}) string {
	mPrint := pipes.Clone()

	for left := range area {
		mPrint.SetVSafe(left, &Pipe{Char: 'x'})
	}

	return mPrint.StringFmtSeparator("", func(pipe *Pipe) string {
		return string(pipe.Char)
	})
}

func walkPath(world World, dir utils.Direction4) (int, map[utils.Vector2i]struct{}, map[utils.Vector2i]struct{}, map[utils.Vector2i]struct{}, bool) {
	current, ok := world.Start, true
	pos := world.StartPos
	steps := 0
	path := make(map[utils.Vector2i]struct{})
	lefts := make(map[utils.Vector2i]struct{})
	rights := make(map[utils.Vector2i]struct{})

	for {
		// on path
		path[pos] = struct{}{}

		// check only straight lines - simplest
		if current.Char == '|' || current.Char == '-' {
			// on left
			leftPos := pos.Add(dir.Rotate(-1).ToStep().InvY())
			lefts[leftPos] = struct{}{}

			// on right
			rightPos := pos.Add(dir.Rotate(1).ToStep().InvY())
			rights[rightPos] = struct{}{}
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

		dir = current.Next2[dir]
		if dir == -1 {
			// pipe does not continue
			return 0, nil, nil, nil, false
		}

	}
}

func walkAreaTiles(path, area map[utils.Vector2i]struct{}, world World) (int, bool) {
	visited := make(map[utils.Vector2i]struct{})

	for areaTile := range area {
		if !walkAreaTile(areaTile, path, visited, world) {
			return 0, false
		}
	}

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
	parseItem := func(char rune) *Pipe {
		return &Pipe{
			Char:  char,
			Next:  [4]*Pipe{},
			Next2: [4]utils.Direction4{-1, -1, -1, -1},
		}
	}

	pipes := parsers.ParseToMatrix(r, parseItem)

	// join pipes
	var start *Pipe
	var startPos utils.Vector2i
	for x := 0; x < pipes.Width; x++ {
		for y := 0; y < pipes.Height; y++ {
			pos := utils.Vector2i{X: x, Y: y}

			pipe := pipes.GetV(pos)

			switch pipe.Char {
			case '|':
				//pipe.Next[utils.Down] = pipes.Columns[x][y-1]
				//pipe.Next[utils.Up] = pipes.Columns[x][y+1]

				pipe.Next2[utils.Up] = utils.Up
				pipe.Next2[utils.Down] = utils.Down

				//pipe.NextA = pipes.Columns[x][y-1]
				//pipe.NextB = pipes.Columns[x][y+1]
			case '-':
				//pipe.Next[utils.Left] = pipes.Columns[x+1][y]
				//pipe.Next[utils.Right] = pipes.Columns[x-1][y]

				pipe.Next2[utils.Right] = utils.Right
				pipe.Next2[utils.Left] = utils.Left

				//pipe.NextA = pipes.Columns[x+1][y]
				//pipe.NextB = pipes.Columns[x-1][y]
			case 'L':
				//pipe.Next[utils.Up] = pipes.Columns[x+1][y]
				//pipe.Next[utils.Right] = pipes.Columns[x][y-1]

				pipe.Next2[utils.Down] = utils.Right
				pipe.Next2[utils.Left] = utils.Up

				//pipe.NextA = pipes.Columns[x+1][y]
				//pipe.NextB = pipes.Columns[x][y-1]
			case 'J':
				//pipe.Next[utils.Up] = pipes.Columns[x-1][y]
				//pipe.Next[utils.Left] = pipes.Columns[x][y-1]

				pipe.Next2[utils.Down] = utils.Left
				pipe.Next2[utils.Right] = utils.Up

				//pipe.NextA = pipes.Columns[x-1][y]
				//pipe.NextB = pipes.Columns[x][y-1]
			case '7':
				//pipe.Next[utils.Down] = pipes.Columns[x-1][y]
				//pipe.Next[utils.Left] = pipes.Columns[x][y+1]

				pipe.Next2[utils.Up] = utils.Left
				pipe.Next2[utils.Right] = utils.Down

				//pipe.NextA = pipes.Columns[x-1][y]
				//pipe.NextB = pipes.Columns[x][y+1]
			case 'F':
				//pipe.Next[utils.Down] = pipes.Columns[x+1][y]
				//pipe.Next[utils.Right] = pipes.Columns[x][y+1]

				pipe.Next2[utils.Up] = utils.Right
				pipe.Next2[utils.Left] = utils.Down

				//pipe.NextA = pipes.Columns[x+1][y]
				//pipe.NextB = pipes.Columns[x][y+1]
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
