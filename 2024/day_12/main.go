package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type World struct {
	Matrix matrix.Matrix[rune]
}

type Region struct {
	Id              int
	Name            rune
	Area, Perimeter int
}

var steps = slices.Reverse(utils.Direction4Steps[:])

func DoWithInputPart01(world World) int {
	used := make(map[utils.Vector2i]struct{})
	var regions []*Region

	for x, column := range world.Matrix.Columns {
		for y := range column {
			pos := utils.Vector2i{X: x, Y: y}

			region := floodFill(pos, nil, used, world.Matrix)
			if region != nil {
				regions = append(regions, region)
			}
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.Area * region.Perimeter
	}

	return totalPrice
}

func floodFill(pos utils.Vector2i, region *Region, used map[utils.Vector2i]struct{}, m matrix.Matrix[rune]) *Region {
	if _, ok := used[pos]; ok {
		// already used
		return nil
	}

	used[pos] = struct{}{}

	value := m.GetV(pos)

	if region == nil {
		region = &Region{
			Id:        0,
			Name:      value,
			Area:      0,
			Perimeter: 0,
		}
	}

	fences := 0
	for _, step := range steps {
		neighbourPos := pos.Add(step)
		neighbourValue, ok := m.GetVSafe(neighbourPos)
		if !ok {
			fences++
			continue
		}

		if neighbourValue != value {
			fences++
		} else {
			floodFill(neighbourPos, region, used, m)
		}
	}

	region.Area += 1
	region.Perimeter += fences

	return region
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) rune {
		return char
	}

	return World{Matrix: parsers.ParseToMatrix(r, parseItem)}
}
