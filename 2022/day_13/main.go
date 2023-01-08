package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Children []*Node
	Value    *int
}

func (n *Node) String() string {
	sb := &strings.Builder{}

	if n.Value != nil {
		return strconv.Itoa(*n.Value)
	}

	sb.WriteRune('[')
	for i, child := range n.Children {
		sb.WriteString(child.String())
		if i != len(n.Children)-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune(']')

	return sb.String()
}

type NodePair struct {
	Index int
	Nodes [2]*Node
}

// compareNodes compares two nodes
//
//	0: if (n1 = n2)
//
// -1: if (n1 < n2)
//
//	1: if (n1 > n2)
func compareNodes(n1, n2 *Node) int {
	// both nodes are values
	if n1.Value != nil && n2.Value != nil {
		return utils.Signum(*n1.Value - *n2.Value)
	}

	// both nodes are lists
	if n1.Value == nil && n2.Value == nil {
		for i := 0; i < utils.Min(len(n1.Children), len(n2.Children)); i++ {
			cmp := compareNodes(n1.Children[i], n2.Children[i])
			if cmp != 0 {
				return cmp
			}
		}

		return utils.Signum(len(n1.Children) - len(n2.Children))
	}

	// list vs value
	if n1.Value == nil && n2.Value != nil {
		n2 = &Node{
			Children: []*Node{n2},
		}
		return compareNodes(n1, n2)
	}

	// value vs list
	n1 = &Node{
		Children: []*Node{n1},
	}
	return compareNodes(n1, n2)
}

func FindInOrder(pairs []NodePair) int {
	indexSum := 0

	for _, pair := range pairs {
		cmp := compareNodes(pair.Nodes[0], pair.Nodes[1])
		if cmp < 0 {
			indexSum += pair.Index
		}
	}

	return indexSum
}

func FindDecoderKey(nodes []*Node) int {
	// add divider nodes
	divider1 := ParseNode([]rune("[[2]]"))
	divider2 := ParseNode([]rune("[[6]]"))
	nodes = append(nodes, divider1, divider2)

	// sort
	sort.Slice(nodes, func(i, j int) bool { return compareNodes(nodes[i], nodes[j]) < 0 })

	key := 1
	for i, node := range nodes {
		if node == divider1 || node == divider2 {
			key *= i + 1
		}
	}

	return key
}

func ParseNode(chars []rune) *Node {
	children := collections.NewStack[*Node]()

	// root node
	children.Push(&Node{})

	var digits []rune

	for i := 0; i < len(chars); i++ {
		char := chars[i]

		if char == ',' || char == ']' {
			if len(digits) > 0 {
				value := utils.ParseInt(string(digits))

				node := &Node{
					Children: nil,
					Value:    &value,
				}

				children.Peek().Children = append(children.Peek().Children, node)
				digits = []rune{}
			}

			if char == ',' {
				continue
			}
		}

		if char == '[' {
			node := &Node{}
			children.Peek().Children = append(children.Peek().Children, node)
			children.Push(node)
			continue
		}

		if char == ']' {
			children.Pop()
			continue
		}

		// char is digit
		digits = append(digits, char)
	}

	return children.Pop().Children[0]
}

func ParseInput(r io.Reader) ([]NodePair, []*Node) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var pairs []NodePair
	var nodes []*Node

	i := 1
	for scanner.Scan() {
		node1 := ParseNode([]rune(scanner.Text()))
		scanner.Scan()
		node2 := ParseNode([]rune(scanner.Text()))
		scanner.Scan()

		pair := NodePair{
			Index: i,
			Nodes: [2]*Node{node1, node2},
		}

		pairs = append(pairs, pair)
		nodes = append(nodes, node1, node2)
		i++
	}

	return pairs, nodes
}
