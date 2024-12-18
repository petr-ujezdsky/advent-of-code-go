package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Rule struct {
	Left, Right int
}

type RuleNode struct {
	Value int
	//LeftNode, RightNode *RuleNode
	//Previous, Next []*RuleNode
	Next []*RuleNode
}

type Update []int

type World struct {
	Rules   []Rule
	Updates []Update
}

func getOrCreateNode(value int, nodes map[int]*RuleNode) *RuleNode {
	if node, ok := nodes[value]; ok {
		return node
	}

	node := &RuleNode{
		Value: value,
		Next:  nil,
	}

	nodes[value] = node

	return node
}

func createGraph(rules []Rule) map[int]*RuleNode {
	nodes := make(map[int]*RuleNode)

	for _, rule := range rules {
		nodeLeft := getOrCreateNode(rule.Left, nodes)
		nodeRight := getOrCreateNode(rule.Right, nodes)

		// connect nodes
		nodeLeft.Next = append(nodeLeft.Next, nodeRight)
	}

	return nodes
}

//
//func findNode(nodes []*RuleNode, value int) *RuleNode {
//	for _, node := range nodes {
//		if node.Value == value {
//			return node
//		}
//
//		if found := findNode(node.Next, value); found != nil {
//			return found
//		}
//	}
//
//	return nil
//}

func findNode(nodes []*RuleNode, value int) *RuleNode {
	for _, node := range nodes {
		if node.Value == value {
			return node
		}
	}

	for _, node := range nodes {
		if found := findNode(node.Next, value); found != nil {
			return found
		}
	}

	return nil
}

func conformsRules(update Update, nodes map[int]*RuleNode) bool {
	node := nodes[update[0]]

	for _, value := range update[1:] {
		node = findNode(node.Next, value)
		if node == nil {
			return false
		}
	}

	return true
}

func DoWithInputPart01(world World) int {
	nodes := createGraph(world.Rules)

	middlesSum := 0
	for _, update := range world.Updates {
		if !conformsRules(update, nodes) {
			fmt.Printf("❌ %v\n", update)
			continue
		}
		fmt.Printf("✅ %v\n", update)

		middle := update[len(update)/2]
		middlesSum += middle
	}

	return middlesSum
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// Rules
	var rules []Rule
	for scanner.Scan() && len(scanner.Text()) > 0 {
		ints := utils.ExtractInts(scanner.Text(), false)
		rules = append(rules, Rule{
			Left:  ints[0],
			Right: ints[1],
		})
	}

	// Updates
	var updates []Update
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)
		updates = append(updates, ints)
	}

	return World{
		Rules:   rules,
		Updates: updates,
	}
}
