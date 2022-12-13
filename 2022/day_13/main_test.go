package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	assert.Equal(t, 8, len(pairs))

	for _, pair := range pairs {
		fmt.Println(pair.Nodes[0])
		fmt.Println(pair.Nodes[1])
		fmt.Println()
	}
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	result := DoWithInput(pairs)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	result := DoWithInput(pairs)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	result := DoWithInput(pairs)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	result := DoWithInput(pairs)
	assert.Equal(t, 0, result)
}
