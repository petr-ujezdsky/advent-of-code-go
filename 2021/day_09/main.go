package day_09

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type HeightMap = utils.Matrix2[int]

func InspectNeighbours(heightMap HeightMap, x, y int, steps []utils.Vector2i) (int, bool) {
	value := heightMap.Get(x, y)

	for _, step := range steps {
		neighbour, ok := heightMap.GetSafe(x+step.X, y+step.Y)
		if ok && neighbour <= value {
			// found neighbour of lower value
			return 0, false
		}
	}

	riskLevel := value + 1
	return riskLevel, true
}

func FindLowPointsAndSum(heightMap HeightMap) int {
	offsetIndexes := []utils.Vector2i{
		// left
		{-1, 0},
		// right
		{1, 0},
		// up
		{0, -1},
		// down
		{0, 1},
	}

	lowPointsRiskLevelsSum := 0

	for x := 0; x < heightMap.Width; x++ {
		for y := 0; y < heightMap.Height; y++ {
			riskLevel, ok := InspectNeighbours(heightMap, x, y, offsetIndexes)
			if ok {
				lowPointsRiskLevelsSum += riskLevel
			}
		}
	}

	return lowPointsRiskLevelsSum
}

func ParseInput(r io.Reader) (HeightMap, error) {
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
