package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"io"
)

type World struct {
	Blueprints []Blueprint
}

const (
	Ore = iota
	Clay
	Obsidian
	Geode
)

type RobotCosts [4]int

type Blueprint struct {
	Id          int
	RobotsCosts [4]RobotCosts
}

type Robot struct {
	Type      int
	SinceTime int
}

type State struct {
	RemainingTime int
	Materials     [4]int
	Robots        []Robot
}

func DoWithInput(world World) int {
	cost := func(state State) int {}

	lowerBound := func(state State) int {}

	nextStatesProvider := func(state State) []State {}

	initialState := State{
		RemainingTime: 24,
		Materials:     [4]int{},
		Robots:        nil,
	}

	min, _ := alg.BranchAndBoundDeepFirst(initialState, cost, lowerBound, nextStatesProvider)
	return min
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var blueprints []Blueprint
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		id := ints[0]
		robotsCosts := [4]RobotCosts{
			{ints[1], 0, 0, 0},
			{ints[2], 0, 0, 0},
			{ints[3], ints[4], 0, 0},
			{ints[5], 0, ints[6], 0},
		}

		item := Blueprint{
			Id:          id,
			RobotsCosts: robotsCosts,
		}

		blueprints = append(blueprints, item)
	}

	return World{Blueprints: blueprints}
}
