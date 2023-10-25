package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Vector2Hex = utils.Vector2i

type World struct {
	InitialPaths []Vector2Hex
}

var dir2vec = map[string]*Vector2Hex{
	"e":  {X: 10, Y: 0},
	"se": {X: 5, Y: -10},
	"sw": {X: -5, Y: -10},
	"w":  {X: -10, Y: 0},
	"nw": {X: -5, Y: 10},
	"ne": {X: 5, Y: 10},
}

func DoWithInputPart01(world World) int {
	black := map[Vector2Hex]struct{}{}

	for _, path := range world.InitialPaths {
		if _, ok := black[path]; ok {
			delete(black, path)
		} else {
			black[path] = struct{}{}
		}
	}

	return len(black)
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseItem(str string) Vector2Hex {
	totalVector := Vector2Hex{}

	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		char1 := runes[i]

		var dir *Vector2Hex

		if i+1 < len(runes) {
			char2 := runes[i+1]
			dirName := string([]rune{char1, char2})
			dir = dir2vec[dirName]

			if dir != nil {
				i++
			}
		}

		if dir == nil {
			dir = dir2vec[string(char1)]
		}

		if dir == nil {
			panic(fmt.Sprintf("No direction for %v", char1))
		}

		totalVector = totalVector.Add(*dir)
	}

	return totalVector
}

func ParseInput(r io.Reader) World {
	items := parsers.ParseToObjects(r, ParseItem)
	return World{InitialPaths: items}
}
