package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"io"
	"regexp"
	"strconv"
)

type Item struct {
	Value  string
	Number int
}

type World struct {
	Items map[utils.Vector2i]Item
}

var regexAny = regexp.MustCompile(`([^\d.]|\d+)`)

func DoWithInputPart01(world World) int {
	sum := 0

	for pos, item := range world.Items {
		if item.Number == -1 {
			continue
		}

		neighbourPositions := findNeighbourPositions(pos, item)

		for _, neighbourPosition := range neighbourPositions {
			if neighbour, ok := world.Items[neighbourPosition]; ok && neighbour.Number == -1 {
				sum += item.Number
				break
			}
		}
	}

	return sum
}

func findNeighbourPositions(pos utils.Vector2i, item Item) []utils.Vector2i {
	var neighbours []utils.Vector2i

	// .xxxxx.
	// .12345.
	// .xxxxx.
	for i := range item.Value {
		neighbours = append(neighbours, utils.Vector2i{X: pos.X + i, Y: pos.Y - 1})
		neighbours = append(neighbours, utils.Vector2i{X: pos.X + i, Y: pos.Y + 1})
	}

	// x.....x
	// x12345x
	// x.....x
	for i := -1; i < 2; i++ {
		neighbours = append(neighbours, utils.Vector2i{X: pos.X - 1, Y: pos.Y + i})
		neighbours = append(neighbours, utils.Vector2i{X: pos.X + len(item.Value), Y: pos.Y + i})
	}

	return neighbours
}

func DoWithInputPart02(world World) int {
	sum := 0

	items := addPositionsPerDigit(world.Items)

	for pos, item := range items {
		if item.Value[0] != '*' {
			continue
		}

		uniqueNumbers := make(map[Item]struct{})
		for _, neighbourPosition := range findNeighbourPositions(pos, item) {
			if neighbour, ok := items[neighbourPosition]; ok && neighbour.Number != -1 {
				uniqueNumbers[neighbour] = struct{}{}
			}

			if len(uniqueNumbers) > 2 {
				break
			}
		}

		if len(uniqueNumbers) == 2 {
			product := 1

			for numberItem := range uniqueNumbers {
				product *= numberItem.Number
			}

			sum += product
		}
	}

	return sum
}

func addPositionsPerDigit(items map[utils.Vector2i]Item) map[utils.Vector2i]Item {
	enhanced := maps.Copy(items)
	for pos, item := range items {
		for i := 1; i < len(item.Value); i++ {
			enhanced[utils.Vector2i{X: pos.X + i, Y: pos.Y}] = item
		}
	}

	return enhanced
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	items := make(map[utils.Vector2i]Item)
	y := 0
	for scanner.Scan() {
		row := scanner.Text()
		values := regexAny.FindAllStringSubmatch(row, -1)
		indices := regexAny.FindAllStringSubmatchIndex(row, -1)

		for i, valueMatch := range values {
			if len(valueMatch) == 0 {
				continue
			}

			x := indices[i][0]
			value := valueMatch[1]
			number := -1
			if n, err := strconv.Atoi(value); err == nil {
				number = n
			}

			item := Item{
				Value:  value,
				Number: number,
			}

			pos := utils.Vector2i{X: x, Y: y}

			items[pos] = item
		}
		y++
	}

	return World{Items: items}
}
