package main

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/tree"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 5, len(world.Modules))
	assert.Equal(t, "broadcaster", world.Modules["broadcaster"].Name)
}

func Test_FlipFlop(t *testing.T) {
	aggregator := &Aggregator{}

	module := &Module{
		Name:          "flip",
		Type:          FlipFlop,
		InputModules:  nil,
		OutputModules: nil,
		State:         collections.NewBitSet64(),
	}

	broadcast := &Module{
		Name:          "broadcast",
		Type:          Broadcast,
		InputModules:  nil,
		OutputModules: []*Module{module},
		State:         collections.NewBitSet64(),
	}

	outputSignal, sent := module.OnSignal(Low, broadcast, aggregator)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)

	outputSignal, sent = module.OnSignal(High, broadcast, aggregator)

	assert.False(t, sent)

	outputSignal, sent = module.OnSignal(Low, broadcast, aggregator)

	assert.True(t, sent)
	assert.Equal(t, Low, outputSignal)

	outputSignal, sent = module.OnSignal(High, broadcast, aggregator)

	assert.False(t, sent)

	outputSignal, sent = module.OnSignal(Low, broadcast, aggregator)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)

	outputSignal, sent = module.OnSignal(High, broadcast, aggregator)

	assert.False(t, sent)
}

func Test_Conjunction(t *testing.T) {
	aggregator := &Aggregator{}

	m1 := &Module{Name: "m1"}
	m2 := &Module{Name: "m2"}

	module := &Module{
		Name:          "conjunction",
		Type:          Conjunction,
		InputModules:  []*Module{m1, m2},
		OutputModules: nil,
		State:         collections.BitSet64{},
	}

	outputSignal, sent := module.OnSignal(High, m1, aggregator)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)

	outputSignal, sent = module.OnSignal(High, m2, aggregator)

	assert.True(t, sent)
	assert.Equal(t, Low, outputSignal)

	outputSignal, sent = module.OnSignal(High, m2, aggregator)

	assert.True(t, sent)
	assert.Equal(t, Low, outputSignal)

	outputSignal, sent = module.OnSignal(Low, m2, aggregator)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 32000000, result)
}

func Test_01_example2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 11687500, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 836127690, result)
}

func Test_02_Print_tree(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	button := world.Button

	str := tree.PrintTree(button, func(node *Module) (string, []*Module) {
		return string(node.Type) + node.Name, node.OutputModules
	})

	fmt.Printf("%s\n", str)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}
