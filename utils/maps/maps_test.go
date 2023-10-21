package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	source := make(map[string]int)

	source["aaa"] = 5
	source["bbb"] = -10

	copied := Copy(source)

	delete(source, "aaa")

	_, ok := source["aaa"]
	assert.False(t, ok)

	assert.Equal(t, 5, copied["aaa"])
}

func TestIntersection(t *testing.T) {
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	map1["a"] = 1
	map1["b"] = 1
	map1["c"] = 1

	map2["b"] = 1

	intersection := Intersection([]map[string]int{map1, map2})

	assert.Equal(t, 1, len(intersection))
	assert.Equal(t, 1, intersection["b"])
}

func TestIntersection_Single(t *testing.T) {
	map1 := make(map[string]int)

	map1["a"] = 1
	map1["b"] = 1
	map1["c"] = 1

	intersection := Intersection([]map[string]int{map1})

	assert.Equal(t, map1, intersection)
}
