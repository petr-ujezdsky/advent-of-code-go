package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type Stone struct {
	Number int
}

type World struct {
	Stones []*Stone
}

func blink(stones []*Stone, times int) int {
	for i := 0; i < times; i++ {
		var extraStones []*Stone
		for _, stone := range stones {
			if stone.Number == 0 {
				stone.Number = 1
				continue
			}

			if digitsCount := utils.DigitsCount(stone.Number); digitsCount%2 == 0 {
				digitsHalf := digitsCount / 2
				k := int(math.Pow10(digitsHalf))

				left := stone.Number / k
				right := stone.Number % k

				stone.Number = left

				extraStones = append(extraStones, &Stone{Number: right})
				continue
			}

			stone.Number *= 2024
		}

		for _, stone := range extraStones {
			stones = append(stones, stone)
		}
	}

	return len(stones)
}

func DoWithInputPart01(world World) int {
	return blink(world.Stones, 25)
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var stones []*Stone
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		for _, value := range ints {
			stones = append(stones, &Stone{Number: value})
		}
	}

	return World{Stones: stones}
}
