package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"io"
)

type Item struct {
	Type     string
	Position utils.Vector2i
}

type World struct {
	Items        map[utils.Vector2i]*Item
	Start        *Item
	Instructions []rune
}

func char2step(char rune) utils.Vector2i {
	switch char {
	case '^':
		return utils.Down.ToStep()
	case 'v':
		return utils.Up.ToStep()
	case '>':
		return utils.Right.ToStep()
	case '<':
		return utils.Left.ToStep()
	}

	panic("Unknown char")
}

func move(position utils.Vector2i, item *Item, items map[utils.Vector2i]*Item) {
	delete(items, item.Position)
	item.Position = position
	items[position] = item
}

func tryMove(step utils.Vector2i, item *Item, items map[utils.Vector2i]*Item) bool {
	if item.Type == "#" {
		// walls can not move
		return false
	}

	nextPos := item.Position.Add(step)

	next, ok := items[nextPos]

	if !ok || tryMove(step, next, items) {
		move(nextPos, item, items)
		return true
	}

	return false
}

func printItems(items map[utils.Vector2i]*Item) {
	mx := utils.Vector2i{}

	// find max boundary
	for pos := range items {
		mx = utils.Vector2i{X: utils.Max(mx.X, pos.X), Y: utils.Max(mx.Y, pos.Y)}
	}

	m := matrix.NewMatrix[string](mx.X+1, mx.Y+1)
	for pos, item := range items {
		m.SetV(pos, item.Type)
	}

	formatter := func(char string, x, y int) string {
		if char == "" {
			return "."
		}
		return char
	}

	str := matrix.StringFmtSeparatorIndexed(m, true, "", formatter)

	fmt.Println(str)
}

func DoWithInputPart01(world World) int {
	items := world.Items
	start := world.Start

	printItems(items)

	for _, instruction := range world.Instructions {
		//fmt.Printf("Instruction #%3d (%s)\n", i, string(instruction))

		step := char2step(instruction)

		tryMove(step, start, items)

		//printItems(items)
		//fmt.Println()
	}
	printItems(items)

	sum := 0
	for _, item := range items {
		if item.Type != "O" {
			continue
		}

		sum += 100*item.Position.Y + item.Position.X
	}

	return sum
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
				Type:     string(char),
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
