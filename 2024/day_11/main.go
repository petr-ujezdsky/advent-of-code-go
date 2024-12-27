package main

import (
	"bufio"
	_ "embed"
	"fmt"
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
	stonesMap := make(map[int]int)
	for _, stone := range stones {
		stonesMap[stone.Number]++
	}

	for i := 0; i < times; i++ {
		fmt.Printf("Step #%v\n", i)
		nextStonesMap := make(map[int]int)
		for stone, count := range stonesMap {
			if stone == 0 {
				// 0 -> 1
				nextStonesMap[1] += count
				continue
			}

			if digitsCount := utils.DigitsCount(stone); digitsCount%2 == 0 {
				digitsHalf := digitsCount / 2
				k := int(math.Pow10(digitsHalf))

				left := stone / k
				right := stone % k

				nextStonesMap[left] += count
				nextStonesMap[right] += count

				continue
			}

			// N * 2024
			nextStonesMap[stone*2024] += count
		}

		stonesMap = nextStonesMap
	}

	total := 0
	for _, count := range stonesMap {
		total += count
	}

	return total
}

func DoWithInputPart01(world World) int {
	return blink(world.Stones, 25)
}

func DoWithInputPart02(world World) int {
	return blink(world.Stones, 75)
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
