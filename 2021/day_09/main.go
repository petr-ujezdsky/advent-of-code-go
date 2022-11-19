package day_09

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"sort"
)

type Matrix2i = utils.Matrix2[int]

var STEPS = []utils.Vector2i{
	// left
	{-1, 0},
	// right
	{1, 0},
	// up
	{0, -1},
	// down
	{0, 1},
}

func inspectNeighbours(heightMap Matrix2i, x, y int) (int, bool) {
	value := heightMap.Get(x, y)

	for _, step := range STEPS {
		neighbour, ok := heightMap.GetSafe(x+step.X, y+step.Y)
		if ok && neighbour <= value {
			// found neighbour of lower value
			return 0, false
		}
	}

	riskLevel := value + 1
	return riskLevel, true
}

func FindLowPointsAndSum(heightMap Matrix2i) (int, []utils.Vector2i) {
	lowPointsRiskLevelsSum := 0
	var lowPoints []utils.Vector2i

	for x := 0; x < heightMap.Width; x++ {
		for y := 0; y < heightMap.Height; y++ {
			riskLevel, ok := inspectNeighbours(heightMap, x, y)
			if ok {
				lowPointsRiskLevelsSum += riskLevel

				// save the low point position
				lowPoints = append(lowPoints, utils.Vector2i{X: x, Y: y})
			}
		}
	}

	return lowPointsRiskLevelsSum, lowPoints
}

func findBasinSizeRecursive(heightMap, basin Matrix2i, position utils.Vector2i) int {
	if basin.Get(position.X, position.Y) != 0 {
		// already inspected -> end
		return 0
	}

	// save basin point location
	basin.Set(position.X, position.Y, 1)

	value := heightMap.Get(position.X, position.Y)

	if value == 9 {
		// 9 is not part of the basin
		return 0
	}

	size := 1
	// inspect neighbours
	for _, step := range STEPS {
		neighbourPosition := position.Add(step)

		neighbourValue, ok := heightMap.GetSafe(neighbourPosition.X, neighbourPosition.Y)

		// neighbour is part of the basin
		if ok && neighbourValue > value {
			// inspect it recursively
			size += findBasinSizeRecursive(heightMap, basin, neighbourPosition)
		}
	}

	return size
}

func findBasinSize(heightMap Matrix2i, position utils.Vector2i) int {
	// create empty matrix to write found basin points
	basin := utils.NewMatrix2[int](heightMap.Width, heightMap.Height)

	// find basin size
	return findBasinSizeRecursive(heightMap, basin, position)
}

func Basins(heightMap Matrix2i) int {
	_, lowPoints := FindLowPointsAndSum(heightMap)

	var basinSizes []int

	for _, lowPoint := range lowPoints {
		basinSize := findBasinSize(heightMap, lowPoint)
		basinSizes = append(basinSizes, basinSize)
	}

	// sort sizes
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	// multiply 3 largest basin sizes
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func ParseInput(r io.Reader) (Matrix2i, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rows [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int

		for _, digitAscii := range []rune(line) {
			digit := int(digitAscii) - int('0')
			row = append(row, digit)
		}

		rows = append(rows, row)
	}

	return utils.NewMatrix2RowNotation(rows), scanner.Err()
}
