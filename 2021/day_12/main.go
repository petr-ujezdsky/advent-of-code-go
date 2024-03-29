package day_12

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Type int

const (
	start Type = iota
	small
	big
	end
)

type Node struct {
	id         string
	nodeType   Type
	neighbours []*Node
}

func (node *Node) String() string {
	return node.id
}

type World struct {
	startNode *Node
	nodes     map[string]*Node
}

func nodeType(id string) Type {
	if id == "start" {
		return start
	}

	if id == "end" {
		return end
	}

	firstChar, _ := utf8.DecodeRuneInString(id)
	if unicode.IsUpper(firstChar) {
		return big
	}

	return small
}

func getOrCreateNode(world *World, id string) *Node {
	node, ok := world.nodes[id]

	if !ok {
		// create new node
		node = &Node{
			id:         id,
			nodeType:   nodeType(id),
			neighbours: nil,
		}

		// store it
		world.nodes[id] = node

		// store start node
		if node.nodeType == start {
			world.startNode = node
		}
	}

	return node
}

func canRevisit(path []*Node) bool {
	counts := make(map[string]int)

	for _, pathNode := range path {
		if pathNode.nodeType == small {
			counts[pathNode.id]++

			if counts[pathNode.id] > 1 {
				return false
			}
		}
	}

	return true
}

func visitable(node *Node, path []*Node, allowSmallRevisit bool) bool {
	switch node.nodeType {
	case start:
		return false
	case end, big:
		return true
	case small:
		for _, pathNode := range path {
			if node == pathNode {
				if !allowSmallRevisit {
					// small node already visited -> can not visit again
					return false
				}

				return canRevisit(path)
			}
		}

		return true
	default:
		panic("Unknown type")
	}
}

func findAllPathsRecursive(node *Node, path []*Node, allPaths *[][]*Node, allowSmallRevisit bool) {
	// add myself to the path
	path = append(path, node)

	// end on end
	if node.nodeType == end {
		*allPaths = append(*allPaths, path)
		return
	}

	// inspect neighbours
	for _, neighbour := range node.neighbours {
		if visitable(neighbour, path, allowSmallRevisit) {
			findAllPathsRecursive(neighbour, path, allPaths, allowSmallRevisit)
		}
	}
}

func FindAllPaths(world World, allowSmallRevisit bool) [][]*Node {
	var paths [][]*Node
	var path []*Node

	findAllPathsRecursive(world.startNode, path, &paths, allowSmallRevisit)
	return paths
}

func ParseInput(r io.Reader) (World, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	world := World{
		nodes: make(map[string]*Node),
	}

	for scanner.Scan() {
		nodeIds := strings.Split(scanner.Text(), "-")

		left := getOrCreateNode(&world, nodeIds[0])
		right := getOrCreateNode(&world, nodeIds[1])

		left.neighbours = append(left.neighbours, right)
		right.neighbours = append(right.neighbours, left)

	}

	return world, scanner.Err()
}
