package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)

	assert.Equal(t, 13, len(cubes))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)

	result := SurfaceArea(cubes)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)

	result := SurfaceArea(cubes)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)

	result := SurfaceArea(cubes)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)

	result := SurfaceArea(cubes)
	assert.Equal(t, 0, result)
}
