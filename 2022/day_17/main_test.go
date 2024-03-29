package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	assert.Equal(t, 40, len(jetDirections))
	assert.Equal(t, "[1 1 1 -1 -1 1 -1 1 1 -1 -1 -1 1 1 -1 1 1 1 -1 -1 -1 1 1 1 -1 -1 -1 1 -1 -1 -1 1 1 -1 1 1 -1 -1 1 1]", fmt.Sprint(jetDirections))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections, 2022)
	assert.Equal(t, 3068, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)
	assert.Equal(t, 10091, len(jetDirections))

	result := InspectFallingRocks(jetDirections, 2022)
	assert.Equal(t, 3227, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	// this never finishes - it seems that there are no repeated "floors" in example data!
	result := InspectFallingRocks(jetDirections, 1_000_000_000_000)
	assert.Equal(t, 0, result)
}

func Test_02_smaller(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections, 20_000_000)
	assert.Equal(t, 31954267, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections, 1_000_000_000_000)
	assert.Equal(t, 1597714285698, result)
}

// Benchmark_02-10    	     216	   5 516 270 ns/op
func Benchmark_02(b *testing.B) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(b, err)

	jetDirections := ParseInput(reader)
	metric.Disable()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := InspectFallingRocks(jetDirections, 1_000_000_000_000)
		assert.Equal(b, 1597714285698, result)
	}
}
