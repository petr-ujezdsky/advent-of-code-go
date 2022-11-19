package day_09

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type HeightMap = utils.Matrix2[int]

func ValueAt(x, y int, heightMap HeightMap, width, height int) (int, bool) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return 0, false
	}

	return heightMap[x][y], true
}

func InspectNeighbours(heightMap HeightMap, x, y, width, height int, offsetIndexes []utils.Vector2i) (int, bool) {
	value := heightMap[x][y]

	for _, dir := range offsetIndexes {
		neighbour, ok := ValueAt(x+dir.X, y+dir.Y, heightMap, width, height)
		if ok && neighbour <= value {
			// found neighbour of lower value
			return 0, false
		}
	}

	riskLevel := value + 1
	return riskLevel, true
}

func FindLowPointsAndSum(heightMap HeightMap) int {
	width := len(heightMap)
	height := len(heightMap[0])

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

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			riskLevel, ok := InspectNeighbours(heightMap, x, y, width, height, offsetIndexes)
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
