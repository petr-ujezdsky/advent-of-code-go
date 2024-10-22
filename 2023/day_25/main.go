package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"io"
	"math/rand/v2"
	"strconv"
	"strings"
)

type Component struct {
	Name       string
	Neighbours []*Component

	NeighboursMap map[string]*Component
	MergedCount   int
}

func (c Component) NeighbourNamesString() string {
	neighbourNames := make([]string, len(c.Neighbours))

	for i, neighbour := range c.Neighbours {
		neighbourNames[i] = neighbour.Name
	}

	return strings.Join(neighbourNames, " ")
}

func (c Component) String() string {
	return fmt.Sprintf("%v: %v", c.Name, c.NeighbourNamesString())
}

type World struct {
	Components     map[string]*Component
	ComponentsList []*Component
}

func DoWithInputPart01(world World) int {
	diff := 0
	for i := 0; i < 1000; i++ {
		c1, c2 := karger(world.Components)

		diff = len(c1.NeighboursMap) - len(c2.NeighboursMap)
		if diff != 0 {
			panic(fmt.Sprintf("Non zero diff, %v vs. %v\n", len(c1.NeighboursMap), len(c2.NeighboursMap)))
		}

		fmt.Printf("#%4d Cut edges %v\n", i, len(c1.NeighboursMap))
	}

	return diff
}

func karger(components map[string]*Component) (*Component, *Component) {
	sequence := 0

	// create shallow copy
	components = initComponents(components)

	for len(components) > 2 {
		// find random edge
		c1 := randComponent(components)
		c2 := randComponentUnique(c1.NeighboursMap, c1)

		// contract the edge
		delete(c1.NeighboursMap, c2.Name)
		delete(c2.NeighboursMap, c1.Name)
		c1Neighbours := unregister(c1, components)
		c2Neighbours := unregister(c2, components)

		cNew := &Component{
			Name:          strconv.Itoa(sequence),
			NeighboursMap: mergeMaps(c1Neighbours, c2Neighbours),
		}

		// link new component
		for _, neighbour := range cNew.NeighboursMap {
			neighbour.NeighboursMap[cNew.Name] = cNew
		}
		components[cNew.Name] = cNew

		sequence++
	}

	c12 := maps.Values(components)
	return c12[0], c12[1]
}

func randComponentUnique(components map[string]*Component, component *Component) *Component {
	for {
		rc := randComponent(components)
		if rc != component {
			return rc
		}
	}
}

func randComponent(components map[string]*Component) *Component {
	index := rand.IntN(len(components))

	i := 0
	for _, component := range components {
		if i == index {
			return component
		}
		i++
	}

	panic("Could not find random component")
}

func mergeMaps(m1, m2 map[string]*Component) map[string]*Component {
	result := make(map[string]*Component)

	for name, component := range m1 {
		result[name] = component
	}

	for name, component := range m2 {
		result[name] = component
	}

	return result
}

func unregister(component *Component, components map[string]*Component) map[string]*Component {
	for _, neighbour := range component.NeighboursMap {
		delete(neighbour.NeighboursMap, component.Name)
	}

	delete(components, component.Name)

	neighbours := component.NeighboursMap
	component.NeighboursMap = nil

	return neighbours
}

func initComponents(components map[string]*Component) map[string]*Component {
	// create shallow copy
	components = maps.Copy(components)

	// rebuild internal maps
	for _, component := range components {
		component.NeighboursMap = make(map[string]*Component)
		component.MergedCount = 0
	}

	for _, c1 := range components {
		for _, c2 := range c1.Neighbours {
			c1.NeighboursMap[c2.Name] = c2
			c2.NeighboursMap[c1.Name] = c1
		}
	}

	return components
}

func DoWithInputPart02(world World) int {
	return 0
}

func getOrCreateComponent(name string, components map[string]*Component, componentsList []*Component) (*Component, []*Component) {
	if component, ok := components[name]; ok {
		return component, componentsList
	}

	component := &Component{
		Name:          name,
		NeighboursMap: make(map[string]*Component),
	}

	components[name] = component

	componentsList = append(componentsList, component)

	return component, componentsList
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	components := make(map[string]*Component)
	var componentsList []*Component
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")

		name := parts[0]
		var component *Component
		component, componentsList = getOrCreateComponent(name, components, componentsList)

		for _, neighbourName := range strings.Split(parts[1], " ") {
			var neighbour *Component
			neighbour, componentsList = getOrCreateComponent(neighbourName, components, componentsList)

			// link both ways
			connect(component, neighbour)
		}
	}

	return World{
		Components:     components,
		ComponentsList: componentsList,
	}
}

func connect(c1, c2 *Component) {
	c1.Neighbours = append(c1.Neighbours, c2)
	c2.Neighbours = append(c2.Neighbours, c1)

	c1.NeighboursMap[c2.Name] = c2
	c2.NeighboursMap[c1.Name] = c1
}
