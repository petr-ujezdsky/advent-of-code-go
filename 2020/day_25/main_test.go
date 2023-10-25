package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_01_example(t *testing.T) {
	result := DoWithInputPart01(5764801, 17807724)
	assert.Equal(t, 14897079, result)
}

func Test_01(t *testing.T) {
	result := DoWithInputPart01(9789649, 3647239)
	assert.Equal(t, 8740494, result)
}
