package main

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 11+2, len(world.Workflows))
	assert.Equal(t, 5, len(world.Parts))

	start := world.Start
	assert.Equal(t, "in", start.Name)
	assert.Equal(t, "qqz", start.Fallback.Name)
	assert.Equal(t, TypeNormal, start.Type)

	condition := start.Conditions[0]
	assert.Equal(t, CategoryS, condition.Category)
	assert.Equal(t, '<', condition.Operand)
	assert.Equal(t, 1351, condition.Amount)
	assert.Equal(t, "px", condition.Next.Name)

	assert.Equal(t, TypeRejects, world.Workflows["crn"].Fallback.Type)
	assert.Equal(t, TypeAccepts, world.Workflows["pv"].Fallback.Type)
}

func Test_01_resolve(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	start := world.Start

	part := world.Parts[0]
	path := make([]*Workflow, 0)
	accepted, path := start.Resolve(part, path, true)

	pathNames := slices.Map(path, func(w *Workflow) string { return w.Name })

	assert.True(t, accepted)
	assert.Equal(t, "[in qqz qs lnx A]", fmt.Sprintf("%v", pathNames))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 19114, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 348378, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}