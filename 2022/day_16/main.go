package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
)

var valvesRegex = regexp.MustCompile(`[A-Z]{2}`)

type World struct {
	RootNode    *ValveNode
	AllNodes    []*ValveNode
	AllNodesMap map[string]*ValveNode
}

type ValveNode struct {
	Name     string
	FlowRate int
	Children []*ValveNode
}

func DoWithInput(world World) int {
	return 0
}

func getOrCreateNodes(names []string, nodes map[string]*ValveNode) []*ValveNode {
	resultNodes := make([]*ValveNode, len(names))

	for i, name := range names {
		name = name[0:2]
		if node, ok := nodes[name]; ok {
			resultNodes[i] = node
		} else {
			node = &ValveNode{Name: name}
			resultNodes[i] = node
			nodes[name] = node
		}
	}

	return resultNodes
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	nodes := make(map[string]*ValveNode)

	for scanner.Scan() {
		flowRate := utils.ExtractInts(scanner.Text(), false)[0]
		names := valvesRegex.FindAllString(scanner.Text(), -1)

		name := names[0]

		node := getOrCreateNodes([]string{name}, nodes)[0]

		node.FlowRate = flowRate
		node.Children = append(node.Children, getOrCreateNodes(names[1:], nodes)...)
	}

	return World{
		RootNode:    nodes["AA"],
		AllNodes:    utils.MapValues(nodes),
		AllNodesMap: nodes,
	}
}
