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

	assert.Equal(t, 15, len(monkeys))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := EvaluateRootMonkey(monkeys)
	assert.Equal(t, 152, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := EvaluateRootMonkey(monkeys)
	assert.Equal(t, 118565889858886, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := FindEqualityForRootMonkey(monkeys)
	assert.Equal(t, 301, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	monkeys := ParseInput(reader)

	result := FindEqualityForRootMonkey(monkeys)
	assert.Equal(t, 3032671800355, result)
}
