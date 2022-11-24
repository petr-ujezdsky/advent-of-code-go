package day_14

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, "NNCB", world.template)
	assert.Equal(t, 16, len(world.rules))
	assert.Equal(t, "B", world.rules[duoHash([]rune("CH"))])
	assert.Equal(t, "C", world.rules[duoHash([]rune("CN"))])
}

func Test_01_example_single(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := world.template
	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NCNBCHB", polymer)

	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NBCCNBBBCBHCB", polymer)

	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", polymer)

	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer)
}

func Test_01_example_multi(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := GrowPolymerIter(world.template, world.rules, 4)
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 10)

	assert.Equal(t, 1588, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 10)

	assert.Equal(t, 3555, score)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 40)
	assert.Equal(t, 2188189693529, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 40)

	assert.Equal(t, 4439442043739, score)
}

// Benchmark_recursive_runified_caching-10    	   17024	      70 349 ns/op
func Benchmark_recursive_runified_caching(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world, err := ParseInput(reader)
	assert.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		score := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 18)
		assert.Equal(b, 480563, score)
	}
}
