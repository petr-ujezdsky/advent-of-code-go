package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"math"
	"regexp"
)

type Direction2 int

const (
	Left Direction2 = iota
	Right
)

var regexMapDef = regexp.MustCompile(`(...) = \((...), (...)\)`)

type MapDef struct {
	Name string
	End  bool
	Next [2]*MapDef
}

type World struct {
	Directions   []Direction2
	Maps         map[string]*MapDef
	StartingMaps []*MapDef
}

func DoWithInputPart01(world World) int {
	i := 0
	current := world.Maps["AAA"]

	for current.Name != "ZZZ" {
		dir := world.Directions[utils.ModFloor(i, len(world.Directions))]
		next := current.Next[dir]
		current = next
		i++
	}

	return i
}

func DoWithInputPart02(world World) int {
	minLength := math.MaxInt
	minLengthFrom := -1

	fromLengths := make([][2]int, len(world.StartingMaps))

	// find smallest period
	for i, startingMap := range world.StartingMaps {
		from, length := findLoop(startingMap, world.Directions)
		fmt.Printf("%s: loop from %d, length %d\n", startingMap.Name, from, length)

		fromLengths[i] = [2]int{from, length}

		if length < minLength {
			minLength = length
			minLengthFrom = from
		}
	}

	// find common ends
	position := minLengthFrom
	for {
		end := true

		for _, fromLength := range fromLengths {
			from := fromLength[0]
			length := fromLength[1]

			if (position-from)%length != 0 {
				end = false
				break
			}
		}

		if end {
			break
		}

		position += minLength
	}

	return position
}

func findLoop(mapDef *MapDef, directions []Direction2) (int, int) {
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
