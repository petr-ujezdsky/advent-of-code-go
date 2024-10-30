package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"io"
)

type World struct {
	Program []int
}

type ColorAndRotation struct {
	Color, Rotation int
}

func paintHull(initialColor int, program []int) map[utils.Vector2i]int {
	input := make(chan int, 1)
	output := make(chan int)
	halt := make(chan int)

	defer close(input)
	defer close(output)
	defer close(halt)

	computer := common.NewIntCodeComputer("Unknown", program, input, output, halt)

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

		// initial color
		hull[position] = initialColor

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

	return paintedHull
}

func DoWithInputPart01(world World) int {
	paintedHull := paintHull(0, world.Program)

	printHull(paintedHull)

	return len(paintedHull)
}

func DoWithInputPart02(world World) string {
	paintedHull := paintHull(1, world.Program)

	printHull(paintedHull)

	return toStringHull(paintedHull)
}

func toStringHull(hull map[utils.Vector2i]int) string {
	bb := utils.NewBoundingRectangle(utils.Vector2i{})

	for pos := range hull {
		bb = bb.Enlarge(pos)
	}

	if min(bb.Width(), bb.Height()) > 300 {
		panic("Too large to print")
	}

	origin := utils.Vector2i{
		X: bb.Horizontal.Low,
		Y: bb.Vertical.Low,
	}

	m := matrix.NewMatrix[int](bb.Width(), bb.Height())

	// print values
	for pos, color := range hull {
		m.SetV(pos.Subtract(origin), color)
	}

	m = m.FlipVertical()

	return matrix.StringFmtSeparatorIndexedOrigin(m, 0, origin, "", matrix.NonIndexedAdapter(matrix.FmtBooleanConst[int](" ", "#")))
}

func printHull(hull map[utils.Vector2i]int) {
	fmt.Println(toStringHull(hull))
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
