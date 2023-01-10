package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
)

type Ship struct {
	Direction utils.Direction4
	Position  utils.Vector2i
}

type Instruction struct {
	Amount       int
	ShipModifier ShipModifier
}

type ShipModifier = func(ship Ship, amount int) Ship

var shipModifiers = map[string]ShipModifier{
	"N": func(ship Ship, amount int) Ship {
		step := utils.Up.ToStep().Multiply(amount)
		return Ship{
			Direction: ship.Direction,
			Position:  ship.Position.Add(step),
		}
	},
	"S": func(ship Ship, amount int) Ship {
		step := utils.Down.ToStep().Multiply(amount)
		return Ship{
			Direction: ship.Direction,
			Position:  ship.Position.Add(step),
		}
	},
	"E": func(ship Ship, amount int) Ship {
		step := utils.Right.ToStep().Multiply(amount)
		return Ship{
			Direction: ship.Direction,
			Position:  ship.Position.Add(step),
		}
	},
	"W": func(ship Ship, amount int) Ship {
		step := utils.Left.ToStep().Multiply(amount)
		return Ship{
			Direction: ship.Direction,
			Position:  ship.Position.Add(step),
		}
	},
	"L": func(ship Ship, amount int) Ship {
		steps := amount / 90
		return Ship{
			Direction: ship.Direction.Rotate(-steps),
			Position:  ship.Position,
		}
	},
	"R": func(ship Ship, amount int) Ship {
		steps := amount / 90
		return Ship{
			Direction: ship.Direction.Rotate(steps),
			Position:  ship.Position,
		}
	},
	"F": func(ship Ship, amount int) Ship {
		step := ship.Direction.ToStep().Multiply(amount)
		return Ship{
			Direction: ship.Direction,
			Position:  ship.Position.Add(step),
		}
	},
}

func DoWithInputPart01(instructions []Instruction) int {
	ship := Ship{
		Direction: utils.Right,
		Position:  utils.Vector2i{},
	}

	for _, instruction := range instructions {
		ship = instruction.ShipModifier(ship, instruction.Amount)
		//fmt.Printf("%v: %v\n", i, ship)
	}

	return ship.Position.LengthManhattan()
}

func DoWithInputPart02(instructions []Instruction) int {
	return len(instructions)
}

func ParseInput(r io.Reader) []Instruction {
	parseItem := func(str string) Instruction {
		amount := utils.ExtractInts(str, false)[0]
		shipModifier := shipModifiers[strs.Substring(str, 0, 1)]

		return Instruction{
			Amount:       amount,
			ShipModifier: shipModifier,
		}
	}

	return parsers.ParseToObjects(r, parseItem)
}
