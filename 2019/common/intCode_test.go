package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun_relativeBase(t *testing.T) {
	program := []int{109, 2000, 109, 19, 204, -2018, 99}
	outputs := RunProgram(nil, program)

	assert.Equal(t, []int{2000}, outputs)
}

func TestRun_bigMemory1(t *testing.T) {
	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	outputs := RunProgram(nil, program)

	assert.Equal(t, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, outputs)
}

func TestRun_bigMemory2(t *testing.T) {
	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	outputs := RunProgram(nil, program)

	assert.Equal(t, []int{1219070632396864}, outputs)
}

func TestRun_bigMemory3(t *testing.T) {
	program := []int{104, 1125899906842624, 99}
	outputs := RunProgram(nil, program)

	assert.Equal(t, []int{1125899906842624}, outputs)
}
