package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type Plant struct {
	Name   rune
	Region *Region
}

type World struct {
	Matrix matrix.Matrix[*Plant]
}

type Region struct {
	Id              int
	Name            rune
	Area, Perimeter int
}

var steps = slices.Reverse(utils.Direction4Steps[:])

func findRegions(m matrix.Matrix[*Plant]) []*Region {
	used := make(map[utils.Vector2i]struct{})
	var regions []*Region

	for x, column := range m.Columns {
		for y := range column {
			pos := utils.Vector2i{X: x, Y: y}

			region := floodFill(pos, nil, used, m)
			if region != nil {
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func DoWithInputPart01(world World) int {
	regions := findRegions(world.Matrix)

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.Area * region.Perimeter
	}

	return totalPrice
}

func floodFill(pos utils.Vector2i, region *Region, used map[utils.Vector2i]struct{}, m matrix.Matrix[*Plant]) *Region {
	if _, ok := used[pos]; ok {
		// already used
		return nil
	}

	used[pos] = struct{}{}

	plant := m.GetV(pos)

	if region == nil {
		region = &Region{
			Id:        0,
			Name:      plant.Name,
			Area:      0,
			Perimeter: 0,
		}
	}

	plant.Region = region

	fences := 0
	for _, step := range steps {
		neighbourPos := pos.Add(step)
		neighbourPlant, ok := m.GetVSafe(neighbourPos)
		if !ok {
			fences++
			continue
		}

		if neighbourPlant.Name != plant.Name {
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
	parseItem := func(char rune) *Plant {
		return &Plant{
			Name:   char,
			Region: nil,
		}
	}

	return World{Matrix: parsers.ParseToMatrix(r, parseItem)}
}
