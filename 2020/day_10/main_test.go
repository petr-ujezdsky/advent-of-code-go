package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	assert.Equal(t, 11, len(adapters))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	result := DoWithInput(adapters)
	assert.Equal(t, 7*5, result)
}

func Test_01_example2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	result := DoWithInput(adapters)
	assert.Equal(t, 22*10, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	result := DoWithInput(adapters)
	assert.Equal(t, 2170, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	result := DoWithInput2(adapters)
	assert.Equal(t, 8, result)
}

func Test_02_example2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	result := DoWithInput2(adapters)
	assert.Equal(t, 19208, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	adapters := ParseInput(reader)

	result := DoWithInput2(adapters)
	assert.Equal(t, 24803586664192, result)
}
