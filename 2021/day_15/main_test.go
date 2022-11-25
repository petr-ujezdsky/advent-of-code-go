package day_15

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	levels, err := ParseInput(reader)
	assert.Nil(t, err)

	expected := `
1 1 6 3 7 5 1 7 4 2
1 3 8 1 3 7 3 6 7 2
2 1 3 6 5 1 1 3 2 8
3 6 9 4 9 3 1 5 6 9
7 4 6 3 4 1 7 1 1 1
1 3 1 9 1 2 8 1 3 7
1 3 5 9 9 1 2 4 2 1
3 1 2 5 4 2 1 6 3 9
1 2 9 3 1 3 8 5 2 1
2 3 1 1 9 4 4 5 8 1`

	assert.Equal(t, expected, "\n"+levels.String())
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	levels, err := ParseInput(reader)
	assert.Nil(t, err)

	score := FindPathScore(levels)
	assert.Equal(t, 40, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	levels, err := ParseInput(reader)
	assert.Nil(t, err)

	score := FindPathScore(levels)
	assert.Equal(t, 40, score)
}
