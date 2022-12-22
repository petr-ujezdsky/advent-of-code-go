package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
)

type NodeType int

const (
	Open NodeType = iota
	Wall
)

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

const maxWidth = 150

type Step struct {
	Rotation    int
	Translation int
}

type Node struct {
	Neighbours [4]*Node
	Type       NodeType
}

type World struct {
	FirstNode *Node
	Steps     []Step
}

func DoWithInput(world World) int {
	return len(world.Steps)
}

func toNodeType(char rune) NodeType {
	if char == '.' {
		return Open
	}

	return Wall
}

func toRotation(char uint8) int {
	if char == 'R' {
		return 1
	}

	return -1
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	firstVerticalNodes := make([]*Node, maxWidth)
	lastVerticalNodes := make([]*Node, maxWidth)
	var firstNode *Node

	for scanner.Scan() && scanner.Text() != "" {
		var firstHorizontalNode *Node
		var lastHorizontalNode *Node

		row := scanner.Text()
		for i, char := range row {
			if char == ' ' {
				continue
			}

			neighbours := [4]*Node{}
			neighbours[Left] = lastHorizontalNode
			neighbours[Up] = lastVerticalNodes[i]

			node := &Node{
				Neighbours: neighbours,
				Type:       toNodeType(char),
			}

			if lastHorizontalNode != nil {
				lastHorizontalNode.Neighbours[Right] = node
			}

			if lastVerticalNodes[i] != nil {
				lastVerticalNodes[i].Neighbours[Down] = node
			}

			lastHorizontalNode = node
			lastVerticalNodes[i] = node

			if firstHorizontalNode == nil {
				firstHorizontalNode = node
			}

			if firstVerticalNodes[i] == nil {
				firstVerticalNodes[i] = node
			}

			if firstNode == nil {
				firstNode = node
			}
		}

		firstHorizontalNode.Neighbours[Left] = lastHorizontalNode
		lastHorizontalNode.Neighbours[Right] = firstHorizontalNode
	}

	for i, firstVerticalNode := range firstVerticalNodes {
		if firstVerticalNode == nil {
			break
		}

		lastVerticalNode := lastVerticalNodes[i]

		firstVerticalNode.Neighbours[Up] = lastVerticalNode
		lastVerticalNode.Neighbours[Down] = firstVerticalNode
	}

	scanner.Scan()

	stepRegex := regexp.MustCompile(`[RL]\d+`)
	stepParts := stepRegex.FindAllString(scanner.Text(), -1)

	steps := make([]Step, len(stepParts)+1)
	steps[0] = Step{
		Rotation:    0,
		Translation: utils.ExtractInts(scanner.Text(), false)[0],
	}

	for i, stepPart := range stepParts {
		steps[i+1] = Step{
			Rotation:    toRotation(stepPart[0]),
			Translation: utils.ExtractInts(stepPart, false)[0],
		}
	}

	return World{
		FirstNode: firstNode,
		Steps:     steps,
	}
}
