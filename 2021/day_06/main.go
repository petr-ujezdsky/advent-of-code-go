package day_06

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

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
	}

	return timers
}

func ParseInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return utils.ToInts(strings.Split(scanner.Text(), ","))
}
