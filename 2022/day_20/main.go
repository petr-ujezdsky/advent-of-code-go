package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Node struct {
	Left, Right *Node
	Value       int
}

// 4, 5, 6, 1, 7, 8, 9
// 4, 5, 6, 7, 1, 8, 9

// 4, -2, 5, 6, 7,  8, 9
// 4,  5, 6, 7, 8, -2, 9

func toNodes(numbers []int) ([]*Node, *Node, *Node) {
	nodes := make([]*Node, len(numbers))
	var zero *Node

	for i, number := range numbers {
		node := &Node{
			Value: number,
		}

		nodes[i] = node

		if number == 0 {
			zero = node
		}
	}

	for i, node := range nodes {
		node.Left = nodes[(i+len(nodes)-1)%len(nodes)]
		node.Right = nodes[(i+1)%len(nodes)]
	}

	return nodes, nodes[0], zero
}

func printNodes(firstNode *Node, nodes []*Node) {
	node := firstNode
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("%v, ", node.Value)
		node = node.Right
	}
	fmt.Println()
}

func getNode(node *Node, steps, totalNodesCount int) *Node {
	count := utils.Abs(steps) % totalNodesCount

	for j := 0; j < count; j++ {
		if steps > 0 {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	return node
}

func MixNumbers(numbers []int) int {
	nodes, firstNode, zeroNode := toNodes(numbers)

	printNodes(firstNode, nodes)
	for _, node := range nodes {
		if node.Value == 0 {
			continue
		}

		// find target node
		targetNode := getNode(node, node.Value, len(nodes))

		if targetNode == node {
			continue
		}

		// remove source node
		//node.Left.Right, node.Right.Left = node.Right, node.Left
		if node == firstNode {
			firstNode = node.Right
		}

		node.Left.Right = node.Right
		node.Right.Left = node.Left

		// put it next to target node
		if node.Value > 0 {
			// to the right
			node.Left = targetNode
			node.Right = targetNode.Right
			targetNode.Right.Left = node
			targetNode.Right = node
		} else {
			// to the left
			node.Right = targetNode
			node.Left = targetNode.Left
			targetNode.Left.Right = node
			targetNode.Left = node
		}

		printNodes(firstNode, nodes)
	}

	a := getNode(zeroNode, 1000, len(nodes)).Value
	b := getNode(zeroNode, 2000, len(nodes)).Value
	c := getNode(zeroNode, 3000, len(nodes)).Value

	return a + b + c

}

func ParseInput(r io.Reader) []int {
	return utils.ParseToIntsP(r)
}
