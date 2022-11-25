package day_15

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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

func Test_01_example_back_propagation(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	//reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	levels, err := ParseInput(reader)
	assert.Nil(t, err)

	bestScore, bestScores := CalcBestScore(levels)

	expected := `
40 40 34 31 30 25 24 28 28 39
39 37 31 30 27 22 21 22 21 37
37 36 33 27 22 21 20 19 19 29
40 36 30 26 22 19 19 14 13 20
33 32 26 23 19 18 14 13 12 19
32 29 28 19 18 16 13 12  9 12
31 28 27 25 16 15 13  9  7 11
28 27 25 20 16 14 13  7  4  2
35 33 24 21 20 17  9  4  2  1
36 33 32 31 22 18 14  9  1  0`
	assert.Equal(t, expected, "\n"+bestScores.StringFmt(utils.FmtFmt[int]("%2d")))
	assert.Equal(t, 40, bestScore)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	levels, err := ParseInput(reader)
	assert.Nil(t, err)

	score, bestScores := CalcBestScore(levels)
	assert.Equal(t, 462, score)

	//fmt.Println(bestScores.StringFmt(utils.FmtFmt[int]("%3d")))
	assert.NotNil(t, bestScores)
}
