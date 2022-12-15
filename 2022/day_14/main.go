package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

const (
	Air    rune = '.'
	Rock        = '#'
	Sand        = 'o'
	Source      = '+'
)

type Vector2i = utils.Vector2i

type World = utils.Matrix[rune]

type RockDef struct {
	From, To Vector2i
}

func DoWithInput(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var items []RockDef
	// start with source dimension
	rangeX, rangeY := utils.IntervalI{500, 500}, utils.IntervalI{0, 0}

	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		for i := 0; i < len(ints)-3; i += 2 {
			rock := RockDef{
				From: Vector2i{ints[i+0], ints[i+1]},
				To:   Vector2i{ints[i+2], ints[i+3]},
			}

			rangeX = rangeX.Enlarge(rock.From.X).Enlarge(rock.To.X)
			rangeY = rangeY.Enlarge(rock.From.Y).Enlarge(rock.To.Y)

			items = append(items, rock)
		}
	}

	world := utils.NewMatrix[rune](rangeX.Size(), rangeY.Size()).SetAll(Air)

	// fill rocks
	for _, rock := range items {
		step := rock.To.Subtract(rock.From).Signum()
		to := rock.To.Add(step)
		for pos := rock.From; pos != to; pos = pos.Add(step) {
			world.Columns[pos.X-rangeX.Low][pos.Y-rangeY.Low] = Rock
		}
	}

	// fill source
	world.Columns[500-rangeX.Low][0-rangeY.Low] = Source

	return world
}
