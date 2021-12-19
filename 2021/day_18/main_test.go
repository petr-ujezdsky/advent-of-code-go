package day_18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	// reader, err := os.Open("data-00-example.txt")
	// assert.Nil(t, err)

	// commands, err := ParseToCommands(reader)
	// assert.Nil(t, err)

	// x, y, result := move(commands)

	// assert.Equal(t, 15, x)
	// assert.Equal(t, 10, y)
	// assert.Equal(t, 150, result)
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

	assert.Equal(t, 1, node.LastLeaf.Value)
	assert.Equal(t, 9, node.LastLeaf.PreviousLeaf.Value)

	assert.Equal(t, node.FirstLeaf, node.LastLeaf.PreviousLeaf)
}

func TestNewNode2(t *testing.T) {
	node, err := NewNode("[9,[8,7]]")
	assert.Nil(t, err)

	assert.Equal(t, 9, node.Left.Value)
	assert.Equal(t, 8, node.Right.Left.Value)
	assert.Equal(t, 7, node.Right.Right.Value)

	assert.Equal(t, 7, node.LastLeaf.Value)
	assert.Equal(t, 8, node.LastLeaf.PreviousLeaf.Value)
	assert.Equal(t, 9, node.LastLeaf.PreviousLeaf.PreviousLeaf.Value)

	assert.Equal(t, node.FirstLeaf, node.LastLeaf.PreviousLeaf.PreviousLeaf)
}

func TestNewNode3(t *testing.T) {
	node, err := NewNode("[[9,8],7]")
	assert.Nil(t, err)

	assert.Equal(t, 9, node.Left.Left.Value)
	assert.Equal(t, 8, node.Left.Right.Value)
	assert.Equal(t, 7, node.Right.Value)

	assert.Equal(t, 7, node.LastLeaf.Value)
	assert.Equal(t, 8, node.LastLeaf.PreviousLeaf.Value)
	assert.Equal(t, 9, node.LastLeaf.PreviousLeaf.PreviousLeaf.Value)

	assert.Equal(t, node.FirstLeaf, node.LastLeaf.PreviousLeaf.PreviousLeaf)
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
}

func TestExplode2(t *testing.T) {
	node, err := NewNode("[7,[6,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[7,[6,[5,[7,0]]]]", node.String())
}

func TestExplode3(t *testing.T) {
	node, err := NewNode("[[6,[5,[4,[3,2]]]],1]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[6,[5,[7,0]]],3]", node.String())
}

func TestExplode4(t *testing.T) {
	node, err := NewNode("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", node.String())
}

func TestExplode5(t *testing.T) {
	node, err := NewNode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", node.String())
}

func TestSplit1(t *testing.T) {
	node, err := NewNode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	assert.Nil(t, err)
	exploded := node.Explode(nil, 0)

	assert.True(t, exploded)
	assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", node.String())
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
