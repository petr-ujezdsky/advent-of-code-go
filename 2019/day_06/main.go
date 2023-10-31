package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"io"
	"strings"
)

type Planet struct {
	Name     string
	Parent   *Planet
	Children []*Planet
}

func NewPlanet(name string) *Planet {
	return &Planet{Name: name}
}

type World struct {
	Root    *Planet
	Planets map[string]*Planet
}

func DoWithInputPart01(world World) int {
	directSum, indirectSum := 0, 0
	for _, planet := range world.Planets {
		direct, indirect := countDirectIndirectOrbits(planet)

		directSum += direct
		indirectSum += indirect
	}

	return directSum + indirectSum
}

func countDirectIndirectOrbits(planet *Planet) (int, int) {
	depth := 0

	parent := planet.Parent

	for parent != nil {
		depth++
		parent = parent.Parent
	}

	if depth != 0 {
		return 1, depth - 1
	}

	return 0, 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	planets := map[string]*Planet{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ")")

		parent := maps.GetOrCompute(planets, parts[0], NewPlanet)
		child := maps.GetOrCompute(planets, parts[1], NewPlanet)

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	}

	return World{
		Root:    planets["COM"],
		Planets: planets,
	}
}
