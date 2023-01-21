package day_15

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

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

	levels := ParseInput(reader)

	score := FindPathScore(levels)
	assert.Equal(t, 40, score)
}

func Test_01_example_back_propagation(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

	bestScore, bestScores := CalcBestScore(levels)

	expected := `
40 40 34 31 30 25 24 25 28 30
39 37 31 30 27 22 21 22 21 28
37 36 33 27 22 21 20 19 19 21
39 36 30 26 22 19 19 14 13 14
33 32 26 23 19 18 14 13 12 13
32 29 28 19 18 16 13 12  9 10
31 28 27 25 16 15 13  9  7  9
28 27 25 20 16 14 13  7  4  2
30 28 23 20 19 16  9  4  2  1
28 25 24 23 20 18 14  9  1  0`
	assert.Equal(t, expected, "\n"+bestScores.StringFmt(matrix.FmtFmt[int]("%2d")))
	//fmt.Println(bestScores.StringFmt(utils.FmtFmt[int]("%2d")))
	assert.Equal(t, 40, bestScore)
}

func Test_01_example_a_star(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

	path, _, score, ok := CalcBestScoreAStar(levels)
	assert.True(t, ok)
	assert.NotNil(t, path)
	assert.Equal(t, 40, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

	score, bestScores := CalcBestScore(levels)
	assert.Equal(t, 462, score)

	//fmt.Println(bestScores.StringFmt(utils.FmtFmt[int]("%3d")))
	assert.NotNil(t, bestScores)
}

func Test_01_a_star(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

	_, _, score, ok := CalcBestScoreAStar(levels)
	assert.True(t, ok)
	assert.Equal(t, 462, score)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

	enlargedLevels := EnlargeWorld(levels)
	//fmt.Println(enlargedLevels.StringFmt(utils.FmtFmt[int]("%2d")))

	_, _, score, ok := CalcBestScoreAStar(enlargedLevels)
	assert.True(t, ok)
	assert.Equal(t, 315, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	levels := ParseInput(reader)

	enlargedLevels := EnlargeWorld(levels)

	_, _, score, ok := CalcBestScoreAStar(enlargedLevels)
	assert.True(t, ok)
	assert.Equal(t, 2846, score)
}

// Benchmark_back_propagation-10    	    5437	    217 993 ns/op
func Benchmark_back_propagation(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	levels := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bestScore, _ := CalcBestScore(levels)
		assert.Equal(b, 40, bestScore)
	}
}

// Benchmark_a_star-10    	   13176	     89 633 ns/op
func Benchmark_a_star(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	levels := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		path, _, score, ok := CalcBestScoreAStar(levels)
		assert.True(b, ok)
		assert.NotNil(b, path)
		assert.Equal(b, 40, score)
	}
}

// Benchmark_back_propagation_big-10    	       1	7 331 581 583 ns/op
func Benchmark_back_propagation_big(b *testing.B) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(b, err)

	levels := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bestScore, _ := CalcBestScore(levels)
		assert.Equal(b, 462, bestScore)
	}
}

// Benchmark_a_star_big-10    	      16	  69 845 609 ns/op
func Benchmark_a_star_big(b *testing.B) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(b, err)

	levels := ParseInput(reader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		path, _, score, ok := CalcBestScoreAStar(levels)
		assert.True(b, ok)
		assert.NotNil(b, path)
		assert.Equal(b, 462, score)
	}
}
