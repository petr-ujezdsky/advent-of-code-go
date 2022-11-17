package day_05

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	lines, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, 10, len(lines))

	line := lines[0]
	assert.Equal(t, NewLine(0, 9, 5, 9), line)

	line = lines[9]
	assert.Equal(t, NewLine(5, 5, 8, 2), line)
}

//func Test_01_example_play(t *testing.T) {
//	reader, err := os.Open("data-00-example.txt")
//	assert.Nil(t, err)
//
//	drawn, bingos, err := ParseInput(reader)
//	assert.Nil(t, err)
//
//	bingo, score := Play(bingos, drawn)
//	assert.NotNil(t, bingo)
//	assert.Equal(t, 4512, score)
//}
//
//func Test_01(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	drawn, bingos, err := ParseInput(reader)
//	assert.Nil(t, err)
//
//	bingo, score := Play(bingos, drawn)
//	assert.NotNil(t, bingo)
//	assert.Equal(t, 6592, score)
//}
//
//func Test_02_example(t *testing.T) {
//	reader, err := os.Open("data-00-example.txt")
//	assert.Nil(t, err)
//
//	drawn, bingos, err := ParseInput(reader)
//	assert.Nil(t, err)
//
//	bingo, score := PlayTillEnd(bingos, drawn)
//	assert.NotNil(t, bingo)
//	assert.Equal(t, 1924, score)
//}
//
//func Test_02(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	drawn, bingos, err := ParseInput(reader)
//	assert.Nil(t, err)
//
//	bingo, score := PlayTillEnd(bingos, drawn)
//	assert.NotNil(t, bingo)
//	assert.Equal(t, 31755, score)
//}
