package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun_relativeBase(t *testing.T) {
	program := []int{109, 2000, 109, 19, 204, -2018, 99}
	outputs := RunProgram2(nil, program)

	assert.Equal(t, []int{2000}, outputs)
}
