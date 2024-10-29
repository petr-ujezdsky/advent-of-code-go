package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/combi"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
	"slices"
)

type Asteroid struct {
	Position utils.Vector2i
}

type World struct {
	Matrix    matrix.Matrix[rune]
	Asteroids []Asteroid
}

type RelativeAsteroid struct {
	OriginalPosition, Direction utils.Vector2i
	Length                      int
}

func findMaxVisibleAsteroids(asteroids []Asteroid) (map[utils.Vector2i][]RelativeAsteroid, utils.Vector2i) {
	maxCount := -1
	var maxPosition utils.Vector2i
	var maxVisibleAsteroids map[utils.Vector2i][]RelativeAsteroid

	for i, asteroidOrigin := range asteroids {
		visibleAsteroids := make(map[utils.Vector2i][]RelativeAsteroid)

		for j, asteroid := range asteroids {
			if i == j {
				continue
			}

			relativePosition := asteroid.Position.Subtract(asteroidOrigin.Position)
			normalizedLength := combi.GCD(utils.Abs(relativePosition.X), utils.Abs(relativePosition.Y))
			normalized := utils.Vector2i{
				X: relativePosition.X / normalizedLength,
				Y: relativePosition.Y / normalizedLength,
			}

			visibleAsteroids[normalized] = append(visibleAsteroids[normalized], RelativeAsteroid{
				OriginalPosition: asteroid.Position,
				Direction:        normalized,
				Length:           normalizedLength,
			})
		}

		visibleCount := len(visibleAsteroids)
		if visibleCount > maxCount {
			maxCount = visibleCount
			maxPosition = asteroidOrigin.Position
			maxVisibleAsteroids = visibleAsteroids
		}
	}

	return maxVisibleAsteroids, maxPosition
}

func DoWithInputPart01(world World) int {
	maxVisibleAsteroids, maxPosition := findMaxVisibleAsteroids(world.Asteroids)

	fmt.Printf("Best is %v with %v other asteroids detected\n", maxPosition, len(maxVisibleAsteroids))

	return len(maxVisibleAsteroids)
}

func DoWithInputPart02(world World) int {
	maxVisibleAsteroids, _ := findMaxVisibleAsteroids(world.Asteroids)

	for _, asteroids := range maxVisibleAsteroids {
		// sort asteroids in given direction by distance
		slices.SortFunc(asteroids, func(a, b RelativeAsteroid) int {
			return cmp.Compare(a.Length, b.Length)
		})
	}

	directions := maps.Keys(maxVisibleAsteroids)
	// sort by angle, zero is at north, rising clockwise
	slices.SortFunc(directions, func(a, b utils.Vector2i) int {
		angleA := angle(a)
		angleB := angle(b)
		return cmp.Compare(angleA, angleB)
	})

	shots := matrix.NewMatrixInt(world.Matrix.Width, world.Matrix.Height)

	shotsCount := 0
	for i := 0; len(maxVisibleAsteroids) > 0; i++ {
		direction := directions[i%len(directions)]
		if _, ok := maxVisibleAsteroids[direction]; !ok {
			continue
		}

		asteroid := maxVisibleAsteroids[direction][0]
		//fmt.Printf("Vaporizing %v\n", asteroid.OriginalPosition)
		shots.SetV(asteroid.OriginalPosition, shotsCount+1)

		// vaporize it
		maxVisibleAsteroids[direction] = maxVisibleAsteroids[direction][1:]
		if len(maxVisibleAsteroids[direction]) == 0 {
			// no more asteroids in given direction, delete it
			delete(maxVisibleAsteroids, direction)
		}

		shotsCount++
		if shotsCount == 200 {
			return 100*asteroid.OriginalPosition.X + asteroid.OriginalPosition.Y
		}
	}

	origin := utils.Vector2i{}

	fmt.Println(matrix.StringFmtSeparatorIndexedOrigin[int](shots, 2, origin, " ", matrix.NonIndexedAdapter(matrix.FmtFmt[int]("%2d"))))

	panic("No solution found")
}

func angle(v utils.Vector2i) float64 {
	// -Y because matrix is upside down
	realAngle := math.Atan2(float64(-v.Y), float64(v.X))

	transformed := -(realAngle - math.Pi/2)
	if transformed >= 0 {
		return transformed
	}

	return 2*math.Pi + transformed
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
