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

var sandSteps = []Vector2i{
	// down
	{0, 1},
	// down-left
	{-1, 1},
	// down-right
	{1, 1},
}

var sourcePos = Vector2i{500, 0}

type Vector2i = utils.Vector2i

type World struct {
	Cave   utils.Matrix[rune]
	Offset Vector2i
}

type RockDef struct {
	From, To Vector2i
}

func PourSand(world World) int {
	settledCount := 0
	for true {
		unitPos := sourcePos
		settled := false

		for !settled {
			for i, step := range sandSteps {
				unitPosNew := unitPos.Add(step)
				item, ok := world.Cave.GetVSafe(unitPosNew.Subtract(world.Offset))

				// out of world bounds -> end
				if !ok {
					return settledCount
				}

				// air -> accept position
				if item == Air {
					unitPos = unitPosNew
					break
				}

				// last step and item is rock or sand -> can not move -> settled
				if i == len(sandSteps)-1 {
					settled = true
					world.Cave.SetV(unitPos.Subtract(world.Offset), Sand)
				}
			}
		}
		settledCount++
	}

	panic("Should not get here!")
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var items []RockDef
	// start with source dimension
	rangeX, rangeY := utils.IntervalI{sourcePos.X, sourcePos.X}, utils.IntervalI{sourcePos.Y, sourcePos.Y}

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

	cave := utils.NewMatrix[rune](rangeX.Size(), rangeY.Size()).SetAll(Air)
	offset := Vector2i{rangeX.Low, rangeY.Low}

	// fill rocks
	for _, rock := range items {
		step := rock.To.Subtract(rock.From).Signum()
		to := rock.To.Add(step)
		for pos := rock.From; pos != to; pos = pos.Add(step) {
			cave.SetV(pos.Subtract(offset), Rock)
		}
	}

	// fill source
	cave.SetV(sourcePos.Subtract(offset), Source)

	return World{cave, offset}
}
