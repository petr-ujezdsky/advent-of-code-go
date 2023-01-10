package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	seatsOccupancy := ParseInput(reader)

	assert.Equal(t, 71, len(seatsOccupancy))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	seatsOccupancy := ParseInput(reader)

	result := DoWithInputPart01(seatsOccupancy)
	assert.Equal(t, 37, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	seatsOccupancy := ParseInput(reader)

	result := DoWithInputPart01(seatsOccupancy)
	assert.Equal(t, 2113, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	seatsOccupancy := ParseInput(reader)

	result := DoWithInputPart02(seatsOccupancy)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	seatsOccupancy := ParseInput(reader)

	result := DoWithInputPart02(seatsOccupancy)
	assert.Equal(t, 0, result)
}
