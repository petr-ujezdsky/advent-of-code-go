package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	assert.Equal(t, 7, len(numbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := MixNumbers(numbers, 1)
	assert.Equal(t, 3, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := MixNumbers(numbers, 1)
	assert.Equal(t, 14888, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := MixNumbersWithDecryptionKey(numbers)
	assert.Equal(t, 1623178306, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := MixNumbersWithDecryptionKey(numbers)
	assert.Equal(t, 3760092545849, result)
}

func TestMixNumberForTest(t *testing.T) {
	type args struct {
		i         int
		firstNode *Node
		nodes     []*Node
	}
	createArgs := func(i int, ints []int) args {
		nodes, first, _ := toNodes(ints)
		return args{
			i:         i,
			firstNode: first,
			nodes:     nodes,
		}
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: createArgs(2, []int{0, 0, 1, 0, 0, 0, 0}), want: []int{0, 0, 0, 1, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 2, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 2, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 3, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, 3, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 4, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, 0, 4}},
		{name: "", args: createArgs(2, []int{0, 0, 5, 0, 0, 0, 0}), want: []int{0, 5, 0, 0, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 6, 0, 0, 0, 0}), want: []int{0, 0, 6, 0, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 7, 0, 0, 0, 0}), want: []int{0, 0, 0, 7, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 8, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 8, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 9, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, 9, 0}},
		{name: "", args: createArgs(2, []int{0, 0, 10, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, 0, 10}},
		{name: "", args: createArgs(2, []int{0, 0, 11, 0, 0, 0, 0}), want: []int{0, 11, 0, 0, 0, 0, 0}},

		{name: "", args: createArgs(2, []int{0, 0, -1, 0, 0, 0, 0}), want: []int{0, -1, 0, 0, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, -2, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, 0, -2}},
		{name: "", args: createArgs(2, []int{0, 0, -3, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, -3, 0}},
		{name: "", args: createArgs(2, []int{0, 0, -4, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, -4, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, -5, 0, 0, 0, 0}), want: []int{0, 0, 0, -5, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, -6, 0, 0, 0, 0}), want: []int{0, 0, -6, 0, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, -7, 0, 0, 0, 0}), want: []int{0, -7, 0, 0, 0, 0, 0}},
		{name: "", args: createArgs(2, []int{0, 0, -8, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, 0, -8}},
		{name: "", args: createArgs(2, []int{0, 0, -9, 0, 0, 0, 0}), want: []int{0, 0, 0, 0, 0, -9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MixNumberForTest(tt.args.i, tt.args.firstNode, tt.args.nodes), "MixNumberForTest(%v, %v, %v)", tt.args.i, tt.args.firstNode, tt.args.nodes)
		})
	}
}
