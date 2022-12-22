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

// 1, [-3], 2, 3, -2, 0, 4
// -3 moves between -2 and 0:
//1, 2, 3, -2, [-3], 0, 4

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

func toNumbers(firstNode *Node, nodes []*Node) []int {
	numbers := make([]int, len(nodes))
	node := firstNode

	for i := 0; i < len(nodes); i++ {
		numbers[i] = node.Value
		node = node.Right
	}

	return numbers
}

func getNode(node *Node, steps, totalNodesCount int) *Node {
	//count := utils.Abs(steps) % (totalNodesCount - 1)
	count := (steps + 2*(totalNodesCount-1)) % (totalNodesCount - 1)

	for j := 0; j < count; j++ {
		node = node.Right
	}

	return node
}

func getNodeDirect(node *Node, steps, totalNodesCount int) *Node {
	count := steps % totalNodesCount

	for j := 0; j < count; j++ {
		node = node.Right
	}

	return node
}

func MixNumberForTest(i int, firstNode *Node, nodes []*Node) []int {
	MixNumber(nodes[i], firstNode, len(nodes))
	return toNumbers(firstNode, nodes)
}

func MixNumber(node, firstNode *Node, totalNodesCount int) {
	if node.Value == 0 {
		return
	}

	// find target node
	targetNode := getNode(node, node.Value, totalNodesCount)

	if targetNode == node {
		return
	}

	// remove source node
	if node == firstNode {
		firstNode = node.Right
	}

	node.Left.Right = node.Right
	node.Right.Left = node.Left

	// put it to the right of the target node
	node.Left = targetNode
	node.Right = targetNode.Right
	targetNode.Right.Left = node
	targetNode.Right = node
}

func MixNumbers(numbers []int) int {
	nodes, firstNode, zeroNode := toNodes(numbers)

	//printNodes(firstNode, nodes)
	for _, node := range nodes {
		MixNumber(node, firstNode, len(nodes))

		//printNodes(firstNode, nodes)
	}

	a := getNodeDirect(zeroNode, 1000, len(nodes)).Value
	b := getNodeDirect(zeroNode, 2000, len(nodes)).Value
	c := getNodeDirect(zeroNode, 3000, len(nodes)).Value

	return a + b + c
}

func ParseInput(r io.Reader) []int {
	return utils.ParseToIntsP(r)
}
