package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	readouts := ParseInput(reader)

	assert.Equal(t, 14, len(readouts))

	assert.Equal(t, Vector2i{2, 18}, readouts[0].Sensor)
	assert.Equal(t, Vector2i{-2, 15}, readouts[0].NearestBeacon)

	assert.Equal(t, Vector2i{20, 1}, readouts[13].Sensor)
	assert.Equal(t, Vector2i{15, 3}, readouts[13].NearestBeacon)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	readouts := ParseInput(reader)

	result := DoWithInput(readouts)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	readouts := ParseInput(reader)

	result := DoWithInput(readouts)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	readouts := ParseInput(reader)

	result := DoWithInput(readouts)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	readouts := ParseInput(reader)

	result := DoWithInput(readouts)
	assert.Equal(t, 0, result)
}
