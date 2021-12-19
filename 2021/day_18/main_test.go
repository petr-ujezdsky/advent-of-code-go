package day_18

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	nodes, err := ParseToNodes(reader)
	assert.Nil(t, err)

	sum := Sum(nodes)

	assert.Equal(t, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", sum.String())
	assert.Equal(t, 4140, sum.Magnitude())
}

func Test_01(t *testing.T) {
	// reader, err := os.Open("data-01.txt")
	// assert.Nil(t, err)

	// commands, err := ParseToCommands(reader)
	// assert.Nil(t, err)

	// x, y, result := move(commands)

	// assert.Equal(t, 1968, x)
	// assert.Equal(t, 1063, y)
	// assert.Equal(t, 2091984, result)
}

func Test_02_example(t *testing.T) {
	// reader, err := os.Open("data-00-example.txt")
	// assert.Nil(t, err)

	// commands, err := ParseToCommands(reader)
	// assert.Nil(t, err)

	// x, y, result := moveByAim(commands)

	// assert.Equal(t, 15, x)
	// assert.Equal(t, 60, y)
	// assert.Equal(t, 900, result)
}

func Test_02(t *testing.T) {
	// assert.Equal(t, 2086261056, result)
}

func TestMagnitude(t *testing.T) {
	node1 := Node{Left: &Node{Value: 9}, Right: &Node{Value: 1}}
	assert.Equal(t, 29, node1.Magnitude())

	node2 := Node{Left: &Node{Value: 1}, Right: &Node{Value: 9}}
	assert.Equal(t, 21, node2.Magnitude())

	node3 := Node{Left: &node1, Right: &node2}
	assert.Equal(t, 129, node3.Magnitude())
}

func TestNewNode(t *testing.T) {
	node, err := NewNode("[9,1]")
	assert.Nil(t, err)

	assert.Equal(t, 9, node.Left.Value)
	assert.Equal(t, 1, node.Right.Value)

	assert.Equal(t, []int{9, 1}, node.LeafValues())
}

func TestNewNode2(t *testing.T) {
	node, err := NewNode("[9,[8,7]]")
	assert.Nil(t, err)

	assert.Equal(t, 9, node.Left.Value)
	assert.Equal(t, 8, node.Right.Left.Value)
	assert.Equal(t, 7, node.Right.Right.Value)

	assert.Equal(t, []int{9, 8, 7}, node.LeafValues())
}

func TestNewNode3(t *testing.T) {
	node, err := NewNode("[[9,8],7]")
	assert.Nil(t, err)

	assert.Equal(t, 9, node.Left.Left.Value)
	assert.Equal(t, 8, node.Left.Right.Value)
	assert.Equal(t, 7, node.Right.Value)

	assert.Equal(t, []int{9, 8, 7}, node.LeafValues())
}

func TestToString(t *testing.T) {
	node, err := NewNode("[9,1]")
	assert.Nil(t, err)

	assert.Equal(t, "[9,1]", node.String())
}

func TestToString2(t *testing.T) {
	node, err := NewNode("[9,[8,7]]")
	assert.Nil(t, err)

	assert.Equal(t, "[9,[8,7]]", node.String())
}

func TestToString3(t *testing.T) {
	node, err := NewNode("[[9,8],7]")
	assert.Nil(t, err)

	assert.Equal(t, "[[9,8],7]", node.String())
}

func TestExplode1(t *testing.T) {
	node, err := NewNode("[[[[[9,8],1],2],3],4]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[[[0,9],2],3],4]", node.String())

	assert.Equal(t, []int{0, 9, 2, 3, 4}, node.LeafValues())
}

func TestExplode2(t *testing.T) {
	node, err := NewNode("[7,[6,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[7,[6,[5,[7,0]]]]", node.String())

	assert.Equal(t, []int{7, 6, 5, 7, 0}, node.LeafValues())
}

func TestExplode3(t *testing.T) {
	node, err := NewNode("[[6,[5,[4,[3,2]]]],1]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[6,[5,[7,0]]],3]", node.String())

	assert.Equal(t, []int{6, 5, 7, 0, 3}, node.LeafValues())
}

func TestExplode4(t *testing.T) {
	node, err := NewNode("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", node.String())

	assert.Equal(t, []int{3, 2, 8, 0, 9, 5, 4, 3, 2}, node.LeafValues())
}

func TestExplode5(t *testing.T) {
	node, err := NewNode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", node.String())
}

func TestSplit1(t *testing.T) {
	node, err := NewNode("[1,10]")
	assert.Nil(t, err)
	splitted := node.Split(nil)

	assert.True(t, splitted)
	assert.Equal(t, "[1,[5,5]]", node.String())

	assert.Equal(t, []int{1, 5, 5}, node.LeafValues())
}

func TestSplit2(t *testing.T) {
	node, err := NewNode("[10,1]")
	assert.Nil(t, err)
	splitted := node.Split(nil)

	assert.True(t, splitted)
	assert.Equal(t, "[[5,5],1]", node.String())

	assert.Equal(t, []int{5, 5, 1}, node.LeafValues())
}

func TestSplit3(t *testing.T) {
	node, err := NewNode("[1,11]")
	assert.Nil(t, err)
	splitted := node.Split(nil)

	assert.True(t, splitted)
	assert.Equal(t, "[1,[5,6]]", node.String())

	assert.Equal(t, []int{1, 5, 6}, node.LeafValues())
}

func TestAdd1(t *testing.T) {
	left, err := NewNode("[[[[4,3],4],4],[7,[[8,4],9]]]")
	assert.Nil(t, err)

	right, err := NewNode("[1,1]")
	assert.Nil(t, err)

	result := Add(left, right)

	assert.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", result.String())
}

func TestAdd2(t *testing.T) {
	left, err := NewNode("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
	assert.Nil(t, err)

	right, err := NewNode("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]")
	assert.Nil(t, err)

	result := Add(left, right)

	assert.Equal(t, "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", result.String())
}

func TestAdd3(t *testing.T) {
	left, err := NewNode("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")
	assert.Nil(t, err)

	right, err := NewNode("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]")
	assert.Nil(t, err)

	result := Add(left, right)

	assert.Equal(t, "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]", result.String())
}

func TestSum1(t *testing.T) {
	reader := strings.NewReader(`[1,1]
	[2,2]
	[3,3]
	[4,4]`)

	nodes, err := ParseToNodes(reader)
	assert.Nil(t, err)

	sum := Sum(nodes)

	assert.Equal(t, "[[[[1,1],[2,2]],[3,3]],[4,4]]", sum.String())
}

func TestSum2(t *testing.T) {
	reader := strings.NewReader(`[1,1]
	[2,2]
	[3,3]
	[4,4]
	[5,5]`)

	nodes, err := ParseToNodes(reader)
	assert.Nil(t, err)

	sum := Sum(nodes)

	assert.Equal(t, "[[[[3,0],[5,3]],[4,4]],[5,5]]", sum.String())
}

func TestSum3(t *testing.T) {
	reader := strings.NewReader(`[1,1]
	[2,2]
	[3,3]
	[4,4]
	[5,5]
	[6,6]`)

	nodes, err := ParseToNodes(reader)
	assert.Nil(t, err)

	sum := Sum(nodes)

	assert.Equal(t, "[[[[5,0],[7,4]],[5,5]],[6,6]]", sum.String())
}

func TestSum4(t *testing.T) {
	reader := strings.NewReader(`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
	[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
	[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
	[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
	[7,[5,[[3,8],[1,4]]]]
	[[2,[2,2]],[8,[8,1]]]
	[2,9]
	[1,[[[9,3],9],[[9,0],[0,7]]]]
	[[[5,[7,4]],7],1]
	[[[[4,2],2],6],[8,7]]`)

	nodes, err := ParseToNodes(reader)
	assert.Nil(t, err)

	sum := Sum(nodes)

	assert.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", sum.String())
}

func ParseToNodes(r io.Reader) ([]*Node, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []*Node

	for scanner.Scan() {
		node, err := NewNode(scanner.Text())
		if err != nil {
			return result, err
		}

		result = append(result, node)
	}

	return result, scanner.Err()
}

// func ScanEntities(data []byte, atEOF bool) (advance int, token []byte, err error) {
// 	if atEOF && len(data) == 0 {
// 		return 0, nil, nil
// 	}

// 	ch := data[0]

// 	if ch == '[' {
// 		return 1, data[0:1], nil
// 	}

// 	start := 0

// 	for ; start < len(data); start++ {
// 		if IsNumber(data[start]) {
// 			break
// 		}
// 	}

// 	for end := start + 1; end < len(data); end++ {
// 		if !IsNumber(data[end]) {
// 			return end, data[start:end], nil
// 		}
// 	}
// 	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
// 	if atEOF && len(data) > start {
// 		return len(data), data[start:], nil
// 	}
// 	// Request more data.
// 	return start, nil, nil
// }

// func ParseNodeOld(scanner *bufio.Scanner) (Node, error) {
// 	scanner.Scan()
// 	entity := scanner.Text()

// 	if entity == "[" {
// 		left, err := ParseNodeOld(scanner)
// 		if err != nil {
// 			return Node{}, err
// 		}

// 		right, err := ParseNodeOld(scanner)
// 		if err != nil {
// 			return Node{}, err
// 		}

// 		return Node{Left: &left, Right: &right}, nil
// 	}

// 	// number value
// 	value, err := strconv.Atoi(entity)
// 	if err != nil {
// 		return Node{}, err
// 	}

// 	return Node{Value: value}, nil
// }

// func TestScanEntities(t *testing.T) {
// 	r := strings.NewReader("[[[[[9,8],1],2],3],4]")

// 	scanner := bufio.NewScanner(r)
// 	scanner.Split(ScanEntities)

// 	var result []string

// 	for scanner.Scan() {
// 		result = append(result, scanner.Text())
// 	}

// 	// expected := []string{"[", "[", "[", "[", "[", "9", ",", "8", "]", ",", "1", "]", ",", "2", "]", ",", "3", "]", ",", "4", "]"}
// 	expected := []string{"[", "[", "[", "[", "[", "9", "8", "1", "2", "3", "4"}

// 	assert.Equal(t, expected, result)
// }

// func TestNewNode(t *testing.T) {
// 	node, err := NewNode("[9,1]")
// 	assert.Nil(t, err)

// 	assert.Equal(t, 9, node.Left.Value)
// 	assert.Equal(t, 1, node.Right.Value)
// }

// func TestNewNode2(t *testing.T) {
// 	node, err := NewNode("[9,[8,7]]")
// 	assert.Nil(t, err)

// 	assert.Equal(t, 9, node.Left.Value)
// 	assert.Equal(t, 8, node.Right.Left.Value)
// 	assert.Equal(t, 7, node.Right.Right.Value)
// }

// func TestScanEntities2(t *testing.T) {
// 	r := strings.NewReader("8],1")

// 	scanner := bufio.NewScanner(r)
// 	scanner.Split(ScanEntities)

// 	var result []string

// 	for scanner.Scan() {
// 		result = append(result, scanner.Text())
// 	}

// 	expected := []string{"8", "1"}

// 	assert.Equal(t, expected, result)
// }
