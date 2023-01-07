package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"sort"
)

type BoardingPass struct {
	Row, Col int
}

func NewBoardingPass(str string) BoardingPass {
	chars := []rune(str)
	return BoardingPass{
		Row: calcRange(
			chars[0:7],
			'F',
			0,
			0,
			127,
		),
		Col: calcRange(
			chars[7:],
			'L',
			0,
			0,
			7,
		),
	}
}

func (bp BoardingPass) getSeatId() int {
	return bp.Row*8 + bp.Col
}

func calcRange(data []rune, lower rune, index, min, max int) int {
	middle := (max-min+1)/2 + min - 1

	if min == max {
		return min
	}

	if data[index] == lower {
		// lower
		return calcRange(data, lower, index+1, min, middle)
	}

	// upper
	return calcRange(data, lower, index+1, middle+1, max)
}

func FindMaxSeatId(boardingPasses []BoardingPass) int {
	maxId := -1

	for _, boardingPass := range boardingPasses {
		maxId = utils.Max(maxId, boardingPass.getSeatId())
	}

	return maxId
}

func FindMissingSeatId(boardingPasses []BoardingPass) int {
	seatIds := slices.Map(boardingPasses, func(bp BoardingPass) int { return bp.getSeatId() })

	sort.Ints(seatIds)

	for i := 0; i < len(seatIds)-1; i++ {
		s1 := seatIds[i]
		s2 := seatIds[i+1]

		if s1+1 != s2 {
			return s1 + 1
		}
	}

	panic("Nothing found")
}

func ParseInput(r io.Reader) []BoardingPass {
	return parsers.ParseToObjects(r, NewBoardingPass)
}
