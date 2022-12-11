package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	assert.Equal(t, 4, len(monkeys))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := PlayKeepAway(monkeys)
	assert.Equal(t, 10605, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := PlayKeepAway(monkeys)
	assert.Equal(t, 111210, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := PlayKeepAway(monkeys)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := PlayKeepAway(monkeys)
	assert.Equal(t, 0, result)
}
