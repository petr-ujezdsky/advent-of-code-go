package day_06

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Fish struct {
	DaysToBirth, Remaining int
}

func PassDay(timers []int) []int {
	var spawned []int

	for i := range timers {
		if timers[i] == 0 {
			// reset and spawn new fish
			timers[i] = 6

			spawned = append(spawned, 8)
		} else {
			timers[i]--
		}
	}

	return append(timers, spawned...)
}

func PassDays(timers []int, days int) []int {
	for i := 0; i < days; i++ {
		timers = PassDay(timers)
		//fmt.Println("Day", i, ", count", len(timers))
	}

	return timers
}

func CountFish(fish Fish) int {
	// 1 = fish itself
	count := 1

	for fish.Remaining >= fish.DaysToBirth {
		fish.Remaining -= fish.DaysToBirth
		fish.DaysToBirth = 7

		child := Fish{9, fish.Remaining}
		count += CountFish(child)
	}

	return count
}

func CountFishMultithread(fish Fish, count chan int) {
	count <- CountFish(fish)
}

func CountManyFish(fish []Fish) int {
	count := make(chan int)

	for _, f := range fish {
		go CountFishMultithread(f, count)
	}

	total := 0
	for range fish {
		total += <-count
	}

	return total
}

func ParseInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return utils.ToInts(strings.Split(scanner.Text(), ","))
}

func ParseFish(r io.Reader, days int) ([]Fish, error) {
	timers, err := ParseInput(r)
	if err != nil {
		return nil, err
	}

	return CreateFish(timers, days), nil
}

func CreateFish(timers []int, days int) []Fish {
	var fish []Fish

	for _, timer := range timers {
		fish = append(fish, Fish{timer + 1, days})
	}

	return fish
}
