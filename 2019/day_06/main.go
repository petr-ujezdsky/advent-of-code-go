package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
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
	youPlanet := world.Planets["YOU"]
	santaPlanet := world.Planets["SAN"]

	youPath := slices.Reverse(pathToRoot(youPlanet))
	fmt.Printf("YOU path %v\n", slices.Map(youPath, func(planet *Planet) string { return planet.Name }))

	santaPath := slices.Reverse(pathToRoot(santaPlanet))
	fmt.Printf("SAN path %v\n", slices.Map(santaPath, func(planet *Planet) string { return planet.Name }))

	for i, youPlanet := range youPath {
		santaPlanet := santaPath[i]

		if youPlanet.Name != santaPlanet.Name {
			return len(youPath) + len(santaPath) - 2*i
		}
	}

	return len(youPath)
}

func pathToRoot(planet *Planet) []*Planet {
	var path []*Planet

	parent := planet.Parent

	for parent != nil {
		path = append(path, parent)
		parent = parent.Parent
	}

	return path
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
