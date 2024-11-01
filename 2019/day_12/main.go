package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Moon struct {
	Id                 int
	Position, Velocity utils.Vector3i
}

type World struct {
	Moons []*Moon
}

func applyGravity(moons []*Moon) {
	for i, moon1 := range moons {
		for _, moon2 := range moons[i+1:] {
			//fmt.Printf("Moons %v, %v\n", moon1.Id, moon2.Id)
			//fmt.Printf("Moons %v, %v\n", moon1.Position, moon2.Position)
			//fmt.Printf("Moons %v, %v\n", moon1.Velocity, moon2.Velocity)

			diffSignum := moon1.Position.Subtract(moon2.Position).Signum()

			moon1.Velocity = moon1.Velocity.Subtract(diffSignum)
			moon2.Velocity = moon2.Velocity.Add(diffSignum)
			//fmt.Printf("Moons %v, %v\n\n", moon1.Velocity, moon2.Velocity)
		}
	}
}

func applyVelocity(moons []*Moon) {
	for _, moon := range moons {
		moon.Position = moon.Position.Add(moon.Velocity)
	}
}

func totalEnergy(moons []*Moon) int {
	total := 0

	for _, moon := range moons {
		potential := moon.Position.ManhattanLength()
		kinetic := moon.Velocity.ManhattanLength()

		total += potential * kinetic
	}

	return total
}

func printMoons(i int, moons []*Moon) {
	if i > 0 {
		fmt.Println()
	}
	fmt.Printf("After %v steps:\n", i)
	for _, moon := range moons {
		fmt.Printf("pos=%v, vel=%v\n", moon.Position, moon.Velocity)
	}

	fmt.Printf("Total energy: %v\n", totalEnergy(moons))
}

func DoWithInputPart01(world World, steps int) int {
	for i := 0; i < steps; i++ {
		//printMoons(i, world.Moons)

		applyGravity(world.Moons)
		applyVelocity(world.Moons)
	}

	return totalEnergy(world.Moons)
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	sequence := -1

	parseItem := func(str string) *Moon {
		sequence++
		numbers := utils.ExtractInts(str, true)

		return &Moon{
			Id:       sequence,
			Position: utils.Vector3i{X: numbers[0], Y: numbers[1], Z: numbers[2]},
			Velocity: utils.Vector3i{},
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Moons: items}
}
