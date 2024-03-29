package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	stacks, ops := ParseInput(reader)

	assert.Equal(t, 3, len(stacks))
	assert.Equal(t, 4, len(ops))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	stacks, ops := ParseInput(reader)

	topCrates := MoveCratesByOps(stacks, ops, false)
	assert.Equal(t, "CMZ", topCrates)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	stacks, ops := ParseInput(reader)

	topCrates := MoveCratesByOps(stacks, ops, false)
	assert.Equal(t, "FJSRQCFTN", topCrates)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	stacks, ops := ParseInput(reader)

	topCrates := MoveCratesByOps(stacks, ops, true)
	assert.Equal(t, "MCD", topCrates)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	stacks, ops := ParseInput(reader)

	topCrates := MoveCratesByOps(stacks, ops, true)
	assert.Equal(t, "CJVLJQPHS", topCrates)
}
