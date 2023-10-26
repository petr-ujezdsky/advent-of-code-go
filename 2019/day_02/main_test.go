package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_01_example(t *testing.T) {
	assert.Equal(t, 2, DoWithInputPart01(ParseInput("1,0,0,0,99")))
	assert.Equal(t, 2, DoWithInputPart01(ParseInput("2,3,0,3,99")))
	assert.Equal(t, 2, DoWithInputPart01(ParseInput("2,4,4,5,99,0")))
	assert.Equal(t, 30, DoWithInputPart01(ParseInput("1,1,1,4,99,5,6,0,99")))

	assert.Equal(t, 3500, DoWithInputPart01(ParseInput("1,9,10,3,2,3,11,0,99,30,40,50")))
}

func Test_01(t *testing.T) {
	program := ParseInput("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,13,19,23,2,23,9,27,1,6,27,31,2,10,31,35,1,6,35,39,2,9,39,43,1,5,43,47,2,47,13,51,2,51,10,55,1,55,5,59,1,59,9,63,1,63,9,67,2,6,67,71,1,5,71,75,1,75,6,79,1,6,79,83,1,83,9,87,2,87,10,91,2,91,10,95,1,95,5,99,1,99,13,103,2,103,9,107,1,6,107,111,1,111,5,115,1,115,2,119,1,5,119,0,99,2,0,14,0")
	PatchProgram(program)

	result := DoWithInputPart01(program)
	assert.Equal(t, 2894520, result)
}

func Test_02(t *testing.T) {
	program := ParseInput("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,13,19,23,2,23,9,27,1,6,27,31,2,10,31,35,1,6,35,39,2,9,39,43,1,5,43,47,2,47,13,51,2,51,10,55,1,55,5,59,1,59,9,63,1,63,9,67,2,6,67,71,1,5,71,75,1,75,6,79,1,6,79,83,1,83,9,87,2,87,10,91,2,91,10,95,1,95,5,99,1,99,13,103,2,103,9,107,1,6,107,111,1,111,5,115,1,115,2,119,1,5,119,0,99,2,0,14,0")

	result := DoWithInputPart02(program)
	assert.Equal(t, 9342, result)
}
