package day_18

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Value                  int
	Left, Right            *Node
	PreviousLeaf, NextLeaf *Node
	FirstLeaf, LastLeaf    *Node
}

func NewNode(str string) (*Node, error) {
	node, _, err := parseNode([]byte(str), nil, nil)
	return node, err
}

func (node Node) Magnitude() int {
	if node.Left == nil {
		return node.Value
	}

	return 3*node.Left.Magnitude() + 2*node.Right.Magnitude()
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

		for node.Explode(nil, 0) {
			modified = true
		}

		for node.Split(nil) {
			modified = true
		}
	}
}

func Add(left, right *Node) *Node {
	result := &Node{Left: left, Right: right, FirstLeaf: left.FirstLeaf, LastLeaf: right.LastLeaf}

	middleLeft := left.LastLeaf
	middleRight := right.FirstLeaf

	middleLeft.NextLeaf = middleRight
	middleRight.PreviousLeaf = middleLeft

	result.Reduce()

	return result
}

func isNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func parseNode(data []byte, leftmost, rightmost *Node) (*Node, []byte, error) {
	var i int
	for i = 0; i < len(data); i++ {
		ch := data[i]

		if ch == '[' {
			left, data, err := parseNode(data[i+1:], leftmost, rightmost)
			if err != nil {
				return &Node{}, nil, err
			}

			right, data, err := parseNode(data, left.FirstLeaf, left.LastLeaf)
			if err != nil {
				return &Node{}, nil, err
			}

			return &Node{Left: left, Right: right, FirstLeaf: right.FirstLeaf, LastLeaf: right.LastLeaf}, data, nil
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
				return &Node{}, nil, err
			}

			valueNode := &Node{Value: value, PreviousLeaf: rightmost}

			if rightmost != nil {
				rightmost.NextLeaf = valueNode
			}

			if leftmost == nil {
				leftmost = valueNode
			}

			valueNode.FirstLeaf = leftmost
			valueNode.LastLeaf = valueNode

			return valueNode, data[end:], nil
		}
	}

	return &Node{}, nil, errors.New("Unfinished input")
}

func (node *Node) String() string {
	if node.Left != nil {
		return fmt.Sprintf("[%v,%v]", node.Left.String(), node.Right.String())
	}

	return strconv.Itoa(node.Value)
}
