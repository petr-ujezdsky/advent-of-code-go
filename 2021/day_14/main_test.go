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
	assert.Equal(t, "B", world.rules["CH"])
	assert.Equal(t, "C", world.rules["CN"])
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

	polymer := GrowPolymerIter(world.template, world.rules, 10)
	score := PolymerScore(polymer)
	assert.Equal(t, 1588, score)
}

func Test_01_example_recursive(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score := GrowPolymerRecursive(world.template, world.rules, 10)
	assert.Equal(t, 1588, score)
}

func Test_01_example_recursive_runified(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	worldRunes := Runify(world)
	score := GrowPolymerRecursiveRune(worldRunes.template, worldRunes.rules, 10, worldRunes.alphabetSize)

	assert.Equal(t, 1588, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := GrowPolymerIter(world.template, world.rules, 10)
	score := PolymerScore(polymer)
	assert.Equal(t, 3555, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	// never finishes
	polymer := GrowPolymerIter(world.template, world.rules, 40)
	score := PolymerScore(polymer)
	assert.Equal(t, -1, score)
}

// Benchmark_recursive-10    	      22	  49 740 915 ns/op
func Benchmark_recursive(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world, err := ParseInput(reader)
	assert.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		score := GrowPolymerRecursive(world.template, world.rules, 18)
		//assert.Equal(b, 1961318, score) // for 20
		assert.Equal(b, 480563, score)
	}
}

// Benchmark_recursive_runified-10    	     151	   7 837 770 ns/op
func Benchmark_recursive_runified(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world, err := ParseInput(reader)
	assert.Nil(b, err)

	worldRunes := Runify(world)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		score := GrowPolymerRecursiveRune(worldRunes.template, worldRunes.rules, 18, worldRunes.alphabetSize)
		//assert.Equal(b, 1961318, score) // for 20
		assert.Equal(b, 480563, score)
	}
}

// Benchmark_iter-10    	      57	  19 790 479 ns/op
func Benchmark_iter(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world, err := ParseInput(reader)
	assert.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		polymer := GrowPolymerIter(world.template, world.rules, 18)
		score := PolymerScore(polymer)
		assert.Equal(b, 480563, score)
	}
}
