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

//type RegionRef struct {
//	Region *Region
//}

var steps = slices.Reverse(utils.Direction4Steps[:])

//func DoWithInputPart01_old(world World) int {
//	regionsByPos := make(map[utils.Vector2i]*RegionRef)
//	regionsById := make(map[int]*Region)
//	idSeq := 0
//
//	for x, column := range world.Matrix.Columns {
//		for y, value := range column {
//			pos := utils.Vector2i{X: x, Y: y}
//			//if pos == (utils.Vector2i{X: 2, Y: 7}) {
//			//	fmt.Println()
//			//}
//
//			fences := 0
//			var regionRef *RegionRef
//			for _, step := range steps {
//				neighbourPos := pos.Add(step)
//				neighbourValue, ok := world.Matrix.GetVSafe(neighbourPos)
//				if !ok {
//					fences++
//					continue
//				}
//
//				if neighbourValue != value {
//					fences++
//				} else {
//					if neighbourRegionRef, ok := regionsByPos[neighbourPos]; ok {
//						neighbourRegion := neighbourRegionRef.Region
//						if regionRef != nil && regionRef.Region != neighbourRegion {
//							// merge two regions
//							regionRef.Region.Area += neighbourRegion.Area
//							regionRef.Region.Perimeter += neighbourRegion.Perimeter
//
//							delete(regionsById, neighbourRegionRef.Region.Id)
//							neighbourRegionRef.Region = regionRef.Region
//						}
//
//						regionRef = neighbourRegionRef
//					}
//				}
//			}
//
//			if regionRef == nil {
//				regionRef = &RegionRef{
//					Region: &Region{
//						Id:        idSeq,
//						Name:      value,
//						Area:      0,
//						Perimeter: 0,
//					},
//				}
//
//				regionsById[regionRef.Region.Id] = regionRef.Region
//				idSeq++
//			}
//
//			regionsByPos[pos] = regionRef
//
//			regionRef.Region.Area += 1
//			regionRef.Region.Perimeter += fences
//		}
//	}
//
//	formatter := func(char rune, x, y int) string {
//		regionRef, ok := regionsByPos[utils.Vector2i{X: x, Y: y}]
//		if !ok {
//			panic("No region found")
//		}
//
//		if regionRef.Region.Name != char {
//			panic("Wrong char")
//		}
//
//		return fmt.Sprintf("%2d", regionRef.Region.Id)
//	}
//	str := matrix.StringFmtSeparatorIndexed(world.Matrix, true, " ", formatter)
//	fmt.Println(str)
//
//	totalPrice := 0
//	for _, region := range regionsById {
//		totalPrice += region.Area * region.Perimeter
//	}
//
//	return totalPrice
//}

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
