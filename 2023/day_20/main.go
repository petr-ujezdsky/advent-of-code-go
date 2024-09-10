package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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

type Signal struct {
	Type     SignalType
	From, To *Module
}

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
	StateIndex                  int
	InputsAggregator            *Aggregator
}

func (m *Module) OnSignal(signal SignalType, from *Module, aggregator *Aggregator, state *collections.BitSet128) (SignalType, bool) {
	// aggregate signal locally
	m.InputsAggregator.aggregate(signal, 1)

	outputSignal, ok := m.checkSignal(signal, from, state)

	if ok {
		m.sendSignal(outputSignal, aggregator, state)
	}

	return outputSignal, ok
}

func (m *Module) OnSignal2(signal SignalType, from *Module, state *collections.BitSet128, signals *collections.Queue[Signal]) (SignalType, bool) {
	// aggregate signal locally
	m.InputsAggregator.aggregate(signal, 1)

	outputSignal, ok := m.checkSignal(signal, from, state)

	if ok {
		for _, outputModule := range m.OutputModules {
			signalNew := Signal{
				Type: outputSignal,
				From: m,
				To:   outputModule,
			}

			signals.Push(signalNew)
		}
	}

	return outputSignal, ok
}

func (m *Module) checkSignal(signal SignalType, from *Module, state *collections.BitSet128) (SignalType, bool) {
	switch m.Type {
	// FlipFlop
	case FlipFlop:
		// contains ~ ON
		if signal == Low {
			state.Invert(m.StateIndex + 0)

			outputSignal := Low
			if state.Contains(m.StateIndex + 0) {
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

		if index == -1 {
			panic("Module not found")
		}

		// store current signal per given input module
		switch signal {
		case Low:
			state.Remove(m.StateIndex + index)
			return High, true
		case High:
			state.Push(m.StateIndex + index)
		}

		// TODO pujezdsky optimize
		// check if all are HIGH
		for i := range m.InputModules {
			if !state.Contains(m.StateIndex + i) {
				// found low, send high
				return High, true
			}
		}

		// all are high -> send low
		return Low, true
	case Broadcast:
		return signal, true
	case Terminal:
		return signal, false
	default:
		panic(fmt.Sprintf("Unknown type %s, %s obtained %v\n", string(m.Type), m.Name, signal))
	}
}

func (m *Module) sendSignal(signal SignalType, aggregator *Aggregator, state *collections.BitSet128) {
	// aggregate counts
	aggregator.aggregate(signal, len(m.OutputModules))

	// send signal - process last module first
	for i := len(m.OutputModules) - 1; i >= 0; i-- {
		output := m.OutputModules[i]
		//fmt.Printf("%s -%v-> %s\n", m.Name, signal, output.Name)
		output.OnSignal(signal, m, aggregator, state)
	}
}

type ModuleType = rune

const (
	Broadcast   ModuleType = 'b'
	FlipFlop               = '%'
	Conjunction            = '&'
	Terminal               = '?'
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
	Modules             Modules
	Button, Broadcaster *Module
}

func DoWithInputPart01(world World) int {
	button := world.Button
	broadcaster := world.Broadcaster
	aggregator := &Aggregator{}

	pushCount := 1000
	state := collections.NewBitSet128()
	buttonPressSignal := Signal{
		Type: Low,
		From: button,
		To:   broadcaster,
	}

	signals := collections.NewQueue[Signal]()

	for i := 0; i < pushCount; i++ {
		//button.OnSignal(Low, nil, aggregator, &state)
		//fmt.Printf("Pushing button #%d\n", i+1)

		signals.Push(buttonPressSignal)
		processSignals(&signals, aggregator, &state)
	}

	//fmt.Printf("Counts %v\n", *aggregator)

	return aggregator.LowCount * aggregator.HighCount
}

func processSignals(signals *collections.Queue[Signal], aggregator *Aggregator, state *collections.BitSet128) {
	for !signals.Empty() {
		signal := signals.Pop()
		aggregator.aggregate(signal.Type, 1)
		//fmt.Printf("%v -%v-> %v\n", signal.From.Name, signal.Type, signal.To.Name)
		signal.To.OnSignal2(signal.Type, signal.From, state, signals)
	}
}

var metricGlobal = utils.NewMetric("Global")

func DoWithInputPart02(world World) int {
	button := world.Button
	broadcaster := world.Broadcaster
	rxModule := world.Modules["rx"]
	aggregator := &Aggregator{}

	pushCount := 1
	state := collections.NewBitSet128()

	metricGlobal.Enable()

	buttonPressSignal := Signal{
		Type: Low,
		From: button,
		To:   broadcaster,
	}

	signals := collections.NewQueue[Signal]()

	for {
		signals.Push(buttonPressSignal)
		processSignals(&signals, aggregator, &state)

		if rxModule.InputsAggregator.LowCount > 0 {
			return pushCount
		}

		//fmt.Printf("#%15d rx: %v\n", pushCount, *rxModule.InputsAggregator)
		metricGlobal.TickCurrent(500_000, pushCount)

		rxModule.InputsAggregator.reset()
		pushCount++
	}
}

func getOrCreateModule(name string, modules Modules) *Module {
	return maps.GetOrCompute(modules, name, func(key string) *Module {
		return &Module{
			Name:             name,
			Type:             Terminal,
			InputsAggregator: &Aggregator{},
		}
	})
}

func ParseInput(r io.Reader) World {
	// prepend "button" module
	buttonLine := strings.NewReader("button -> broadcaster\n")
	r = io.MultiReader(buttonLine, r)

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

	// find starting state index for each module
	// state size is 89
	stateIndex := 0
	for _, module := range modules {
		module.StateIndex = stateIndex

		switch module.Type {
		case FlipFlop:
			stateIndex++
		case Conjunction:
			stateIndex += len(module.InputModules)
		}
	}

	return World{
		Modules:     modules,
		Button:      modules["button"],
		Broadcaster: modules["broadcaster"],
	}
}
