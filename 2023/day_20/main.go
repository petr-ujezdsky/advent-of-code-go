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

func (s SignalType) String() string {
	if s {
		return "high"
	}

	return "low"
}

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
	InputsAggregator            *Aggregator
}

func (m *Module) OnSignal(signal SignalType, from *Module, aggregator *Aggregator) (SignalType, bool) {
	// aggregate signal locally
	m.InputsAggregator.aggregate(signal, 1)

	outputSignal, ok := m.checkSignal(signal, from)

	if ok {
		m.sendSignal(outputSignal, aggregator)
	}

	return outputSignal, ok
}

func (m *Module) checkSignal(signal SignalType, from *Module) (SignalType, bool) {
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
		for i := range m.InputModules {
			if !m.State.Contains(i) {
				allHigh = false
				break
			}
		}

		outputSignal := High
		if allHigh {
			outputSignal = Low
		}

		return outputSignal, true
	case Broadcast:
		return signal, true
	}

	//fmt.Printf("Unknown type, %s obtained %v\n", m.Name, signal)
	return signal, false
}

func (m *Module) sendSignal(signal SignalType, aggregator *Aggregator) {
	// aggregate counts
	aggregator.aggregate(signal, len(m.OutputModules))

	// send signal - process last module first
	for i := len(m.OutputModules) - 1; i >= 0; i-- {
		output := m.OutputModules[i]
		//fmt.Printf("%s -%v-> %s\n", m.Name, signal, output.Name)
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

func (a *Aggregator) aggregate(signal SignalType, count int) {
	switch signal {
	case Low:
		a.LowCount += count
	case High:
		a.HighCount += count
	}
}

func (a *Aggregator) reset() {
	a.LowCount = 0
	a.HighCount = 0
}

type Modules = map[string]*Module

type World struct {
	Modules Modules
	Button  *Module
}

func DoWithInputPart01(world World) int {
	button := world.Button
	aggregator := &Aggregator{}

	pushCount := 1000

	for i := 0; i < pushCount; i++ {
		button.OnSignal(Low, nil, aggregator)
	}

	fmt.Printf("Counts %v\n", *aggregator)

	return aggregator.LowCount * aggregator.HighCount
}

func DoWithInputPart02(world World) int {
	button := world.Button
	rxModule := world.Modules["rx"]
	aggregator := &Aggregator{}
	targetRxCounts := Aggregator{LowCount: 1, HighCount: 0}

	pushCount := 1

	for {
		button.OnSignal(Low, nil, aggregator)

		if *rxModule.InputsAggregator == targetRxCounts {
			return pushCount
		}

		//fmt.Printf("#%15d rx: %v\n", pushCount, *rxModule.InputsAggregator)

		rxModule.InputsAggregator.reset()
		pushCount++
	}
}

func getOrCreateModule(name string, modules Modules) *Module {
	return maps.GetOrCompute(modules, name, func(key string) *Module {
		return &Module{
			Name:             name,
			Type:             '?',
			State:            collections.NewEmptyBitSet64(),
			InputsAggregator: &Aggregator{},
		}
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
	}

	return World{
		Modules: modules,
		Button:  modules["button"],
	}
}
