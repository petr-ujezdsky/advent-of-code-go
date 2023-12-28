package tree

import (
	"fmt"
	"github.com/emirpasic/gods/sets/linkedhashset"
	"strings"
)

type Node interface {
	//Children() iterators.Iterator[Node]
	Children() []Node
}

func PrintTreeNode(root Node) string {
	return traversePreOrderRoot(root, func(node Node) (string, []Node) {
		return fmt.Sprintf("%v", node), node.Children()
	})
}

func PrintTree[T any](root T, extractor func(node T) (string, []T)) string {
	return traversePreOrderRoot(root, extractor)
}

// see https://www.baeldung.com/java-print-binary-tree-diagram
func traversePreOrderChild[T any](builder *strings.Builder, padding, pointer string, node T, last bool, path *linkedhashset.Set, extractor func(node T) (string, []T)) {
	//if node == nil {
	//	return
	//}
	name, children := extractor(node)

	builder.WriteRune('\n')
	builder.WriteString(padding)
	builder.WriteString(pointer)

	builder.WriteString(name)

	if path.Contains(name) {
		builder.WriteString(" --cycle detected--")
		return
	}
	path.Add(name)

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

		traversePreOrderChild(builder, paddingForBoth, pointer, subNode, last, path, extractor)
	}

	path.Remove(name)
}

func traversePreOrderRoot[T any](root T, extractor func(node T) (string, []T)) string {
	//if root == nil {
	//	return ""
	//}
	name, children := extractor(root)
	path := linkedhashset.New(name)

	builder := &strings.Builder{}
	builder.WriteString(name)

	for i, child := range children {
		pointer := "├──"

		last := i == len(children)-1
		if last {
			pointer = "└──"
		}

		traversePreOrderChild(builder, "", pointer, child, last, path, extractor)
	}

	return builder.String()
}
