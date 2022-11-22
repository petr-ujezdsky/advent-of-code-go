package day_12

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strconv"
	"strings"
)

type Fold struct {
	index int
	dirX  bool
}

type World struct {
	points []utils.Vector2i
	folds  []Fold
}

func ParseInput(r io.Reader) (World, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var points []utils.Vector2i

	var maxX, maxY int

	// points
	for scanner.Scan() && scanner.Text() != "" {
		coordinate := strings.Split(scanner.Text(), ",")

		pos := utils.Vector2i{
			X: int(coordinate[0][0] - '0'),
			Y: int(coordinate[1][0] - '0'),
		}

		maxX = utils.Max(maxX, pos.X)
		maxY = utils.Max(maxY, pos.Y)

		points = append(points, pos)
	}

	// folds
	var folds []Fold

	for scanner.Scan() {
		foldParts := strings.Split(scanner.Text(), "=")

		index, err := strconv.Atoi(foldParts[1])
		if err != nil {
			return World{}, err
		}

		fold := Fold{
			index: index,
			dirX:  foldParts[0][len(foldParts[0])-1] == 'x',
		}

		folds = append(folds, fold)
	}

	world := World{points: points, folds: folds}
	return world, scanner.Err()
}
