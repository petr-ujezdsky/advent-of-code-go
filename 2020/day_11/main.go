package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	SeatsOccupancy SeatsOccupancy
	Bounds         utils.BoundingRectangle
}
type SeatsOccupancy = map[utils.Vector2i]bool

type OccupiedNeighbourDetector = func(pos, step utils.Vector2i, occupancy SeatsOccupancy, bounds utils.BoundingRectangle) bool

func directNeighbour(pos, step utils.Vector2i, occupancy SeatsOccupancy, _ utils.BoundingRectangle) bool {
	neighbour := pos.Add(step)
	if neighbourOccupied, ok := occupancy[neighbour]; ok && neighbourOccupied {
		return true
	}

	return false
}

func seenNeighbour(pos, step utils.Vector2i, occupancy SeatsOccupancy, bounds utils.BoundingRectangle) bool {
	for {
		pos = pos.Add(step)

		if !bounds.Contains(pos) {
			return false
		}

		neighbourOccupied, ok := occupancy[pos]
		if !ok {
			continue
		}

		return neighbourOccupied
	}
}

func neighboursOccupiedCountInInterval(pos utils.Vector2i, allowedCount utils.IntervalI, occupiedNeighbourDetector OccupiedNeighbourDetector, occupancy SeatsOccupancy, bounds utils.BoundingRectangle) bool {
	neighboursOccupiedCount := 0

	for _, step := range utils.Direction8Steps {
		// is there an occupied neighbour?
		if occupiedNeighbourDetector(pos, step, occupancy, bounds) {
			neighboursOccupiedCount++
		}

		if !allowedCount.Contains(neighboursOccupiedCount) {
			return false
		}
	}

	return true
}

func round(occupiedNeighbourDetector OccupiedNeighbourDetector, countLimit int, occupancy SeatsOccupancy, bounds utils.BoundingRectangle) (SeatsOccupancy, bool) {
	nextOccupancy := make(SeatsOccupancy)
	anySeatChanged := false

	for pos, occupied := range occupancy {
		if occupied {
			if !neighboursOccupiedCountInInterval(pos, utils.IntervalI{Low: 0, High: countLimit - 1}, occupiedNeighbourDetector, occupancy, bounds) {
				// empty the seat
				nextOccupancy[pos] = false
				anySeatChanged = true
				continue
			}
		} else {
			if neighboursOccupiedCountInInterval(pos, utils.IntervalI{Low: 0, High: 0}, occupiedNeighbourDetector, occupancy, bounds) {
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

func countSeats(occupiedNeighbourDetector OccupiedNeighbourDetector, countLimit int, world World) int {
	occupancy := world.SeatsOccupancy
	i := 0
	for {
		nextOccupancy, anySeatChanged := round(occupiedNeighbourDetector, countLimit, occupancy, world.Bounds)
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

func DoWithInputPart01(world World) int {
	return countSeats(directNeighbour, 4, world)
}

func DoWithInputPart02(world World) int {
	return countSeats(seenNeighbour, 5, world)
}

func ParseInput(r io.Reader) World {
	seatsOccupancy := make(SeatsOccupancy)

	parseItem := func(char rune, x, y int) int {
		if char == 'L' {
			pos := utils.Vector2i{X: x, Y: y}
			seatsOccupancy[pos] = false
		}
		return 0
	}

	m := parsers.ParseToMatrixIndexed(r, parseItem)

	return World{
		SeatsOccupancy: seatsOccupancy,
		Bounds:         m.Bounds(),
	}
}
