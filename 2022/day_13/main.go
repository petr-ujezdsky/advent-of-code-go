package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
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

func DoWithInput(pairs []NodePair) int {
	return len(pairs)
}

func ParseNode(chars []rune) *Node {
	children := utils.NewStack[*Node]()

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

func ParseInput(r io.Reader) []NodePair {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var pairs []NodePair
	i := 0
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
		i++
	}

	return pairs
}
