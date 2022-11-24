package day_14

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, []rune("NNCB"), world.template)
	assert.Equal(t, 'B', world.rules[duoHash([]rune("CH"))])
	assert.Equal(t, 'C', world.rules[duoHash([]rune("CN"))])
}

func Test_01_example_single(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	_, polymer := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 1, &strings.Builder{})
	assert.Equal(t, "NCNBCHB", polymer.String())

	_, polymer = GrowPolymerRecursiveRuneCaching(world.template, world.rules, 2, &strings.Builder{})
	assert.Equal(t, "NBCCNBBBCBHCB", polymer.String())

	_, polymer = GrowPolymerRecursiveRuneCaching(world.template, world.rules, 3, &strings.Builder{})
	assert.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", polymer.String())

	_, polymer = GrowPolymerRecursiveRuneCaching(world.template, world.rules, 4, &strings.Builder{})
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer.String())
}

func Test_01_example_multi(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	_, polymer := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 4, &strings.Builder{})
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer.String())
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score, _ := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 10, nil)

	assert.Equal(t, 1588, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score, _ := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 10, nil)

	assert.Equal(t, 3555, score)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score, _ := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 40, nil)
	assert.Equal(t, 2188189693529, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	score, _ := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 40, nil)

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
		score, _ := GrowPolymerRecursiveRuneCaching(world.template, world.rules, 18, nil)
		assert.Equal(b, 480563, score)
	}
}
