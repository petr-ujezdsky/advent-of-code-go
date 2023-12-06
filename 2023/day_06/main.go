package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Round struct {
	Time, Distance int
}

type World struct {
	Rounds []Round
}

func DoWithInputPart01(world World) int {
	product := 1
	for _, round := range world.Rounds {
		recordBeated := 0
		for timePushed := 0; timePushed <= round.Time; timePushed++ {
			timeMoved := round.Time - timePushed
			traveled := timeMoved * timePushed
			fmt.Printf("Time pushed %d, traveled %d\n", timePushed, traveled)

			if traveled > round.Distance {
				recordBeated++
			}
		}
		product *= recordBeated
		fmt.Println()
	}

	return product
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	times := utils.ExtractInts(scanner.Text(), false)

	scanner.Scan()
	distances := utils.ExtractInts(scanner.Text(), false)

	rounds := make([]Round, len(times))
	for i, time := range times {
		distance := distances[i]

		rounds[i] = Round{
			Time:     time,
			Distance: distance,
		}
	}

	return World{Rounds: rounds}
}
