package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strconv"
	"strings"
)

type Command struct {
	Name      string
	Args      []string
	Output    []string
	Evaluator Evaluator
}

type Type byte

const (
	File      Type = iota
	Directory      = iota
)

type Node struct {
	Parent *Node
	Nodes  []*Node

	Name string
	Type Type
	Size int
}

func (n *Node) AddNode(n2 *Node) {
	n.Nodes = append(n.Nodes, n2)

	// update sizes in parents
	if n2.Type == File {
		for cn := n; cn != nil; {
			cn.Size += n2.Size
			cn = cn.Parent
		}
	}
}

func (n *Node) string(sb *strings.Builder, prefix string) *strings.Builder {
	// write itself
	sb.WriteString(prefix + "- " + n.Name)
	if n.Type == Directory {
		sb.WriteString(" (dir)\n")
	} else {
		sb.WriteString(" (file, size=" + strconv.Itoa(n.Size) + ")\n")
	}

	// write children
	for _, node := range n.Nodes {
		node.string(sb, prefix+"  ")
	}

	return sb
}

func (n *Node) String() string {
	sb := &strings.Builder{}
	return n.string(sb, "").String()
}

type Evaluator func(currentNode *Node, command Command) *Node

var evaluators = map[string]Evaluator{
	"cd": cd,
	"ls": ls,
}

func cd(currentNode *Node, command Command) *Node {
	dirName := command.Args[0]

	// special dir
	if dirName == ".." {
		return currentNode.Parent
	}

	for _, node := range currentNode.Nodes {
		if node.Name == dirName {
			return node
		}
	}

	panic("Dir '" + dirName + "' not found!")
}

func ls(currentNode *Node, command Command) *Node {
	// parse output - create files and dirs
	for _, output := range command.Output {
		parts := strings.Split(output, " ")
		var node *Node

		if parts[0] == "dir" {
			node = &Node{
				Parent: currentNode,
				Nodes:  nil,
				Name:   parts[1],
				Type:   Directory,
				Size:   0,
			}
		} else {
			node = &Node{
				Parent: currentNode,
				Nodes:  nil,
				Name:   parts[1],
				Type:   File,
				Size:   utils.ParseInt(parts[0]),
			}
		}

		currentNode.AddNode(node)
	}
	return currentNode
}

func ReplayCommands(commands []Command) *Node {
	// skip first command (cd /)
	commands = commands[1:]

	// create root node
	root := &Node{
		Parent: nil,
		Nodes:  nil,
		Name:   "/",
		Type:   Directory,
		Size:   0,
	}

	// replay commands
	currentNode := root
	for _, command := range commands {
		currentNode = command.Evaluator(currentNode, command)
	}

	return root
}

func findDirsOfSizeLessThan(limit int, node *Node, dirs *[]*Node) {
	if node.Type == File {
		return
	}

	if node.Size <= limit {
		*dirs = append(*dirs, node)
	}

	for _, child := range node.Nodes {
		findDirsOfSizeLessThan(limit, child, dirs)
	}
}

func Filter100k(root *Node) int {
	limit := 100_000
	dirs := &[]*Node{}

	findDirsOfSizeLessThan(limit, root, dirs)

	// sum dir sizes
	sum := 0
	for _, dir := range *dirs {
		sum += dir.Size
	}

	return sum
}

func minMax(node *Node, limit int, min *Node) *Node {
	if node.Type == File || node.Size < limit {
		return min
	}

	if node.Size < min.Size {
		min = node
	}

	for _, child := range node.Nodes {
		min = minMax(child, limit, min)
	}

	return min
}

func Deletable(root *Node) int {
	free := 70_000_000 - root.Size
	needed := 30000000 - free

	min := minMax(root, needed, root)

	return min.Size
}

func ParseInput(r io.Reader) []Command {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var commands []Command

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		if parts[0] == "$" {
			command := Command{
				Name:      parts[1],
				Args:      parts[2:],
				Evaluator: evaluators[parts[1]],
			}

			commands = append(commands, command)
			continue
		}

		currentCommand := &commands[len(commands)-1]

		currentCommand.Output = append(currentCommand.Output, scanner.Text())
	}

	return commands
}
