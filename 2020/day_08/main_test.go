package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	operations := ParseInput(reader)

	assert.Equal(t, 9, len(operations))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	operations := ParseInput(reader)

	result := ValueBeforeCycle(operations)
	assert.Equal(t, 5, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	operations := ParseInput(reader)

	result := ValueBeforeCycle(operations)
	assert.Equal(t, 1753, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	operations := ParseInput(reader)

	result := ValueBeforeCycle(operations)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	operations := ParseInput(reader)

	result := ValueBeforeCycle(operations)
	assert.Equal(t, 0, result)
}
