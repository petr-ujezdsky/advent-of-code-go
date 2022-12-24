package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector2i = utils.Vector2i

type Elf = Vector2i

type World map[Elf]struct{}

type Proposition struct {
	DirectionsToCheck [3]Vector2i
	Direction         utils.Direction4
}

var propositionRules = [4]Proposition{
	{
		DirectionsToCheck: [3]Vector2i{utils.North.ToStep(), utils.NorthEast.ToStep(), utils.NorthWest.ToStep()},
		Direction:         utils.Up,
	},
	{
		DirectionsToCheck: [3]Vector2i{utils.South.ToStep(), utils.SouthEast.ToStep(), utils.SouthWest.ToStep()},
		Direction:         utils.Down,
	},
	{
		DirectionsToCheck: [3]Vector2i{utils.West.ToStep(), utils.NorthWest.ToStep(), utils.SouthWest.ToStep()},
		Direction:         utils.Left,
	},
	{
		DirectionsToCheck: [3]Vector2i{utils.East.ToStep(), utils.NorthEast.ToStep(), utils.SouthEast.ToStep()},
		Direction:         utils.Right,
	},
}

func containsAnyElf(elf Elf, directions [3]Vector2i, elves World) bool {
	for _, direction := range directions {
		if _, ok := elves[elf.Add(direction)]; ok {
			return true
		}
	}

	return false
}

func propositionStep(elf Elf, propositionOffset int, elves World) utils.Vector2i {
	emptyCount := 0
	var firstEmpty *Proposition

	for j := 0; j < len(propositionRules); j++ {
		rule := propositionRules[(j+propositionOffset)%len(propositionRules)]

		if containsAnyElf(elf, rule.DirectionsToCheck, elves) {
			if firstEmpty != nil {
				break
			}

			continue
		}

		emptyCount++
		if firstEmpty == nil {
			firstEmpty = &rule
		}
	}

	// all empty or all occupied -> no movement
	if emptyCount == len(propositionRules) || firstEmpty == nil {
		return Vector2i{}
	}

	return firstEmpty.Direction.ToStep()
}

func DoWithInput(elves World) int {
	propositionOffset := 0
	for i := 0; i < 10; i++ {
		proposedPositions := make(map[Vector2i][]Elf)
		// propositions phase
		for elf := range elves {
			// make proposition
			step := propositionStep(elf, propositionOffset, elves)

			proposedPosition := elf.Add(step)
			proposedPositions[proposedPosition] = append(proposedPositions[proposedPosition], elf)
		}

		// move elves
		for elfNew, elvesSamePosition := range proposedPositions {
			if len(elvesSamePosition) > 1 {
				continue
			}

			elfOld := elvesSamePosition[0]
			delete(elves, elfOld)
			elves[elfNew] = struct{}{}
		}
		propositionOffset++
	}

	// bounding box
	boundingBox := utils.BoundingBox{}
	for elf := range elves {
		boundingBox = boundingBox.Enlarge(elf)
	}

	return boundingBox.Horizontal.Size()*boundingBox.Vertical.Size() - len(elves)
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	elves := make(World)

	y := 0
	for scanner.Scan() {
		for x, char := range scanner.Text() {
			if char == '.' {
				continue
			}

			elf := Elf{X: x, Y: y}
			elves[elf] = struct{}{}
		}
		y--
	}

	return elves
}
