package tree

import (
	"fmt"
	"strings"
)

type Node interface {
	//Children() iterators.Iterator[Node]
	Children() []Node
}

func PrintTree(root Node) string {
	return traversePreOrderRoot(root)
}

// see https://www.baeldung.com/java-print-binary-tree-diagram
func traversePreOrderChild(builder *strings.Builder, padding, pointer string, node Node, last bool) {
	if node == nil {
		return
	}
	children := node.Children()

	builder.WriteRune('\n')
	builder.WriteString(padding)
	builder.WriteString(pointer)

	builder.WriteString(fmt.Sprintf("%v", node))

	paddingBuilder := &strings.Builder{}
	paddingBuilder.WriteString(padding)

	if last {
		paddingBuilder.WriteString("   ")
	} else {
		paddingBuilder.WriteString("│  ")
	}

	paddingForBoth := paddingBuilder.String()

	//it := node.Children()
	//for it.HasNext() {
	//	traversePreOrder(builder, it.Next())
	//}

	for i, subNode := range children {
		pointer = "├──"

		last := i == len(children)-1
		if last {
			pointer = "└──"
		}

		traversePreOrderChild(builder, paddingForBoth, pointer, subNode, last)
	}
}

func traversePreOrderRoot(root Node) string {
	if root == nil {
		return ""
	}

	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v", root))

	children := root.Children()
	for i, child := range children {
		pointer := "├──"

		last := i == len(children)-1
		if last {
			pointer = "└──"
		}

		traversePreOrderChild(builder, "", pointer, child, last)
	}

	return builder.String()
}
