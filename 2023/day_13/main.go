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
	return 0
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
