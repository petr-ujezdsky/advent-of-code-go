package day_18

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
)

type Node struct {
	Value                  int
	Left, Right            *Node
	PreviousLeaf, NextLeaf *Node
}

func NewNode(str string) (*Node, error) {
	node, _, _, err := parseNode([]byte(str), nil)
	return node, err
}

func (node Node) Magnitude() int {
	if node.Left == nil {
		return node.Value
	}

	return 3*node.Left.Magnitude() + 2*node.Right.Magnitude()
}

func (node *Node) FirstLeaf() *Node {
	if node.Left != nil {
		return node.Left.FirstLeaf()
	}

	return node
}

func (node *Node) LastLeaf() *Node {
	if node.Right != nil {
		return node.Right.LastLeaf()
	}

	return node
}

func (node *Node) LeafValues() []int {
	var values []int

	for leaf := node.FirstLeaf(); leaf != nil; leaf = leaf.NextLeaf {
		values = append(values, leaf.Value)
	}

	return values
}

func (node *Node) Explode(parent *Node, depth int) bool {
	if node.Left == nil {
		if depth >= 5 {
			left := parent.Left
			right := parent.Right

			// exploding
			// left
			if previous := left.PreviousLeaf; previous != nil {
				previous.Value += left.Value
				previous.NextLeaf = parent
				parent.PreviousLeaf = previous
			}

			// right
			if next := right.NextLeaf; next != nil {
				next.Value += right.Value
				next.PreviousLeaf = parent
				parent.NextLeaf = next
			}

			// make parent leaf with value 0
			parent.Value = 0
			parent.Left = nil
			parent.Right = nil

			return true
		}

		return false
	}

	if node.Left.Explode(node, depth+1) {
		return true
	}

	if node.Right.Explode(node, depth+1) {
		return true
	}

	return false
}

func (node *Node) Split(parent *Node) bool {
	if node.Left == nil {
		if node.Value >= 10 {

			left := &Node{Value: node.Value / 2}
			right := &Node{Value: (node.Value + 1) / 2}

			left.NextLeaf = right
			right.PreviousLeaf = left

			if previous := node.PreviousLeaf; previous != nil {
				previous.NextLeaf = left
				left.PreviousLeaf = previous
			}

			if next := node.NextLeaf; next != nil {
				next.PreviousLeaf = right
				right.NextLeaf = next
			}

			node.PreviousLeaf = nil
			node.NextLeaf = nil
			node.Value = 0
			node.Left = left
			node.Right = right

			return true
		}

		return false
	}

	if node.Left.Split(node) {
		return true
	}

	if node.Right.Split(node) {
		return true
	}

	return false
}

func (node *Node) Reduce() {
	for modified := true; modified; {
		modified = false

		if node.Explode(nil, 0) {
			// fmt.Printf("after explode:  %v\n", node)
			modified = true
			continue
		}

		if node.Split(nil) {
			// fmt.Printf("after split:    %v\n", node)
			modified = true
			continue
		}
	}
}

func Add(left, right *Node) *Node {
	// fmt.Printf("  %v\n", left)
	// fmt.Printf("+ %v\n", right)

	result := &Node{Left: left, Right: right}

	middleLeft := left.LastLeaf()
	middleRight := right.FirstLeaf()

	middleLeft.NextLeaf = middleRight
	middleRight.PreviousLeaf = middleLeft

	// fmt.Printf("after addition: %v\n", result)

	result.Reduce()

	// fmt.Printf("= %v\n\n", result)

	return result
}

func AddString(left, right string) (*Node, error) {
	leftNode, err := NewNode(left)
	if err != nil {
		return &Node{}, err
	}

	rightNode, err := NewNode(right)
	if err != nil {
		return &Node{}, err
	}

	return Add(leftNode, rightNode), nil
}

func Sum(nodes []*Node) *Node {
	var sum *Node = nil

	for _, node := range nodes {
		if sum == nil {
			sum = node
		} else {
			sum = Add(sum, node)
		}
	}

	return sum
}

// Strings must be used because the Nodes are mutated during Add
func MaxSumMagnitude(nodes []string) (int, error) {
	max := 0
	var nodeMax1, nodeMax2 string

	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1 == n2 {
				continue
			}

			s1, err := AddString(n1, n2)
			if err != nil {
				return 0, err
			}

			s2, err := AddString(n2, n1)
			if err != nil {
				return 0, err
			}

			magnitude := utils.Max(s1.Magnitude(), s2.Magnitude())

			max = utils.Max(max, magnitude)

			nodeMax1 = n1
			nodeMax2 = n2
		}
	}

	fmt.Printf("1: %v\n", nodeMax1)
	fmt.Printf("2: %v\n", nodeMax2)

	return max, nil
}

func isNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func parseNode(data []byte, lastLeaf *Node) (*Node, *Node, []byte, error) {
	var i int
	for i = 0; i < len(data); i++ {
		ch := data[i]

		if ch == '[' {
			left, lastLeaf, data, err := parseNode(data[i+1:], lastLeaf)
			if err != nil {
				return &Node{}, nil, nil, err
			}

			right, lastLeaf, data, err := parseNode(data, lastLeaf)
			if err != nil {
				return &Node{}, nil, nil, err
			}

			return &Node{Left: left, Right: right}, lastLeaf, data, nil
		}

		if isNumber(ch) {
			end := i + 1
			for ; end < len(data); end++ {
				if !isNumber(data[end]) {
					break
				}
			}

			// number value
			value, err := strconv.Atoi(string(data[i:end]))
			if err != nil {
				return &Node{}, nil, nil, err
			}

			valueNode := &Node{Value: value, PreviousLeaf: lastLeaf}

			if lastLeaf != nil {
				lastLeaf.NextLeaf = valueNode
			}

			return valueNode, valueNode, data[end:], nil
		}
	}

	return &Node{}, nil, nil, errors.New("unfinished input")
}

func (node *Node) String() string {
	if node.Left != nil {
		return fmt.Sprintf("[%v,%v]", node.Left.String(), node.Right.String())
	}

	return strconv.Itoa(node.Value)
}
