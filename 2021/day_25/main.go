package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type World struct {
	CucumbersRight map[Cucumber]struct{}
	CucumbersDown  map[Cucumber]struct{}
	Area           utils.BoundingRectangle
}

type Cucumber struct {
	Position  utils.Vector2i
	Direction utils.Direction4
}

func DoWithInput(world World) int {
	return len(world.CucumbersRight)
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	cucumbersRight := make(map[Cucumber]struct{})
	cucumbersDown := make(map[Cucumber]struct{})

	y := 0
	xMax := 0
	for scanner.Scan() {
		xMax = len(scanner.Text())

		for x, char := range scanner.Text() {
			if char == '.' {
				continue
			}

			switch char {
			case '>':
				cucumber := Cucumber{
					Position:  utils.Vector2i{X: x, Y: y},
					Direction: utils.Right,
				}
				cucumbersRight[cucumber] = struct{}{}
			case 'v':
				cucumber := Cucumber{
					Position:  utils.Vector2i{X: x, Y: y},
					Direction: utils.Up,
				}
				cucumbersDown[cucumber] = struct{}{}
			}
		}
		y++
	}

	area := utils.BoundingRectangle{
		Horizontal: utils.IntervalI{
			Low:  0,
			High: xMax - 1,
		},
		Vertical: utils.IntervalI{
			Low:  0,
			High: y - 1,
		},
	}

	return World{
		CucumbersRight: cucumbersRight,
		CucumbersDown:  cucumbersDown,
		Area:           area,
	}
}
