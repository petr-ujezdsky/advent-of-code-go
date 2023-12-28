package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxFinder(t *testing.T) {
	maxFinder := NewMaxFinder[string]()

	maxFinder.Inspect(5, "five")
	maxFinder.Inspect(10, "ten")
	maxFinder.Inspect(1, "one")
	maxFinder.Inspect(2, "two")

	maximum, item := maxFinder.Result()
	assert.Equal(t, 10, maximum)
	assert.Equal(t, "ten", item)
	assert.Equal(t, "Max 10 for ten", maxFinder.String())
}
