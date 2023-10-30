package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_01(t *testing.T) {
	world := World{
		From: 231832,
		To:   767346,
	}

	result := DoWithInputPart01(world)
	assert.Equal(t, 1330, result)
}

func Test_02_example(t *testing.T) {
	world := World{
		From: 231832,
		To:   767346,
	}

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	world := World{
		From: 231832,
		To:   767346,
	}

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}
