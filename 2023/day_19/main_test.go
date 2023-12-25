package main

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//example tree
//in
//├──px
//│  ├──qkq
//│  │  ├──A
//│  │  └──crn
//│  │     ├──A
//│  │     └──R
//│  ├──A
//│  └──rfg
//│     ├──gd
//│     │  ├──R
//│     │  └──R
//│     ├──R
//│     └──A
//└──qqz
//   ├──qs
//   │  ├──A
//   │  └──lnx
//   │     ├──A
//   │     └──A
//   ├──hdj
//   │  ├──A
//   │  └──pv
//   │     ├──R
//   │     └──A
//   └──R

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 11+2, len(world.Workflows))
	assert.Equal(t, 5, len(world.Parts))

	workflow := world.Start
	assert.Equal(t, "in", workflow.Name)
	assert.Equal(t, "qqz", workflow.Conditions[len(workflow.Conditions)-1].Next.Name)
	assert.Equal(t, TypeNormal, workflow.Type)

	condition := workflow.Conditions[0]
	assert.Equal(t, CategoryS, condition.Category)
	assert.Equal(t, '<', condition.Operand)
	assert.Equal(t, 1351, condition.Amount)
	assert.Equal(t, "px", condition.Next.Name)

	workflow = world.Workflows["crn"]
	assert.Equal(t, TypeRejects, workflow.Conditions[len(workflow.Conditions)-1].Next.Type)
	assert.Equal(t, workflow, workflow.Conditions[len(workflow.Conditions)-1].Owner)

	workflow = world.Workflows["pv"]
	assert.Equal(t, TypeAccepts, workflow.Conditions[len(workflow.Conditions)-1].Next.Type)
}

func Test_01_reverse_links(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	accepting := world.Workflows["A"]

	// find A in path in -> qqz -> qs -> A
	condition := accepting.ParentConditions[5]
	assert.Equal(t, "qs", condition.Owner.Name)
	assert.Nil(t, condition.Previous)

	condition = condition.Owner.ParentCondition
	assert.Equal(t, "qqz", condition.Owner.Name)
	assert.Nil(t, condition.Previous)

	condition = condition.Owner.ParentCondition
	assert.Equal(t, "in", condition.Owner.Name)
	assert.Equal(t, "px", condition.Previous.Next.Name)
	assert.Nil(t, condition.Previous.Previous)
	assert.Nil(t, condition.Owner.ParentCondition)
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
