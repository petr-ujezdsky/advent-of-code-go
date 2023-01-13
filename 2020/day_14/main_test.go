package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	assert.Equal(t, 4, len(items))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart01(items)
	assert.Equal(t, 165, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart01(items)
	assert.Equal(t, 9879607673316, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart02(items)
	assert.Equal(t, 208, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart02(items)
	assert.Equal(t, 0, result)
}
