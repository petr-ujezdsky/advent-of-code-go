package day_21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dieRollOutcomes(t *testing.T) {
	assert.Equal(t, 7, len(dieRollOutcomes))
	assert.Equal(t, 1, dieRollOutcomes[3])
	assert.Equal(t, 3, dieRollOutcomes[4])
	assert.Equal(t, 6, dieRollOutcomes[5])
	// there are 7 combinations to get sum of 6
	assert.Equal(t, 7, dieRollOutcomes[6])
	assert.Equal(t, 6, dieRollOutcomes[7])
	assert.Equal(t, 3, dieRollOutcomes[8])
	assert.Equal(t, 1, dieRollOutcomes[9])
}

func Test_01_example(t *testing.T) {
	p1 := Player{
		Position: 4,
		Score:    0,
	}

	p2 := Player{
		Position: 8,
		Score:    0,
	}

	result, dieRolls, p1, p2 := Play(p1, p2, 1000)

	assert.Equal(t, 1000, p1.Score)
	assert.Equal(t, 745, p2.Score)
	assert.Equal(t, 993, dieRolls)
	assert.Equal(t, 739785, result)
}

func Test_01(t *testing.T) {
	p1 := Player{
		Position: 8,
		Score:    0,
	}

	p2 := Player{
		Position: 10,
		Score:    0,
	}

	result, dieRolls, p1, p2 := Play(p1, p2, 1000)

	assert.Equal(t, 1000, p1.Score)
	assert.Equal(t, 810, p2.Score)
	assert.Equal(t, 747, dieRolls)
	assert.Equal(t, 605070, result)
}

func Test_02_example(t *testing.T) {
	p1 := Player{
		Position: 4,
		Score:    0,
	}

	p2 := Player{
		Position: 8,
		Score:    0,
	}

	p1wins, p2wins, max := PlayFaster(p1, p2)

	assert.Equal(t, 444356092776315, p1wins)
	assert.Equal(t, 341960390180808, p2wins)
	assert.Equal(t, 444356092776315, max)
}

func Test_02(t *testing.T) {
	p1 := Player{
		Position: 8,
		Score:    0,
	}

	p2 := Player{
		Position: 10,
		Score:    0,
	}

	p1wins, p2wins, max := PlayFaster(p1, p2)

	assert.Equal(t, 218433063958910, p1wins)
	assert.Equal(t, 189371397363999, p2wins)
	assert.Equal(t, 218433063958910, max)
}

// Benchmark_fast-10    	     200	   5 966 090 ns/op
func Benchmark_fast(b *testing.B) {
	p1 := Player{
		Position: 4,
		Score:    0,
	}

	p2 := Player{
		Position: 8,
		Score:    0,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p1wins, p2wins, max := PlayFaster(p1, p2)

		assert.Equal(b, 444356092776315, p1wins)
		assert.Equal(b, 341960390180808, p2wins)
		assert.Equal(b, 444356092776315, max)
	}
}
