package main

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
)

type Component struct {
	Name       string
	Neighbours []*Component
}

type World struct {
	Components map[string]*Component
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func getOrCreateComponent(name string, components map[string]*Component) *Component {
	if component, ok := components[name]; ok {
		return component
	}

	component := &Component{Name: name}
	components[name] = component

	return component
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	components := make(map[string]*Component)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")

		name := parts[0]
		component := getOrCreateComponent(name, components)

		for _, neighbourName := range strings.Split(parts[1], " ") {
			neighbour := getOrCreateComponent(neighbourName, components)

			// link both ways
			component.Neighbours = append(component.Neighbours, neighbour)
			neighbour.Neighbours = append(neighbour.Neighbours, component)
		}
	}

	return World{Components: components}
}
