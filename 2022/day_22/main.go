package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"io"
	"regexp"
	"strings"
)

type NodeType int

const (
	Open NodeType = iota
	Wall
)

type Direction = int

const (
	Right Direction = iota
	Down
	Left
	Up
)

var steps = [4]utils.Vector2i{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
}

var footsteps = [4]string{
	">",
	"v",
	"<",
	"^",
}

type Step struct {
	Rotation    int
	Translation int
}

type Node struct {
	Neighbours [4]*Node
	Type       NodeType
	Position   utils.Vector2i
	FaceId     int
	Directions [4]*int
	Footstep   *string
}

func (n *Node) String() string {
	if n == nil {
		return " "
	}

	if n.Footstep != nil {
		return *n.Footstep
	}

	switch n.Type {
	case Open:
		return "."
	case Wall:
		return "#"
	}

	panic("Unknown node type")
}

//func (n *Node) StringFaceId() string {
//	if n == nil {
//		return " "
//	}
//
//	if n.Direction != nil {
//		crossNode := n.Neighbours[(*n.Direction+2)%4]
//		return fmt.Sprintf("%x", crossNode.FaceId)
//	}
//
//	return fmt.Sprintf("%x", n.FaceId)
//}

func (n *Node) StringDirections() string {
	if n == nil {
		return " "
	}

	sb := &strings.Builder{}
	for _, inputDirection := range []int{Right, Down, Left, Up} {
		outputDirection := n.Directions[inputDirection]
		if outputDirection != nil {
			sb.WriteString(fmt.Sprintf("%d:%d,", inputDirection, *outputDirection))
		}
	}
	if sb.Len() > 0 {
		return "[" + sb.String() + "]"
	}

	return fmt.Sprintf(" %x ", n.FaceId)
}

type Matrix = matrix.Matrix[*Node]

type World struct {
	FirstNode *Node
	Steps     []Step
	Matrix    Matrix
}

func Walk(world World) int {
	node := world.FirstNode
	direction := Right
	node.Footstep = &footsteps[direction]

	for _, step := range world.Steps {
		// rotation
		direction = (direction + step.Rotation + 4) % 4

		// translation
		for i := 0; i < step.Translation; i++ {
			nextNode := node.Neighbours[direction]
			if nextNode.Type == Wall {
				break
			}

			nextNode.Footstep = &footsteps[direction]
			node.Footstep = &footsteps[direction]

			// moved across the edge
			if node.FaceId != nextNode.FaceId && nextNode.Directions[direction] != nil {
				direction = *nextNode.Directions[direction]
			}

			node = nextNode
		}
	}

	f := "x"
	node.Footstep = &f

	fmt.Println(world.Matrix.StringFmtSeparator("", func(node *Node) string { return node.String() }))

	return 1000*node.Position.Y + 4*node.Position.X + direction
}

type Edge struct {
	From      utils.Vector2i
	Direction Direction
}

type PatchDef struct {
	Edge               Edge
	NewDirection       int
	OtherEdgeDirection Direction
}

func patchEdge(patch1, patch2 PatchDef, m Matrix, edgeLength int) {
	step1 := steps[patch1.Edge.Direction]
	step2 := steps[patch2.Edge.Direction]

	pos1 := patch1.Edge.From
	pos2 := patch2.Edge.From
	for i := 0; i < edgeLength; i++ {
		node1 := m.GetV(pos1)
		node2 := m.GetV(pos2)

		if node1 == nil || node2 == nil {
			fmt.Print("Nil!")
		}
		// connect nodes
		node1.Neighbours[patch1.OtherEdgeDirection] = node2
		node2.Neighbours[patch2.OtherEdgeDirection] = node1

		// set directions
		node1.Directions[patch2.OtherEdgeDirection] = &patch1.NewDirection
		node2.Directions[patch1.OtherEdgeDirection] = &patch2.NewDirection

		// move to next index
		pos1 = pos1.Add(step1)
		pos2 = pos2.Add(step2)
	}
}
func patchEdgesExample(m Matrix) {
	l := 4

	// 1
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2 * l, Y: l - 1},
			Direction: Up,
		},
		NewDirection:       Right,
		OtherEdgeDirection: Left,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2*l - 1, Y: l},
			Direction: Left,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, m, l)

	// 2
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2 * l, Y: 0},
			Direction: Right,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l - 1, Y: l},
			Direction: Left,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, m, l)

	// 3
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2*l - 1, Y: 2*l - 1},
			Direction: Left,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2 * l, Y: 2 * l},
			Direction: Down,
		},
		NewDirection:       Right,
		OtherEdgeDirection: Left,
	}, m, l)

	// 4
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2 * l, Y: 3*l - 1},
			Direction: Right,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l - 1, Y: 2*l - 1},
			Direction: Left,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, m, l)

	// 5
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 3 * l, Y: 2 * l},
			Direction: Right,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 3*l - 1, Y: 2*l - 1},
			Direction: Up,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, m, l)

	// 6
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 0, Y: l},
			Direction: Down,
		},
		NewDirection:       Right,
		OtherEdgeDirection: Left,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 4*l - 1, Y: 3*l - 1},
			Direction: Left,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, m, l)

	// 7
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 3*l - 1, Y: 0},
			Direction: Down,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 4*l - 1, Y: 3*l - 1},
			Direction: Up,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, m, l)
}
func patchEdgesMain(m Matrix) {
	l := 50

	// 1
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l, Y: 2*l - 1},
			Direction: Up,
		},
		NewDirection: Right,

		OtherEdgeDirection: Left,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l - 1, Y: 2 * l},
			Direction: Left,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, m, l)

	// 2
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l, Y: l - 1},
			Direction: Up,
		},
		NewDirection:       Right,
		OtherEdgeDirection: Left,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 0, Y: 2 * l},
			Direction: Down,
		},
		NewDirection:       Right,
		OtherEdgeDirection: Left,
	}, m, l)

	// 3
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l, Y: 0},
			Direction: Right,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 0, Y: 3 * l},
			Direction: Down,
		},
		NewDirection:       Right,
		OtherEdgeDirection: Left,
	}, m, l)

	// 4
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2 * l, Y: 0},
			Direction: Right,
		},
		NewDirection:       Down,
		OtherEdgeDirection: Up,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 0, Y: 4*l - 1},
			Direction: Right,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, m, l)

	// 5
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: l - 1, Y: 4*l - 1},
			Direction: Up,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2*l - 1, Y: 3*l - 1},
			Direction: Left,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, m, l)

	// 6
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2*l - 1, Y: 2 * l},
			Direction: Down,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 3*l - 1, Y: l - 1},
			Direction: Up,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, m, l)

	// 7
	patchEdge(PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2*l - 1, Y: l},
			Direction: Down,
		},
		NewDirection:       Left,
		OtherEdgeDirection: Right,
	}, PatchDef{
		Edge: Edge{
			From:      utils.Vector2i{X: 2 * l, Y: l - 1},
			Direction: Right,
		},
		NewDirection:       Up,
		OtherEdgeDirection: Down,
	}, m, l)
}

func Walk3D(world World, nodesPatcher func(matrix Matrix)) int {
	// patch edges
	nodesPatcher(world.Matrix)

	// standard walk
	return Walk(world)
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

func ParseInput(r io.Reader, edgeLength int) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	firstVerticalNodes := make([]*Node, 4*edgeLength)
	lastVerticalNodes := make([]*Node, 4*edgeLength)
	var firstNode *Node
	y := 0
	m := matrix.NewMatrix[*Node](4*edgeLength, 4*edgeLength)

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

			// face id
			xFace := i / edgeLength
			yFace := y / edgeLength
			faceId := xFace + 4*yFace

			node := &Node{
				Neighbours: neighbours,
				Type:       toNodeType(char),
				Position: utils.Vector2i{
					X: i + 1,
					Y: y + 1,
				},
				Directions: [4]*int{},
				FaceId:     faceId,
			}

			// store node, at zero based index
			m.SetV(node.Position.Subtract(utils.Vector2i{X: 1, Y: 1}), node)

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
		y++
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
		Matrix:    m,
	}
}
