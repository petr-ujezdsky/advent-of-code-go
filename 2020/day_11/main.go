package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type SeatsOccupancy = map[utils.Vector2i]bool

func neighboursOccupiedCountInInterval(pos utils.Vector2i, allowedCount utils.IntervalI, occupancy SeatsOccupancy) bool {
	neighboursOccupiedCount := 0

	for _, step := range utils.Direction8Steps {
		neighbour := pos.Add(step)
		if neighbourOccupied, ok := occupancy[neighbour]; ok && neighbourOccupied {
			neighboursOccupiedCount++
		}

		if !allowedCount.Contains(neighboursOccupiedCount) {
			return false
		}
	}

	return true
}

func round(occupancy SeatsOccupancy) (SeatsOccupancy, bool) {
	nextOccupancy := make(SeatsOccupancy)
	anySeatChanged := false

	for pos, occupied := range occupancy {
		if occupied {
			if !neighboursOccupiedCountInInterval(pos, utils.IntervalI{Low: 0, High: 3}, occupancy) {
				// empty the seat
				nextOccupancy[pos] = false
				anySeatChanged = true
				continue
			}
		} else {
			if neighboursOccupiedCountInInterval(pos, utils.IntervalI{Low: 0, High: 0}, occupancy) {
				// fill the seat
				nextOccupancy[pos] = true
				anySeatChanged = true
				continue
			}
		}

		// no change
		nextOccupancy[pos] = occupied
	}

	return nextOccupancy, anySeatChanged
}

func DoWithInputPart01(occupancy SeatsOccupancy) int {
	i := 0
	for {
		nextOccupancy, anySeatChanged := round(occupancy)
		if !anySeatChanged {
			occupiedCount := 0

			for _, occupied := range nextOccupancy {
				if occupied {
					occupiedCount++
				}
			}

			return occupiedCount
		}
		occupancy = nextOccupancy
		i++
	}
}

func DoWithInputPart02(occupancy SeatsOccupancy) int {
	return 0
}

func ParseInput(r io.Reader) SeatsOccupancy {
	seatsOccupancy := make(SeatsOccupancy)

	parseItem := func(char rune, x, y int) int {
		if char == 'L' {
			pos := utils.Vector2i{X: x, Y: y}
			seatsOccupancy[pos] = false
		}
		return 0
	}

	parsers.ParseToMatrixIndexed(r, parseItem)

	return seatsOccupancy
}
