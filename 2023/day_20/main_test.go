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

	assert.Equal(t, 6, len(world.Modules))
	assert.Equal(t, "broadcaster", world.Modules["broadcaster"].Name)
}

func Test_FlipFlop(t *testing.T) {
	aggregator := &Aggregator{}
	state := collections.NewBitSet128()

	module := &Module{
		Name:             "flip",
		Type:             FlipFlop,
		InputModules:     nil,
		OutputModules:    nil,
		InputsAggregator: &Aggregator{},
	}

	broadcast := &Module{
		Name:             "broadcast",
		Type:             Broadcast,
		InputModules:     nil,
		OutputModules:    []*Module{module},
		InputsAggregator: &Aggregator{},
	}

	outputSignal, sent := module.OnSignal(Low, broadcast, aggregator, &state)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)

	outputSignal, sent = module.OnSignal(High, broadcast, aggregator, &state)

	assert.False(t, sent)

	outputSignal, sent = module.OnSignal(Low, broadcast, aggregator, &state)

	assert.True(t, sent)
	assert.Equal(t, Low, outputSignal)

	outputSignal, sent = module.OnSignal(High, broadcast, aggregator, &state)

	assert.False(t, sent)

	outputSignal, sent = module.OnSignal(Low, broadcast, aggregator, &state)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)

	outputSignal, sent = module.OnSignal(High, broadcast, aggregator, &state)

	assert.False(t, sent)
}

func Test_Conjunction(t *testing.T) {
	aggregator := &Aggregator{}
	state := collections.NewBitSet128()

	m1 := &Module{Name: "m1"}
	m2 := &Module{Name: "m2"}

	module := &Module{
		Name:             "conjunction",
		Type:             Conjunction,
		InputModules:     []*Module{m1, m2},
		OutputModules:    nil,
		InputsAggregator: &Aggregator{},
	}

	outputSignal, sent := module.OnSignal(High, m1, aggregator, &state)

	assert.True(t, sent)
	assert.Equal(t, High, outputSignal)

	outputSignal, sent = module.OnSignal(High, m2, aggregator, &state)

	assert.True(t, sent)
	assert.Equal(t, Low, outputSignal)

	outputSignal, sent = module.OnSignal(High, m2, aggregator, &state)

	assert.True(t, sent)
	assert.Equal(t, Low, outputSignal)

	outputSignal, sent = module.OnSignal(Low, m2, aggregator, &state)

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

func Test_02_Print_nodes(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	for _, module := range world.Modules {
		var in, out []string

		for _, m := range module.InputModules {
			in = append(in, m.Name)
		}

		for _, m := range module.OutputModules {
			out = append(out, m.Name)
		}

		fmt.Printf("%v%v - IN: %v, OUT: %v\n", string(module.Type), module.Name, in, out)
	}
}

func Test_02_Print_reverse_tree(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	rx := world.Modules["rx"]

	str := tree.PrintTree(rx, func(node *Module) (string, []*Module) {
		return string(node.Type) + node.Name, node.InputModules
	})

	fmt.Printf("%s\n", str)
}

func Test_02_State_size(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	stateSize := 0

	for _, module := range world.Modules {
		switch module.Type {
		case FlipFlop:
			stateSize++
		case Conjunction:
			stateSize += len(module.InputModules)
		}
	}
	fmt.Printf("Need state of size %d\n", stateSize)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Benchmark_1000Steps(b *testing.B) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(b, err)

	world := ParseInput(reader)

	for i := 0; i < b.N; i++ {
		result := DoWithInputPart01(world)
		assert.Equal(b, 836127690, result)
	}
}
