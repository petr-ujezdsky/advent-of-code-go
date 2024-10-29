package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type World struct {
	Program []int
}

type ColorAndRotation struct {
	Color, Rotation int
}

func DoWithInputPart01(world World) int {
	input := make(chan int, 1)
	output := make(chan int)
	halt := make(chan int)

	defer close(input)
	defer close(output)
	defer close(halt)

	computer := common.NewIntCodeComputer("Unknown", world.Program, input, output, halt)

	go common.Run(computer)

	formattedOutput := make(chan ColorAndRotation)
	defer close(formattedOutput)

	// read 2 values from output and send them to formattedOutput channel
	go func() {
		for color := range output {
			rotation := <-output

			formattedOutput <- ColorAndRotation{
				Color:    color,
				Rotation: rotation,
			}
		}
	}()

	end := make(chan map[utils.Vector2i]int)
	defer close(end)

	go func() {
		hull := make(map[utils.Vector2i]int)

		// start at [0,0]
		position := utils.Vector2i{}
		direction := utils.Up

		for {
			color := hull[position]
			input <- color

			// wait for output or halt
			select {
			case o := <-formattedOutput:
				// paint current position
				hull[position] = o.Color

				// rotate and move
				direction = direction.Rotate(o.Rotation*2 - 1)
				position = position.Add(direction.ToStep())
			case <-halt:
				end <- hull
				return
			}
		}
	}()

	// wait for painting
	paintedHull := <-end

	return len(paintedHull)
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
