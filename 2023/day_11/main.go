package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type Item struct {
	Char     rune
	Star     bool
	Id       int
	Position utils.Vector2i
}

type World struct {
	Matrix        matrix.Matrix[Item]
	Stars         []Item
	StarsInColumn []int
	StarsInRow    []int
}

func DoWithInputPart01(world World) int {
	sum := 0

	for i, starFrom := range world.Stars {
		for _, starTo := range world.Stars[i+1:] {
			sum += calculateDistance(starFrom, starTo, world, 2)
		}
	}

	return sum
}

func calculateDistance(starFrom, starTo Item, world World, voidSize int) int {
	distance := starTo.Position.Subtract(starFrom.Position).LengthManhattan()

	for x := utils.Min(starFrom.Position.X, starTo.Position.X); x < utils.Max(starFrom.Position.X, starTo.Position.X); x++ {
		if world.StarsInColumn[x] == 0 {
			distance += voidSize - 1
		}
	}

	for y := utils.Min(starFrom.Position.Y, starTo.Position.Y); y < utils.Max(starFrom.Position.Y, starTo.Position.Y); y++ {
		if world.StarsInRow[y] == 0 {
			distance += voidSize - 1
		}
	}

	return distance
}

func DoWithInputPart02(world World, voidSize int) int {
	sum := 0

	for i, starFrom := range world.Stars {
		for _, starTo := range world.Stars[i+1:] {
			sum += calculateDistance(starFrom, starTo, world, voidSize)
		}
	}

	return sum
}

func ParseInput(r io.Reader) World {
	var stars []Item
	idSeq := 1

	parseItem := func(char rune, x, y int) Item {
		star := char == '#'

		item := Item{
			Char:     char,
			Star:     star,
			Position: utils.Vector2i{X: x, Y: y},
		}

		if star {
			item.Id = idSeq
			idSeq++
			stars = append(stars, item)
		}

		return item
	}

	m := parsers.ParseToMatrixIndexed(r, parseItem)

	starsInColumn := slices.Filled(0, m.Width)
	starsInRow := slices.Filled(0, m.Height)

	for _, star := range stars {
		starsInColumn[star.Position.X]++
		starsInRow[star.Position.Y]++
	}

	return World{
		Matrix:        m,
		Stars:         stars,
		StarsInColumn: starsInColumn,
		StarsInRow:    starsInRow,
	}
}
