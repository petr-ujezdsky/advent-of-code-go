package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	scanner := ParseInput(reader)

	assert.Equal(t, 14, len(scanner.Readouts))

	assert.Equal(t, Vector2i{2, 18}, scanner.Readouts[0].Sensor)
	assert.Equal(t, Vector2i{-2, 15}, scanner.Readouts[0].NearestBeacon)

	assert.Equal(t, Vector2i{20, 1}, scanner.Readouts[13].Sensor)
	assert.Equal(t, Vector2i{15, 3}, scanner.Readouts[13].NearestBeacon)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	scanner := ParseInput(reader)

	result := NoBeaconPositionsCount(scanner, 10)
	assert.Equal(t, 26, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	scanner := ParseInput(reader)

	result := NoBeaconPositionsCount(scanner, 2000000)
	assert.Equal(t, 5299855, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	scanner := ParseInput(reader)

	result := NoBeaconPositionsCount(scanner, 10)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	scanner := ParseInput(reader)

	result := NoBeaconPositionsCount(scanner, 10)
	assert.Equal(t, 0, result)
}
