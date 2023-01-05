package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"io"
)

type World struct {
	CucumbersRight map[Cucumber]struct{}
	CucumbersDown  map[Cucumber]struct{}
	Dimensions     utils.Vector2i
}

type Cucumber struct {
	Position  utils.Vector2i
	Direction utils.Direction4
}

func moduloV(v1 utils.Vector2i, m utils.Vector2i) utils.Vector2i {
	return utils.Vector2i{
		X: v1.X % m.X,
		Y: v1.Y % m.Y,
	}
}

func moveDirection(moving, other map[Cucumber]struct{}, dimensions utils.Vector2i) (next map[Cucumber]struct{}, anyMoved bool) {
	otherDirection := maps.FirstKey(other).Direction
	next = make(map[Cucumber]struct{})
	anyMoved = false

	for cucumber := range moving {
		nextPos := cucumber.Position.Add(cucumber.Direction.ToStep())
		nextPos = moduloV(nextPos, dimensions)

		// check moving set
		if _, ok := moving[Cucumber{Position: nextPos, Direction: cucumber.Direction}]; ok {
			next[cucumber] = struct{}{}
			continue
		}

		// check other set
		if _, ok := other[Cucumber{Position: nextPos, Direction: otherDirection}]; ok {
			next[cucumber] = struct{}{}
			continue
		}

		// position is free -> move
		movedCucumber := Cucumber{
			Position:  nextPos,
			Direction: cucumber.Direction,
		}
		next[movedCucumber] = struct{}{}
		anyMoved = true
	}

	return next, anyMoved
}

func moveRound(world *World) bool {
	anyMoved := false

	// move right
	nextCucumbers, moved := moveDirection(world.CucumbersRight, world.CucumbersDown, world.Dimensions)
	world.CucumbersRight = nextCucumbers
	anyMoved = anyMoved || moved

	// move move down
	nextCucumbers, moved = moveDirection(world.CucumbersDown, world.CucumbersRight, world.Dimensions)
	world.CucumbersDown = nextCucumbers
	anyMoved = anyMoved || moved

	return anyMoved
}

func RoundsUntilStill(world World) int {
	rounds := 1
	for moveRound(&world) {
		rounds++
	}

	return rounds
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	cucumbersRight := make(map[Cucumber]struct{})
	cucumbersDown := make(map[Cucumber]struct{})

	y := 0
	xMax := 0
	for scanner.Scan() {
		xMax = len(scanner.Text())

		for x, char := range scanner.Text() {
			if char == '.' {
				continue
			}

			switch char {
			case '>':
				cucumber := Cucumber{
					Position:  utils.Vector2i{X: x, Y: y},
					Direction: utils.Right,
				}
				cucumbersRight[cucumber] = struct{}{}
			case 'v':
				cucumber := Cucumber{
					Position:  utils.Vector2i{X: x, Y: y},
					Direction: utils.Up,
				}
				cucumbersDown[cucumber] = struct{}{}
			}
		}
		y++
	}

	dimensions := utils.Vector2i{X: xMax, Y: y}

	return World{
		CucumbersRight: cucumbersRight,
		CucumbersDown:  cucumbersDown,
		Dimensions:     dimensions,
	}
}
