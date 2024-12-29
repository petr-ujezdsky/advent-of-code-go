package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type Plant struct {
	Name   rune
	Region *Region
	Fences [4]bool
}

type World struct {
	Matrix matrix.Matrix[*Plant]
}

type Region struct {
	Id                                 int
	Name                               rune
	Area, Perimeter, PerimeterStraight int
	PerimeterPlants                    map[utils.Vector2i]*Plant
}

func (r Region) String() string {
	return string(r.Name)
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

func floodFill(pos utils.Vector2i, region *Region, used map[utils.Vector2i]struct{}, m matrix.Matrix[*Plant]) *Region {
	if _, ok := used[pos]; ok {
		// already used
		return nil
	}

	used[pos] = struct{}{}

	plant := m.GetV(pos)

	if region == nil {
		region = &Region{
			Id:              0,
			Name:            plant.Name,
			Area:            0,
			Perimeter:       0,
			PerimeterPlants: make(map[utils.Vector2i]*Plant),
		}
	}

	plant.Region = region

	fences := 0
	for i, step := range steps {
		neighbourPos := pos.Add(step)
		neighbourPlant, ok := m.GetVSafe(neighbourPos)
		if !ok {
			fences++
			plant.Fences[i] = true
			region.PerimeterPlants[pos] = plant
			continue
		}

		if neighbourPlant.Name != plant.Name {
			fences++
			plant.Fences[i] = true
			region.PerimeterPlants[pos] = plant
		} else {
			floodFill(neighbourPos, region, used, m)
		}
	}

	region.Area += 1
	region.Perimeter += fences

	return region
}

func DoWithInputPart01(world World) int {
	regions := findRegions(world.Matrix)

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.Area * region.Perimeter
	}

	return totalPrice
}

func findFenceIndex(plant *Plant) (int, bool) {
	for i, fence := range plant.Fences {
		if fence {
			return i, true
		}
	}

	return -1, false
}

func findStraightPerimeters(region *Region) int {
	straightsCount := 0

	for len(region.PerimeterPlants) > 0 {
		pos, plant := maps.FirstEntry(region.PerimeterPlants)

		fenceIndex, ok := findFenceIndex(plant)
		if !ok {
			delete(region.PerimeterPlants, pos)
			continue
		}

		// perpendicular step
		step := steps[(fenceIndex+1)%4]

		// one direction
		length1 := walkPerimeter(pos, fenceIndex, step, region)

		// the other direction
		step = step.Multiply(-1)
		pos = pos.Add(step)
		length2 := walkPerimeter(pos, fenceIndex, step, region)

		if length1+length2 > 0 {
			straightsCount++
		}
	}

	return straightsCount
}

func walkPerimeter(pos utils.Vector2i, fenceIndex int, step utils.Vector2i, region *Region) int {
	plant := region.PerimeterPlants[pos]

	if plant != nil && plant.Fences[fenceIndex] {
		plant.Fences[fenceIndex] = false
		return 1 + walkPerimeter(pos.Add(step), fenceIndex, step, region)
	}

	return 0
}

func DoWithInputPart02(world World) int {
	regions := findRegions(world.Matrix)

	for _, region := range regions {
		straightPerimeters := findStraightPerimeters(region)
		region.PerimeterStraight = straightPerimeters
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.Area * region.PerimeterStraight
	}

	return totalPrice
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
