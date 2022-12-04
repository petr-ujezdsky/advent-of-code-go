package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterval_Contains(t *testing.T) {
	assert.Equal(t, false, NewInterval(10, 20).Contains(5))
	assert.Equal(t, true, NewInterval(10, 20).Contains(10))
	assert.Equal(t, true, NewInterval(10, 20).Contains(15))
	assert.Equal(t, true, NewInterval(10, 20).Contains(20))
	assert.Equal(t, false, NewInterval(10, 20).Contains(25))
}

func TestInterval_Subtract(t *testing.T) {
	assert.Equal(t, []IntervalI{NewInterval(10, 14)}, NewInterval(10, 20).Subtract(NewInterval(15, 25)))
	assert.Equal(t, []IntervalI{NewInterval(16, 20)}, NewInterval(10, 20).Subtract(NewInterval(5, 15)))
	assert.Equal(t, []IntervalI{NewInterval(10, 11), NewInterval(18, 20)}, NewInterval(10, 20).Subtract(NewInterval(12, 17)))
	assert.Equal(t, []IntervalI(nil), NewInterval(10, 20).Subtract(NewInterval(5, 25)))
	assert.Equal(t, []IntervalI{NewInterval(10, 20)}, NewInterval(10, 20).Subtract(NewInterval(5, 7)))

	assert.Equal(t, []IntervalI{NewInterval(11, 20)}, NewInterval(10, 20).Subtract(NewInterval(5, 10)))
	assert.Equal(t, []IntervalI{NewInterval(10, 19)}, NewInterval(10, 20).Subtract(NewInterval(20, 25)))

	assert.Equal(t, []IntervalI{NewInterval(10, 14), NewInterval(16, 20)}, NewInterval(10, 20).Subtract(NewInterval(15, 15)))
}
