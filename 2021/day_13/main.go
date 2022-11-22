package day_12

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strconv"
	"strings"
)

type Fold struct {
	index      int
	horizontal bool
}

type World struct {
	points []*utils.Vector2i
	folds  []Fold
}

func foldPoint(point *utils.Vector2i, fold Fold) utils.Vector2i {
	if fold.horizontal && point.Y > fold.index {
		// fold up
		y := fold.index - (point.Y - fold.index)
		return utils.Vector2i{point.X, y}
	} else if !fold.horizontal && point.X > fold.index {
		// fold left
		x := fold.index - (point.X - fold.index)
		return utils.Vector2i{x, point.Y}
	}

	// no change
	return *point
}

func foldPoints(points []*utils.Vector2i, fold Fold) {
	for _, point := range points {
		*point = foldPoint(point, fold)
	}
}

func countUniquePoints(points []*utils.Vector2i) int {
	uniquePoints := make(map[int]int)

	for _, point := range points {
		hash := point.X*1_000_000 + point.Y
		uniquePoints[hash]++
	}

	return len(uniquePoints)
}

func printPoints(points []*utils.Vector2i) {
	var maxX, maxY int

	for _, point := range points {
		maxX = utils.Max(maxX, point.X)
		maxY = utils.Max(maxY, point.Y)
	}

	matrix := utils.NewMatrix2[int](maxX+1, maxY+1)

	for _, point := range points {
		matrix.Columns[point.X][point.Y]++
	}

	fmt.Println(matrix.StringFmt(utils.FmtBoolean[int]))
}

func FoldPaper(world World, foldsCount int) int {
	for i, fold := range world.folds {
		if i >= foldsCount {
			break
		}
		foldPoints(world.points, fold)
	}

	//printPoints(world.points)

	return countUniquePoints(world.points)
}

func ParseInput(r io.Reader) (World, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var points []*utils.Vector2i

	var maxX, maxY int

	// points
	for scanner.Scan() && scanner.Text() != "" {
		coordinate := strings.Split(scanner.Text(), ",")

		x, err := strconv.Atoi(coordinate[0])
		if err != nil {
			return World{}, err
		}

		y, err := strconv.Atoi(coordinate[1])
		if err != nil {
			return World{}, err
		}

		pos := utils.Vector2i{x, y}

		maxX = utils.Max(maxX, pos.X)
		maxY = utils.Max(maxY, pos.Y)

		points = append(points, &pos)
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
			index:      index,
			horizontal: foldParts[0][len(foldParts[0])-1] == 'y',
		}

		folds = append(folds, fold)
	}

	world := World{points: points, folds: folds}
	return world, scanner.Err()
}
