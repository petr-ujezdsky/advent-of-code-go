package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Report struct {
	Levels []int
}

type World struct {
	Reports []Report
}

func isReportSafe(levels []int) bool {
	previousLevel := levels[0]
	previousSignum := 0

	for _, level := range levels[1:] {
		diff := level - previousLevel

		if diff == 0 || utils.Abs(diff) > 3 {
			return false
		}

		signum := utils.Signum(diff)
		if previousSignum != 0 && previousSignum != signum {
			return false
		}

		previousLevel = level
		previousSignum = signum
	}

	return true
}

func DoWithInputPart01(world World) int {
	safeCount := 0

	for _, report := range world.Reports {
		if isReportSafe(report.Levels) {
			safeCount++
		}
	}

	return safeCount
}

func isReportSafeDampened(report Report) bool {
	for i := -1; i < len(report.Levels); i++ {
		var levels []int

		if i == -1 {
			levels = report.Levels
		} else {
			for index, level := range report.Levels {
				if index == i {
					continue
				}
				levels = append(levels, level)
			}
		}

		if isReportSafe(levels) {
			return true
		}
	}

	return false
}

func DoWithInputPart02(world World) int {
	safeCount := 0

	for _, report := range world.Reports {
		if isReportSafeDampened(report) {
			safeCount++
		}
	}

	return safeCount
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Report {
		ints := utils.ExtractInts(str, true)

		return Report{Levels: ints}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Reports: items}
}
