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
	Winning                        bool
}

func NewBingo(numbers [5][5]int) *Bingo {
	sumAll := 0

	for _, row := range numbers {
		for _, number := range row {
			sumAll += number
		}
	}

	return &Bingo{
		Numbers:   numbers,
		SumAll:    sumAll,
		SumMarked: 0,
	}
}

func (bingo *Bingo) Mark(numberDrawn int) (bool, int) {
	for iRow, row := range bingo.Numbers {
		for iCol, number := range row {
			if number == numberDrawn {
				bingo.MarkedCountCol[iCol]++
				bingo.MarkedCountRow[iRow]++
				bingo.SumMarked += number

				if bingo.MarkedCountCol[iCol] == 5 || bingo.MarkedCountRow[iRow] == 5 {
					score := (bingo.SumAll - bingo.SumMarked) * numberDrawn
					bingo.Winning = true

					return true, score
				}

				return false, -1
			}
		}
	}

	return false, -1
}

func Play(bingos []*Bingo, drawn []int) (*Bingo, int) {
	for _, number := range drawn {
		for _, bingo := range bingos {
			winning, score := bingo.Mark(number)
			if winning {
				return bingo, score
			}
		}
	}

	return nil, -1
}

func PlayTillEnd(bingos []*Bingo, drawn []int) (*Bingo, int) {
	var winBingo *Bingo
	winScore := -1

	for _, number := range drawn {
		for _, bingo := range bingos {
			if bingo.Winning {
				continue
			}

			winning, score := bingo.Mark(number)
			if winning {
				winBingo = bingo
				winScore = score
			}
		}
	}

	return winBingo, winScore
}

func ParseInput(r io.Reader) ([]int, []*Bingo, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	drawn, err := utils.ToInts(strings.Split(scanner.Text(), ","))
	if err != nil {
		return nil, nil, err
	}
	scanner.Scan()

	var bingos []*Bingo
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

	bingo := NewBingo(numbers)
	bingos = append(bingos, bingo)

	return drawn, bingos, scanner.Err()
}
