package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Item struct {
	Type     rune
	Position utils.Vector2i
}

type World struct {
	Items        map[utils.Vector2i]*Item
	Start        *Item
	Instructions []rune
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	items := make(map[utils.Vector2i]*Item)
	var start *Item

	y := 0
	for scanner.Scan() && scanner.Text() != "" {
		for x, char := range scanner.Text() {
			if char == '.' {
				// skip empty space
				continue
			}

			pos := utils.Vector2i{X: x, Y: y}

			item := &Item{
				Type:     char,
				Position: pos,
			}

			items[pos] = item

			if char == '@' {
				start = item
			}
		}

		y++
	}

	var instructions []rune
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			instructions = append(instructions, char)
		}
	}

	return World{
		Items:        items,
		Start:        start,
		Instructions: instructions,
	}
}
