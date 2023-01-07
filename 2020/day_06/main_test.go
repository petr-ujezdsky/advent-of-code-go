package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	groups := ParseInput(reader)

	assert.Equal(t, 5, len(groups))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	groups := ParseInput(reader)

	result := CountTrueAnswersPerGroup(groups)
	assert.Equal(t, 11, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	groups := ParseInput(reader)

	result := CountTrueAnswersPerGroup(groups)
	assert.Equal(t, 6249, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	groups := ParseInput(reader)

	result := CountTrueAnswersPerGroup(groups)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	groups := ParseInput(reader)

	result := CountTrueAnswersPerGroup(groups)
	assert.Equal(t, 0, result)
}
