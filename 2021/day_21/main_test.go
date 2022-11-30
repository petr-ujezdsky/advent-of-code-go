package day_21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
