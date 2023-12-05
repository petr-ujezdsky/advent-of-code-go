package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
	"strings"
)

type Mapper = func(source int) (int, bool)

type Mapping struct {
	Name    string
	Mappers []Mapper
}

func (m Mapping) Map(source int) int {
	for _, mapper := range m.Mappers {
		if destination, ok := mapper(source); ok {
			return destination
		}
	}

	return source
}

type World struct {
	Seeds    []int
	Mappings []Mapping
}

func DoWithInputPart01(world World) int {
	lowest := math.MaxInt

	for _, seed := range world.Seeds {
		location := findLocationForSeed(seed, world.Mappings)

		lowest = utils.Min(lowest, location)
	}

	return lowest
}

func findLocationForSeed(seed int, mappings []Mapping) int {
	sourceDestination := seed

	for _, mapping := range mappings {
		sourceDestination = mapping.Map(sourceDestination)
	}

	return sourceDestination
}

func DoWithInputPart02(world World) int {
	return 0
}

func parseMapper(str string) Mapper {
	numbers := utils.ExtractInts(str, false)

	destination := numbers[0]
	source := numbers[1]
	length := numbers[2]

	return func(src int) (int, bool) {
		if src >= source && src < source+length {
			return destination + (src - source), true
		}

		return 0, false
	}
}

func parseMappings(lines []string, _ int) Mapping {
	name := lines[0]

	mappers := make([]Mapper, len(lines)-1)
	for i, line := range lines[1:] {
		mappers[i] = parseMapper(line)
	}

	return Mapping{
		Name:    name,
		Mappers: mappers,
	}
}

func ParseInput(r io.Reader) World {
	rows := parsers.ParseToStrings(r)
	seeds := utils.ExtractInts(rows[0], false)

	mappersString := strings.Join(rows[2:], "\n")
	r = strings.NewReader(mappersString)

	mappings := parsers.ParseToGroups(r, parseMappings)

	return World{
		Seeds:    seeds,
		Mappings: mappings,
	}
}
