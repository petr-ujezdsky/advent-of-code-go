package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Operation = func(a, b int) int

var operations = func() map[rune]Operation {
	m := make(map[rune]Operation, 4)

	m['+'] = func(a, b int) int { return a + b }
	m['-'] = func(a, b int) int { return a - b }
	m['*'] = func(a, b int) int { return a * b }
	m['/'] = func(a, b int) int { return a / b }

	return m
}()

type Monkey struct {
	Name        string
	Value       *int
	Left, Right *Monkey
	Operation   Operation
}

func (m Monkey) GetValue() int {
	if m.Value == nil {
		value := m.Operation(m.Left.GetValue(), m.Right.GetValue())
		m.Value = &value
	}
	return *m.Value
}

func EvaluateRootMonkey(monkeys map[string]*Monkey) int {
	rootMonkey := monkeys["root"]

	return rootMonkey.GetValue()
}

func getOrCreateMonkey(name string, monkeys map[string]*Monkey) *Monkey {
	if monkey, ok := monkeys[name]; ok {
		return monkey
	}

	monkey := &Monkey{Name: name}
	monkeys[name] = monkey

	return monkey
}

func ParseInput(r io.Reader) map[string]*Monkey {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	monkeys := make(map[string]*Monkey)

	for scanner.Scan() {
		parts := strings.FieldsFunc(scanner.Text(), func(r rune) bool { return r == ' ' || r == ':' })

		name := parts[0]
		monkey := getOrCreateMonkey(name, monkeys)

		// value type
		if len(parts) == 2 {
			value := utils.ParseInt(parts[1])
			monkey.Value = &value
		} else {
			monkey.Left = getOrCreateMonkey(parts[1], monkeys)
			monkey.Operation = operations[[]rune(parts[2])[0]]
			monkey.Right = getOrCreateMonkey(parts[3], monkeys)
		}
	}

	return monkeys
}
