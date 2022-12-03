package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rucksack := ParseInput(reader)

	assert.Equal(t, 6, len(rucksack))
	assert.Equal(t, []rune("CrZsJsPPZsGz"), rucksack[5].Left)
	assert.Equal(t, []rune("wwsLwLmpwMDw"), rucksack[5].Right)
}

func Test_01_commonItem(t *testing.T) {
	ci := commonItem(NewRucksack("vJrwpWtwJgWrhcsFMMfFFhFp"))
	assert.Equal(t, 'p', ci)
}

func Test_01_commonChars(t *testing.T) {
	ch := commonChar([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg"})

	assert.Equal(t, 'r', ch)
}

func Test_01_itemScore(t *testing.T) {
	assert.Equal(t, 1, itemScore('a'))
	assert.Equal(t, 26, itemScore('z'))
	assert.Equal(t, 27, itemScore('A'))
	assert.Equal(t, 52, itemScore('Z'))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rucksacks := ParseInput(reader)

	score := Score(rucksacks)
	assert.Equal(t, 157, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rucksacks := ParseInput(reader)

	score := Score(rucksacks)
	assert.Equal(t, 8139, score)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rucksacks := ParseInput(reader)

	score := GroupsScore(rucksacks)
	assert.Equal(t, 70, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rucksacks := ParseInput(reader)

	score := GroupsScore(rucksacks)
	assert.Equal(t, 70, score)
}
