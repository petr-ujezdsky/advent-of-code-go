package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
)

type Direction2 int

const (
	Left Direction2 = iota
	Right
)

var regexMapDef = regexp.MustCompile(`(...) = \((...), (...)\)`)

type Shortcut struct {
	EndMap *MapDef
	Steps  int
}

type MapDef struct {
	Name          string
	End           bool
	Next          [2]*MapDef
	ShortcutToEnd map[int]Shortcut
}

type World struct {
	Directions   []Direction2
	Maps         map[string]*MapDef
	StartingMaps []*MapDef
}

func DoWithInputPart01(world World) int {
	stepper := &Stepper{
		Directions:      world.Directions,
		Position:        0,
		Current:         world.Maps["AAA"],
		LastEndMap:      nil,
		LastEndPosition: -1,
	}

	for stepper.Current.Name != "ZZZ" {
		stepper.Move(1)
	}

	return stepper.Position
}

type Stepper struct {
	Directions []Direction2
	Position   int

	Current         *MapDef
	LastEndMap      *MapDef
	LastEndPosition int
}

func (s *Stepper) Move(steps int) {
	for i := 0; i < steps; i++ {
		dirIndex := utils.ModFloor(s.Position, len(s.Directions))
		dir := s.Directions[dirIndex]

		next := s.Current.Next[dir]
		s.Current = next
		s.Position++

		// fill shortcut
		if next.End {
			if s.LastEndMap != nil {
				shortcut := Shortcut{
					EndMap: next,
					Steps:  s.Position - s.LastEndPosition,
				}

				lastEndIndex := utils.ModFloor(s.LastEndPosition, len(s.Directions))
				s.LastEndMap.ShortcutToEnd[lastEndIndex] = shortcut
			}

			s.LastEndMap = next
			s.LastEndPosition = s.Position
		}
	}
}

func FindLoop(mapDef *MapDef, directions []Direction2) (int, int) {
	position := 0

	visited := make(map[*MapDef][]int)

	current := mapDef
	lastEndPosition := -1
	for {
		if current.End {
			lastEndPosition = position
			//fmt.Printf("%s end @ %d\n", mapDef.Name, position)
		}

		dirIndex := utils.ModFloor(position, len(directions))
		dir := directions[dirIndex]

		visitedPositions, ok := visited[current]
		if !ok {
			visitedPositions = slices.Filled(-1, len(directions))
			visited[current] = visitedPositions
		}

		if previousPosition := visitedPositions[dirIndex]; previousPosition != -1 {
			return lastEndPosition, position - previousPosition
		}

		visitedPositions[dirIndex] = position

		next := current.Next[dir]
		current = next
		position++
	}
}

func (s *Stepper) MaxStepSize() int {
	dirIndex := utils.ModFloor(s.Position, len(s.Directions))

	if shortcut, ok := s.Current.ShortcutToEnd[dirIndex]; ok {
		return shortcut.Steps
	}

	return 1
}

func DoWithInputPart02(world World) int {
	maxFrom := -1
	for _, startingMap := range world.StartingMaps {
		from, length := FindLoop(startingMap, world.Directions)
		fmt.Printf("%s: loop from %d, length %d\n", startingMap.Name, from, length)

		maxFrom = utils.Max(maxFrom, from)
	}

	return 0
}

func countStepsToEnd(current *MapDef, directions []Direction2) int {
	i := 0

	for !current.End {
		dir := directions[utils.ModFloor(i, len(directions))]
		next := current.Next[dir]
		current = next
		i++
	}

	return i
}

func getOrCreateMapDef(name string, maps map[string]*MapDef) *MapDef {
	mapDef, ok := maps[name]

	if !ok {
		mapDef = &MapDef{
			Name:          name,
			Next:          [2]*MapDef{nil, nil},
			ShortcutToEnd: make(map[int]Shortcut),
		}

		maps[name] = mapDef
	}

	return mapDef
}

func parseDirections(str string) []Direction2 {
	directions := make([]Direction2, len(str))

	for i, char := range str {
		if char == 'R' {
			directions[i] = Right
		} else {
			directions[i] = Left
		}
	}

	return directions
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	directions := parseDirections(scanner.Text())
	scanner.Scan()

	maps := make(map[string]*MapDef)
	var startingMaps []*MapDef
	for scanner.Scan() {
		parts := regexMapDef.FindStringSubmatch(scanner.Text())

		name := parts[1]

		left := getOrCreateMapDef(parts[2], maps)
		right := getOrCreateMapDef(parts[3], maps)

		this := getOrCreateMapDef(name, maps)
		this.Next[Left] = left
		this.Next[Right] = right
		this.End = name[2] == 'Z'

		if name[2] == 'A' {
			startingMaps = append(startingMaps, this)
		}
	}

	return World{
		Directions:   directions,
		Maps:         maps,
		StartingMaps: startingMaps,
	}
}
