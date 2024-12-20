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

type StringSet = map[string]struct{}

type Edge struct {
	C1, C2 *Component
	//Count         int
	OriginalEdges []*Edge
}

func (edge *Edge) other(component *Component) *Component {
	if edge.C1 == component {
		return edge.C2
	}

	if edge.C2 == component {
		return edge.C1
	}

	panic("Component not on edge")
}

func (edge *Edge) swap(component, cNew *Component) {
	if edge.C1 == component {
		edge.C1 = cNew
		return
	}

	if edge.C2 == component {
		edge.C2 = cNew
		return
	}

	panic("Component not on edge")
}

func (edge *Edge) String() string {
	return fmt.Sprintf("%v/%v", edge.C1.Name, edge.C2.Name)
}

type Component struct {
	Name       string
	Neighbours []*Component

	Edges map[string]*Edge
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

		e1 := maps.FirstValue(c1.Edges)
		e2 := maps.FirstValue(c2.Edges)

		diff = len(c1.Edges) - len(c2.Edges)
		if diff != 0 {
			panic(fmt.Sprintf("Non zero diff, %v vs. %v\n", len(c1.Edges), len(c2.Edges)))
		}

		if e1 != e2 {
			panic(fmt.Sprintf("Edge is not symmetric\n"))
		}

		if len(e1.OriginalEdges) == 3 {
			fmt.Printf("* #%4d min-cut size %v, %v\n", i, len(e1.OriginalEdges), e1.OriginalEdges)

			countL, countR := countLeftRightSizes(e1.OriginalEdges, world.Components)
			fmt.Printf("Left: %v, right: %v\n", countL, countR)

			return countL * countR
		} else {
			fmt.Printf("  #%4d min-cut size %v\n", i, len(e1.OriginalEdges))
		}
	}

	panic("Found no solution")
}

func countLeftRightSizes(edges []*Edge, components map[string]*Component) (int, int) {
	components = initComponents(components)

	// do the cut
	for _, edge := range edges {
		delete(edge.C1.Edges, edge.C2.Name)
		delete(edge.C2.Edges, edge.C1.Name)
	}

	// count vertices in graph
	visited := make(StringSet)
	visitAllVertices(edges[0].C1, visited)
	left := len(visited)
	right := len(components) - left

	return left, right
}

func visitAllVertices(c *Component, visited StringSet) {
	if _, ok := visited[c.Name]; ok {
		return
	}

	visited[c.Name] = struct{}{}

	for _, edge := range c.Edges {
		neighbour := edge.other(c)
		visitAllVertices(neighbour, visited)
	}
}

// karger implements Karger's algorithm
// see https://en.wikipedia.org/wiki/Karger%27s_algorithm
func karger(components map[string]*Component) (*Component, *Component) {
	sequence := 0

	// create shallow copy
	components = initComponents(components)

	for len(components) > 2 {
		// find random edge
		cr := randMapValue(components)
		edge := randMapValue(cr.Edges)

		c1 := edge.C1
		c2 := edge.C2

		cNew := &Component{
			Name:  strconv.Itoa(sequence),
			Edges: make(map[string]*Edge),
		}
		components[cNew.Name] = cNew

		// contract the edge
		delete(c1.Edges, c2.Name)
		delete(c2.Edges, c1.Name)
		relink(c1, cNew, components)
		relink(c2, cNew, components)

		sequence++
	}

	c12 := maps.Values(components)
	return c12[0], c12[1]
}

func randMapValue[K comparable, V any](m map[K]*V) *V {
	index := rand.IntN(len(m))

	i := 0
	for _, component := range m {
		if i == index {
			return component
		}
		i++
	}

	panic("Could not find random component")
}

func relink(component, cNew *Component, components map[string]*Component) {
	for _, edge := range component.Edges {
		neighbour := edge.other(component)

		delete(neighbour.Edges, component.Name)

		if edgeNew, ok := neighbour.Edges[cNew.Name]; ok {
			// merge original edges
			for _, originalEdge := range edge.OriginalEdges {
				edgeNew.OriginalEdges = append(edgeNew.OriginalEdges, originalEdge)
			}
			continue
		}

		edge.swap(component, cNew)
		neighbour.Edges[cNew.Name] = edge
		cNew.Edges[neighbour.Name] = edge
	}

	delete(components, component.Name)
}

func initComponents(components map[string]*Component) map[string]*Component {
	// create shallow copy
	components = maps.Copy(components)

	// rebuild internal maps
	for _, component := range components {
		component.Edges = make(map[string]*Edge)
	}

	for _, c1 := range components {
		for _, c2 := range c1.Neighbours {
			// edge exists, skip
			if _, ok := c1.Edges[c2.Name]; ok {
				continue
			}

			edge := &Edge{
				C1:            c1,
				C2:            c2,
				OriginalEdges: []*Edge{{C1: c1, C2: c2}},
			}

			c1.Edges[c2.Name] = edge
			c2.Edges[c1.Name] = edge
		}
	}

	return components
}

func getOrCreateComponent(name string, components map[string]*Component, componentsList []*Component) (*Component, []*Component) {
	if component, ok := components[name]; ok {
		return component, componentsList
	}

	component := &Component{Name: name}

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
}
