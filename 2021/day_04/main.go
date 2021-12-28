package day_04

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Bingo struct {
	Numbers [5][5]int
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
	bingo := Bingo{}

	for iRow := 0; scanner.Scan(); iRow++ {
		row := scanner.Text()

		if row == "" {
			bingos = append(bingos, bingo)
			bingo = Bingo{}
			iRow = -1
			continue
		}

		numbers, err := utils.ToInts(strings.Fields(row))
		if err != nil {
			return nil, nil, err
		}

		copy(bingo.Numbers[iRow][:], numbers)
	}

	return drawn, bingos, scanner.Err()
}
