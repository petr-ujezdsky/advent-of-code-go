package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type TileType int

const (
	Empty TileType = iota
	Movable
	Solid
)

type Tile struct {
	Char rune
	Type TileType
}

type World struct {
	Tiles matrix.Matrix[Tile]
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) Tile {
		var tileType TileType
		switch char {
		case '.':
			tileType = Empty
		case 'O':
			tileType = Movable
		case '#':
			tileType = Solid
		}

		return Tile{
			Char: char,
			Type: tileType,
		}
	}

	return World{Tiles: parsers.ParseToMatrix(r, parseItem)}
}
