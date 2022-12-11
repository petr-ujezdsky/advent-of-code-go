package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"sort"
	"strings"
)

type Operation func(int, int) int

type Monkey struct {
	Id                      int
	Items                   []int
	Operation               Operation
	OperationArg            int
	Test                    int
	MonkeyTrue, MonkeyFalse int
	Inspections             int
}

func (m *Monkey) AddItem(item int) {
	m.Items = append(m.Items, item)
}

func add(old, i int) int {
	return old + i
}

func multiply(old, i int) int {
	return old * i
}

func square(old, _ int) int {
	return old * old
}

func printState(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d: %v (inspections %v)\n", monkey.Id, monkey.Items, monkey.Inspections)
	}
}

func top2inspections(monkeys []*Monkey) (int, int) {
	counts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		counts[i] = monkey.Inspections
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	return counts[0], counts[1]
}

func PlayKeepAway(monkeys []*Monkey) int {
	for round := 0; round < 20; round++ {
		fmt.Printf("Round %2d\n", round+1)

		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				worryLevel := monkey.Operation(item, monkey.OperationArg)
				worryLevel /= 3
				if worryLevel%monkey.Test == 0 {
					monkeys[monkey.MonkeyTrue].AddItem(worryLevel)
				} else {
					monkeys[monkey.MonkeyFalse].AddItem(worryLevel)
				}
				monkey.Inspections++
			}
			monkey.Items = []int{}
		}
		printState(monkeys)
		fmt.Println()
	}

	first, second := top2inspections(monkeys)

	return first * second
}

func ParseInput(r io.Reader) []*Monkey {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var monkeys []*Monkey
	for scanner.Scan() {
		id := utils.ExtractInts(scanner.Text(), false)[0]

		scanner.Scan()
		items := utils.ExtractInts(scanner.Text(), false)

		scanner.Scan()
		op := multiply
		if strings.ContainsRune(scanner.Text(), '+') {
			op = add
		}

		arg := 0
		if strings.Contains(scanner.Text(), "* old") {
			op = square
		} else {
			arg = utils.ExtractInts(scanner.Text(), false)[0]
		}

		scanner.Scan()
		test := utils.ExtractInts(scanner.Text(), false)[0]

		scanner.Scan()
		monkeyTrue := utils.ExtractInts(scanner.Text(), false)[0]

		scanner.Scan()
		monkeyFalse := utils.ExtractInts(scanner.Text(), false)[0]
		scanner.Scan()

		monkey := Monkey{
			Id:           id,
			Items:        items,
			Operation:    op,
			OperationArg: arg,
			Test:         test,
			MonkeyTrue:   monkeyTrue,
			MonkeyFalse:  monkeyFalse,
		}

		monkeys = append(monkeys, &monkey)
	}

	return monkeys
}
