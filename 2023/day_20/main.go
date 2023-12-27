package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strings"
)

type SignalType bool

const (
	Low  SignalType = false
	High SignalType = true
)

//type StateFlipFlop struct {
//	On bool
//}
//
//type StateConjunction struct {
//
//}

type Module struct {
	Name                        string
	Type                        ModuleType
	InputModules, OutputModules []*Module
	State                       collections.BitSet64
}

func (m *Module) OnSignal(signal SignalType, from *Module, aggregator *Aggregator) (SignalType, bool) {
	switch m.Type {
	// FlipFlop
	case FlipFlop:
		// contains ~ ON
		if signal == Low {
			m.State.Invert(0)

			outputSignal := Low
			if m.State.Contains(0) {
				outputSignal = High
			}

			m.sendSignal(outputSignal, aggregator)

			return outputSignal, true
		}

		return Low, false
	// Conjunction
	case Conjunction:
		// TODO pujezdsky optimize
		// find input module index
		index := -1
		for i, inputModule := range m.InputModules {
			if inputModule == from {
				index = i
				break
			}
		}

		// store current signal per given input module
		switch signal {
		case Low:
			m.State.Remove(index)
		case High:
			m.State.Push(index)
		}

		// TODO pujezdsky optimize
		// check if all are HIGH
		allHigh := true
		for i, _ := range m.InputModules {
			if !m.State.Contains(i) {
				allHigh = false
				break
			}
		}

		outputSignal := High
		if allHigh {
			outputSignal = Low
		}

		m.sendSignal(outputSignal, aggregator)

		return outputSignal, true
	case Broadcast:
		m.sendSignal(signal, aggregator)

		return signal, true
	}

	panic("Not implemented")
}

func (m *Module) sendSignal(signal SignalType, aggregator *Aggregator) {
	// aggregate counts
	switch signal {
	case Low:
		aggregator.LowCount += len(m.OutputModules)
	case High:
		aggregator.HighCount += len(m.OutputModules)
	}

	// send signal
	for _, output := range m.OutputModules {
		output.OnSignal(signal, m, aggregator)
	}
}

type ModuleType = rune

const (
	Broadcast   ModuleType = 'b'
	FlipFlop               = '%'
	Conjunction            = '&'
)

type Aggregator struct {
	LowCount, HighCount int
}

type Modules = map[string]*Module

type World struct {
	Modules   Modules
	Broadcast *Module
}

func DoWithInputPart01(world World) int {
	broadcast := world.Broadcast
	aggregator := &Aggregator{}

	pushCount := 1000

	for i := 0; i < pushCount; i++ {
		broadcast.OnSignal(Low, nil, aggregator)
	}

	fmt.Printf("Counts %v", *aggregator)

	return aggregator.LowCount * aggregator.HighCount
}

func DoWithInputPart02(world World) int {
	return 0
}

func getOrCreateModule(name string, modules Modules) *Module {
	return maps.GetOrCompute(modules, name, func(key string) *Module {
		return &Module{Name: name}
	})
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	modules := make(Modules)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		mtype := parts[0][0]
		name := parts[0][1:]

		// exception for 'broadcaster'
		if mtype == 'b' {
			name = parts[0]
		}

		module := getOrCreateModule(name, modules)

		outputsRaw := strings.Split(parts[1], ", ")
		outputs := slices.Map(outputsRaw, func(moduleName string) *Module {
			output := getOrCreateModule(moduleName, modules)
			output.InputModules = append(output.InputModules, module)
			return output
		})

		module.Type = rune(mtype)
		module.OutputModules = outputs
		module.State = collections.NewEmptyBitSet64()
	}

	return World{
		Modules:   modules,
		Broadcast: modules["broadcaster"],
	}
}
