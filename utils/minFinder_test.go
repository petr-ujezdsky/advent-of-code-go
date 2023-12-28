package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinFinder(t *testing.T) {
	minFinder := NewMinFinder[string]()

	minFinder.Inspect(5, "five")
	minFinder.Inspect(10, "ten")
	minFinder.Inspect(1, "one")
	minFinder.Inspect(2, "two")

	minimum, item := minFinder.Result()
	assert.Equal(t, 1, minimum)
	assert.Equal(t, "one", item)
}
