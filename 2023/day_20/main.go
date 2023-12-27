package main

import (
	"bufio"
	_ "embed"
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

func (m *Module) OnSignal(signal SignalType, from *Module) (SignalType, bool) {
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

			m.sendSignal(outputSignal)

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

		m.sendSignal(outputSignal)

		return outputSignal, true
	}

	panic("Not implemented")
}

func (m *Module) sendSignal(signal SignalType) {
	for _, output := range m.OutputModules {
		output.OnSignal(signal, m)
	}
}

type ModuleType = rune

const (
	Broadcast   ModuleType = 'b'
	FlipFlop               = '%'
	Conjunction            = '&'
)

type Modules = map[string]*Module

type World struct {
	Modules Modules
}

func DoWithInputPart01(world World) int {
	return 0
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

		// exception for 'broadcast'
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

	return World{Modules: modules}
}
