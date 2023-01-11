package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type World struct {
	ArriveTime int
	BusIds     []int
}

func DoWithInputPart01(world World) int {
	minWaitTime := math.MaxInt
	minBusId := -1

	for _, busId := range world.BusIds {
		var waitTime int

		if world.ArriveTime%busId == 0 {
			waitTime = 0
		} else {
			waitTime = busId - world.ArriveTime%busId
		}

		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minBusId = busId
		}
	}

	return minWaitTime * minBusId
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	arriveTime := utils.ParseInt(scanner.Text())

	scanner.Scan()
	busIds := utils.ExtractInts(scanner.Text(), false)

	return World{
		ArriveTime: arriveTime,
		BusIds:     busIds,
	}
}
