package day_22

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_1(t *testing.T) {
	reader, err := os.Open("data-00-example-1.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 4, len(cubes))

	count := CountOnCubes(cubes)
	assert.Equal(t, 39, count)
}

func Test_01_example_2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 22, len(cubes))

	count := CountOnCubes(cubes[:20])
	assert.Equal(t, 590784, count)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 420, len(cubes))

	count := CountOnCubes(cubes[:20])
	assert.Equal(t, 620241, count)
}

func Test_02_example_3(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 60, len(cubes))

	count := CountOnCubes(cubes)
	assert.Equal(t, 2758514936282235, count)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 420, len(cubes))

	count := CountOnCubes(cubes)
	assert.Equal(t, 1284561759639324, count)
}

// Benchmark_CountOnCubes-10    	     556	   2 136 827 ns/op
func Benchmark_CountOnCubes(b *testing.B) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(b, err)

	cubes := ParseInput(reader)
	assert.Equal(b, 22, len(cubes))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := CountOnCubes(cubes[:20])
		assert.Equal(b, 590784, count)
	}
}
