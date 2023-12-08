package main

import (
	"bufio"
	_ "embed"
	"io"
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
	Next [2]*MapDef
}

type World struct {
	Directions []Direction2
	Maps       map[string]*MapDef
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
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
	for scanner.Scan() {
		parts := regexMapDef.FindStringSubmatch(scanner.Text())

		name := parts[1]

		left := getOrCreateMapDef(parts[2], maps)
		right := getOrCreateMapDef(parts[3], maps)

		this := getOrCreateMapDef(name, maps)
		this.Next[Left] = left
		this.Next[Right] = right
	}

	return World{
		Directions: directions,
		Maps:       maps,
	}
}
