package main

import (
	"bufio"
	_ "embed"
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
		Directions: world.Directions,
		Position:   0,
		Current:    world.Maps["AAA"],
		LastEnd:    nil,
	}

	for stepper.Current.Name != "ZZZ" {
		stepper.Move(1)
	}

	return stepper.Position
}

type Stepper struct {
	Directions []Direction2
	Position   int

	Current *MapDef
	LastEnd *MapDef
}

func (s *Stepper) Move(steps int) {
	for i := 0; i < steps; i++ {
		dirIndex := utils.ModFloor(s.Position, len(s.Directions))
		dir := s.Directions[dirIndex]

		next := s.Current.Next[dir]
		s.Current = next
		s.Position++
	}
}

func DoWithInputPart02(world World) int {
	i := 0

	//histories := make([][]*MapDef, len(world.StartingMaps))
	lastEnds := make([]*Shortcut, len(world.StartingMaps))
	currents := slices.Clone(world.StartingMaps)
	end := false

	for !end {
		dirIndex := utils.ModFloor(i, len(world.Directions))
		dir := world.Directions[dirIndex]
		end = true

		for j := 0; j < len(currents); j++ {
			// store current into history
			//histories[j] = append(histories[j], currents[j])

			next := currents[j].Next[dir]
			currents[j] = next

			if next.End {
				// fill shortcuts
				//for ih, mapHistory := range histories[j] {
				//	//steps := i + 1 - ih
				//	steps := len(histories[j]) - ih
				//	Shortcut{
				//		EndMap: next,
				//		Steps:  0,
				//	}
				//	mapHistory.ShortcutToEnd[]
				//}

				// fill shortcut
				lastEnd := lastEnds[j]
				if lastEnd == nil {
					lastEnds[j] = &Shortcut{
						EndMap: next,
						Steps:  i + 1,
					}
				}
			} else {
				// not ending map -> must continue
				end = false
			}
		}

		i++
	}

	return i
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
			Name: name,
			Next: [2]*MapDef{nil, nil},
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
