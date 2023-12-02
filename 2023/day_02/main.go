package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"regexp"
	"strings"
)

var regexGame = regexp.MustCompile(`Game (\d+): (.*)`)
var regexRed = regexp.MustCompile(`(\d+) red`)
var regexGreen = regexp.MustCompile(`(\d+) green`)
var regexBlue = regexp.MustCompile(`(\d+) blue`)

type CubeSet struct {
	Red, Green, Blue int
}

func (cs CubeSet) Max(other CubeSet) CubeSet {
	return CubeSet{
		Red:   utils.Max(cs.Red, other.Red),
		Green: utils.Max(cs.Green, other.Green),
		Blue:  utils.Max(cs.Blue, other.Blue),
	}
}

func (cs CubeSet) IsLowerThan(other CubeSet) bool {
	return cs.Red <= other.Red && cs.Green <= other.Green && cs.Blue <= other.Blue
}

type Game struct {
	Id       int
	Examples []CubeSet
}

func (game Game) IsPossibleFor(config CubeSet) bool {
	for _, example := range game.Examples {
		if !example.IsLowerThan(config) {
			return false
		}
	}

	return true
}

type World struct {
	Games []Game
}

func DoWithInputPart01(world World) int {
	bagConfig := CubeSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sum := 0
	for _, game := range world.Games {
		if game.IsPossibleFor(bagConfig) {
			sum += game.Id
		}
	}

	return sum
}

func DoWithInputPart02(world World) int {
	sum := 0
	for _, game := range world.Games {
		max := CubeSet{}
		for _, example := range game.Examples {
			max = max.Max(example)
		}
		sum += max.Red * max.Green * max.Blue
	}

	return sum
}

func toNumber(matches []string) int {
	if len(matches) == 0 {
		return 0
	}

	return utils.ParseInt(matches[1])
}

func parseExample(str string) CubeSet {
	reds := toNumber(regexRed.FindStringSubmatch(str))
	greens := toNumber(regexGreen.FindStringSubmatch(str))
	blues := toNumber(regexBlue.FindStringSubmatch(str))

	return CubeSet{
		Red:   reds,
		Green: greens,
		Blue:  blues,
	}
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Game {
		rowTokens := regexGame.FindStringSubmatch(str)

		id := utils.ParseInt(rowTokens[1])

		examplesRaw := strings.Split(rowTokens[2], ";")
		examples := make([]CubeSet, len(examplesRaw))

		for i, example := range examplesRaw {
			examples[i] = parseExample(example)
		}

		return Game{
			Id:       id,
			Examples: examples,
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Games: items}
}
