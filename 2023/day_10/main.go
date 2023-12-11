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

		dir = current.Next2[dir.Rotate(2)]
		if dir == -1 {
			// pipe does not continue
			return 0, false
		}

	}
}

func DoWithInputPart02(world World) int {
	return 0
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

				pipe.Next2[utils.Down] = utils.Up
				pipe.Next2[utils.Up] = utils.Down

				//pipe.NextA = pipes.Columns[x][y-1]
				//pipe.NextB = pipes.Columns[x][y+1]
			case '-':
				//pipe.Next[utils.Left] = pipes.Columns[x+1][y]
				//pipe.Next[utils.Right] = pipes.Columns[x-1][y]

				pipe.Next2[utils.Left] = utils.Right
				pipe.Next2[utils.Right] = utils.Left

				//pipe.NextA = pipes.Columns[x+1][y]
				//pipe.NextB = pipes.Columns[x-1][y]
			case 'L':
				//pipe.Next[utils.Up] = pipes.Columns[x+1][y]
				//pipe.Next[utils.Right] = pipes.Columns[x][y-1]

				pipe.Next2[utils.Up] = utils.Right
				pipe.Next2[utils.Right] = utils.Up

				//pipe.NextA = pipes.Columns[x+1][y]
				//pipe.NextB = pipes.Columns[x][y-1]
			case 'J':
				//pipe.Next[utils.Up] = pipes.Columns[x-1][y]
				//pipe.Next[utils.Left] = pipes.Columns[x][y-1]

				pipe.Next2[utils.Up] = utils.Left
				pipe.Next2[utils.Left] = utils.Up

				//pipe.NextA = pipes.Columns[x-1][y]
				//pipe.NextB = pipes.Columns[x][y-1]
			case '7':
				//pipe.Next[utils.Down] = pipes.Columns[x-1][y]
				//pipe.Next[utils.Left] = pipes.Columns[x][y+1]

				pipe.Next2[utils.Down] = utils.Left
				pipe.Next2[utils.Left] = utils.Down

				//pipe.NextA = pipes.Columns[x-1][y]
				//pipe.NextB = pipes.Columns[x][y+1]
			case 'F':
				//pipe.Next[utils.Down] = pipes.Columns[x+1][y]
				//pipe.Next[utils.Right] = pipes.Columns[x][y+1]

				pipe.Next2[utils.Down] = utils.Right
				pipe.Next2[utils.Right] = utils.Down

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
