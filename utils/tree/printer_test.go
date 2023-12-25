package tree

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type node struct {
	Name  string
	Nodes []*node
}

func (n *node) Children() []Node {
	children := make([]Node, len(n.Nodes))

	for i, child := range n.Nodes {
		children[i] = child
	}

	return children
}

//func (n *node) Children() iterators.Iterator[Node] {
//	it := iterators.NewSliceIterator[Node](n.Nodes)
//	return it
//}

func (n *node) String() string {
	return n.Name
}

func TestPrintTree(t *testing.T) {
	node9 := &node{Name: "node9"}
	node8 := &node{Name: "node8"}
	node7 := &node{Name: "node7", Nodes: []*node{node8, node9}}
	node6 := &node{Name: "node6"}
	node5 := &node{Name: "node5"}
	node4 := &node{Name: "node4"}
	node3 := &node{Name: "node3", Nodes: []*node{node7}}
	node2 := &node{Name: "node2", Nodes: []*node{node5, node6}}
	node1 := &node{Name: "node1", Nodes: []*node{node3, node4}}
	root := &node{Name: "root", Nodes: []*node{node1, node2}}

	expected := utils.Msg(`
root
├──node1
│  ├──node3
│  │  └──node7
│  │     ├──node8
│  │     └──node9
│  └──node4
└──node2
   ├──node5
   └──node6`)

	actual := PrintTree(root)

	assert.Equal(t, expected, actual)
}
