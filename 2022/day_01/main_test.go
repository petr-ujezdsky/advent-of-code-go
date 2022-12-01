package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	elves := ParseInput(reader)

	max := FindMax(elves)
	assert.Equal(t, 24000, max)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	elves := ParseInput(reader)

	max := FindMax(elves)
	assert.Equal(t, 64929, max)
}
