package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"strings"
)

type Round struct {
	Time, Distance int
}

type World struct {
	Rounds    []Round
	LongRound Round
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
	round := world.LongRound

	// solve quadratic equation
	// -T_t*T_t + T*T_t > D_r
	// where
	// T	total race time
	// T_t  duration for which the button is pushed
	// D_r  record distance
	timePushedLowF := (float64(-round.Time) + math.Sqrt(float64(round.Time*round.Time-4*round.Distance))) / (-2)
	timePushedHighF := (float64(-round.Time) - math.Sqrt(float64(round.Time*round.Time-4*round.Distance))) / (-2)

	// round up
	timePushedLowI := int(math.Round(timePushedLowF + 0.5))

	// round down
	timePushedHighI := int(math.Round(timePushedHighF - 0.5))

	return timePushedHighI - timePushedLowI + 1
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	times := utils.ExtractInts(scanner.Text(), false)
	longTime := utils.ExtractInts(strings.Replace(scanner.Text(), " ", "", -1), false)[0]

	scanner.Scan()
	distances := utils.ExtractInts(scanner.Text(), false)
	longDistance := utils.ExtractInts(strings.Replace(scanner.Text(), " ", "", -1), false)[0]

	rounds := make([]Round, len(times))
	for i, time := range times {
		distance := distances[i]

		rounds[i] = Round{
			Time:     time,
			Distance: distance,
		}
	}

	longRound := Round{
		Time:     longTime,
		Distance: longDistance,
	}

	return World{
		Rounds:    rounds,
		LongRound: longRound,
	}
}
