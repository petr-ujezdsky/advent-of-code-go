package day_04

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Bingo struct {
	Numbers                        [5][5]int
	MarkedCountRow, MarkedCountCol [5]int
	SumAll, SumMarked              int
}

func NewBingo(numbers [5][5]int) Bingo {
	sumAll := 0

	for _, row := range numbers {
		for _, number := range row {
			sumAll += number
		}
	}

	return Bingo{
		Numbers:   numbers,
		SumAll:    sumAll,
		SumMarked: 0,
	}
}

func ParseInput(r io.Reader) ([]int, []Bingo, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	drawn, err := utils.ToInts(strings.Split(scanner.Text(), ","))
	if err != nil {
		return nil, nil, err
	}
	scanner.Scan()

	var bingos []Bingo
	numbers := [5][5]int{}

	for iRow := 0; scanner.Scan(); iRow++ {
		row := scanner.Text()

		if row == "" {
			bingo := NewBingo(numbers)
			bingos = append(bingos, bingo)
			numbers = [5][5]int{}
			iRow = -1
			continue
		}

		numbersRow, err := utils.ToInts(strings.Fields(row))
		if err != nil {
			return nil, nil, err
		}

		copy(numbers[iRow][:], numbersRow)
	}

	return drawn, bingos, scanner.Err()
}
