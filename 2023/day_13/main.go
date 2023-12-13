package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strings"
)

type Reading struct {
	Columns, Rows []uint64
}

type World struct {
	Readings []Reading
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, reading := range world.Readings {
		sum += CountReading(reading)
	}

	return sum
}

func CountReading(reading Reading) int {
	if columnsBefore, ok := FindMirror(reading.Columns); ok {
		return columnsBefore + 1
	}

	if rowsBefore, ok := FindMirror(reading.Rows); ok {
		return 100 * (rowsBefore + 1)
	}

	panic("No mirror found")
}

func FindMirror(items []uint64) (int, bool) {
	for mirror := 0; mirror < len(items)-1; mirror++ {
		if checkMirror(mirror, items) {
			return mirror, true
		}
	}

	return 0, false
}

func checkMirror(mirror int, items []uint64) bool {
	maxStep := utils.Min(mirror, len(items)-mirror-2)

	for step := 0; step <= maxStep; step++ {
		backward := items[mirror-step]
		forward := items[mirror+1+step]

		if backward != forward {
			return false
		}
	}

	return true
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseGroup := func(lines []string, i int) Reading {
		reader := strings.NewReader(strings.Join(lines, "\n"))

		boolMatrix := parsers.ParseToMatrix(reader, parsers.MapperBoolean('#', '.'))

		columns := slices.Map(boolMatrix.Columns, utils.ParseBinaryBool64)

		boolMatrix = boolMatrix.Transpose()
		rows := slices.Map(boolMatrix.Columns, utils.ParseBinaryBool64)

		return Reading{
			Columns: columns,
			Rows:    rows,
		}
	}

	readings := parsers.ParseToGroups(r, parseGroup)
	return World{Readings: readings}
}
