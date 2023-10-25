package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Vector2Hex = utils.Vector2i

type Vectors2Hex = map[Vector2Hex]struct{}

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
	black := initFloor(world)

	return len(black)
}

func initFloor(world World) Vectors2Hex {
	black := Vectors2Hex{}

	for _, path := range world.InitialPaths {
		if _, ok := black[path]; ok {
			delete(black, path)
		} else {
			black[path] = struct{}{}
		}
	}

	return black
}

func DoWithInputPart02(world World, daysCount int) int {
	black := initFloor(world)

	for i := 0; i < daysCount; i++ {
		toBeWhite := findToBeWhite(black)
		toBeBlack := findToBeBlack(black)

		for vec := range toBeWhite {
			delete(black, vec)
		}

		for vec := range toBeBlack {
			black[vec] = struct{}{}
		}

		fmt.Printf("Day %d: %v\n", i+1, len(black))
	}

	return len(black)
}

func findToBeWhite(black Vectors2Hex) Vectors2Hex {
	toBeWhite := Vectors2Hex{}

	for vec := range black {
		blackNeighboursCount := countBlackNeighbours(vec, black, 3)

		if blackNeighboursCount == 0 || blackNeighboursCount > 2 {
			toBeWhite[vec] = struct{}{}
		}
	}

	return toBeWhite
}

func findToBeBlack(black Vectors2Hex) Vectors2Hex {
	toBeBlack := Vectors2Hex{}

	for vecBlack := range black {
		for _, dir := range dir2vec {
			neighbour := vecBlack.Add(*dir)

			if _, ok := black[neighbour]; ok {
				// skip black neighbour
				continue
			}

			// neighbour is white, inspect it's neighbours
			blackNeighboursCount := countBlackNeighbours(neighbour, black, 3)
			if blackNeighboursCount == 2 {
				toBeBlack[neighbour] = struct{}{}
			}
		}
	}

	return toBeBlack
}

func countBlackNeighbours(vec Vector2Hex, black Vectors2Hex, maxCount int) int {
	blackNeighboursCount := 0

	for _, dir := range dir2vec {
		neighbour := vec.Add(*dir)

		if _, ok := black[neighbour]; ok {
			blackNeighboursCount++
		}

		if blackNeighboursCount == maxCount {
			break
		}
	}

	return blackNeighboursCount
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
