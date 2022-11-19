package day_07

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	entries, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, entries[0].Digits)
	assert.Equal(t, []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}, entries[0].Outputs)

	assert.Equal(t, []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, entries[9].Digits)
	assert.Equal(t, []string{"fgae", "cfgab", "fg", "bagce"}, entries[9].Outputs)
}

func Test_01_example_easy_outputs_count(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	entries, err := ParseInput(reader)
	assert.Nil(t, err)

	count := CountEasyOutputs(entries)

	assert.Equal(t, 26, count)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	entries, err := ParseInput(reader)
	assert.Nil(t, err)

	count := CountEasyOutputs(entries)

	assert.Equal(t, 470, count)
}

func Test_02_example_mini(t *testing.T) {
	entry := Entry{
		Digits:  []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		Outputs: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
	}

	output, ok := TryDecodeOutput([]rune("deafgbc"), entry)
	assert.True(t, ok)
	assert.Equal(t, 5353, output)
}

func Test_02_example_one(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	entries, err := ParseInput(reader)
	assert.Nil(t, err)

	output, decoder, iterations, ok := BruteForceDecode(entries[0])
	fmt.Println("Decoder:", string(decoder), "iterations:", iterations)
	assert.True(t, ok)
	assert.Equal(t, 8394, output)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	entries, err := ParseInput(reader)
	assert.Nil(t, err)

	sum, ok := DecodeAndSum(entries)
	assert.True(t, ok)
	assert.Equal(t, 61229, sum)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	entries, err := ParseInput(reader)
	assert.Nil(t, err)

	sum, ok := DecodeAndSum(entries)
	assert.True(t, ok)
	assert.Equal(t, 989396, sum)
}
