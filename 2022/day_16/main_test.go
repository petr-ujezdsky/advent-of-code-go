package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 10, len(world.AllNodes))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := FindMaxPressureReleaseStateMinMax(world)
	assert.Equal(t, 1651, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := FindMaxPressureReleaseStateMinMax(world)
	assert.Equal(t, 1944, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := FindMaxPressureReleaseStateMinMax(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := FindMaxPressureReleaseStateMinMax(world)
	assert.Equal(t, 0, result)
}

// Benchmark_01_example-10    	    6877	    167 797 ns/op
// Benchmark_01_example-10    	   14564	    103 739 ns/op
// Benchmark_01_example-10    	   22880	     52 456 ns/op
func Benchmark_01_example(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := FindMaxPressureReleaseStateMinMax(world)
		assert.Equal(b, 1651, result)
	}
}

// Benchmark_01_example_generalized-10    	   10000	    102 752 ns/op
func Benchmark_01_example_generalized(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := FindMaxPressureReleaseStateMinMaxGeneralized(world)
		assert.Equal(b, 1651, result)
	}
}

// Benchmark_01-10    	     271	   4 130 114 ns/op
func Benchmark_01(b *testing.B) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(b, err)

	world := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := FindMaxPressureReleaseStateMinMax(world)
		assert.Equal(b, 1944, result)
	}
}

// Benchmark_01_generalized-10    	     247	   4 459 456 ns/op
func Benchmark_01_generalized(b *testing.B) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(b, err)

	world := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := FindMaxPressureReleaseStateMinMaxGeneralized(world)
		assert.Equal(b, 1944, result)
	}
}
