package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/combi"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Asteroid struct {
	Position utils.Vector2i
}

type World struct {
	Matrix    matrix.Matrix[rune]
	Asteroids []Asteroid
}

func DoWithInputPart01(world World) int {
	maxCount := -1
	var maxPosition utils.Vector2i

	for i, asteroidOrigin := range world.Asteroids {
		visibleAsteroids := make(map[utils.Vector2i]utils.Vector2i)

		for j, asteroid := range world.Asteroids {
			if i == j {
				continue
			}

			relativePosition := asteroid.Position.Subtract(asteroidOrigin.Position)
			normalizedLength := combi.GCD(utils.Abs(relativePosition.X), utils.Abs(relativePosition.Y))
			normalized := utils.Vector2i{
				X: relativePosition.X / normalizedLength,
				Y: relativePosition.Y / normalizedLength,
			}

			visibleAsteroids[normalized] = asteroid.Position
		}

		if asteroidOrigin.Position == (utils.Vector2i{X: 5, Y: 8}) {
			fmt.Printf("Best is %v with %v other asteroids detected\n", maxPosition, maxCount)
		}

		visibleCount := len(visibleAsteroids)
		if visibleCount > maxCount {
			maxCount = visibleCount
			maxPosition = asteroidOrigin.Position
		}
	}

	fmt.Printf("Best is %v with %v other asteroids detected\n", maxPosition, maxCount)

	return maxCount
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	var asteroids []Asteroid

	parseItem := func(char rune, x, y int) rune {
		if char == '#' {
			asteroids = append(asteroids, Asteroid{
				Position: utils.Vector2i{X: x, Y: y},
			})
		}
		return char
	}

	return World{
		Matrix:    parsers.ParseToMatrixIndexed(r, parseItem),
		Asteroids: asteroids,
	}
}
