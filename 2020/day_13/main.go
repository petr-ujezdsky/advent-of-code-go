package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"sort"
	"strings"
)

type Bus struct {
	Id         int
	TimeOffset int
}

type World struct {
	ArriveTime int
	Buses      []Bus
}

func DoWithInputPart01(world World) int {
	minWaitTime := math.MaxInt
	minBusId := -1

	for _, bus := range world.Buses {
		var waitTime int

		if world.ArriveTime%bus.Id == 0 {
			waitTime = 0
		} else {
			waitTime = bus.Id - world.ArriveTime%bus.Id
		}

		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minBusId = bus.Id
		}
	}

	return minWaitTime * minBusId
}

func DoWithInputPart02(world World) int {
	// sort by busId, the largest first
	buses := world.Buses
	sort.Slice(buses, func(i, j int) bool { return buses[i].Id > buses[j].Id })

	t := 0
	i := 0
	step := 1
	for i < len(buses) {
		bus := buses[i]
		for (t+bus.TimeOffset)%bus.Id != 0 {
			t += step
		}

		step *= bus.Id
		i++
	}

	return t
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	arriveTime := utils.ParseInt(scanner.Text())

	scanner.Scan()
	var buses []Bus
	parts := strings.Split(scanner.Text(), ",")
	for i, part := range parts {
		if part == "x" {
			continue
		}

		buses = append(buses, Bus{
			Id:         utils.ParseInt(part),
			TimeOffset: i,
		})
	}

	return World{
		ArriveTime: arriveTime,
		Buses:      buses,
	}
}
