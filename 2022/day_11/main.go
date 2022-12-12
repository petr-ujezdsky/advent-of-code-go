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

type ModuloNumber struct {
	Modulos map[int]int
}

func NewModuloNumber(i int, modulos map[int]struct{}) ModuloNumber {
	modulosMap := make(map[int]int, len(modulos))
	for modulo := range modulos {
		modulosMap[modulo] = 0
	}

	moduloNumber := ModuloNumber{modulosMap}

	moduloNumber.Eval(add, i)

	return moduloNumber
}

func (m ModuloNumber) Eval(op Operation, i int) {
	for n, residue := range m.Modulos {
		m.Modulos[n] = (op(residue, i)) % n
	}
}

func (m ModuloNumber) String() string {
	return fmt.Sprintf("%v", m.Modulos)
}

type Monkey struct {
	Id                      int
	Items                   []int
	ModuloItems             []ModuloNumber
	Operation               Operation
	OperationArg            int
	Test                    int
	MonkeyTrue, MonkeyFalse int
	Inspections             int
}

func (m *Monkey) AddItem(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) AddModuloItem(item ModuloNumber) {
	m.ModuloItems = append(m.ModuloItems, item)
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
		fmt.Printf("Monkey %d: (inspections %6d) %v\n", monkey.Id, monkey.Inspections, monkey.Items)
	}
}

func printModuloState(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d: (inspections %6d) %v\n", monkey.Id, monkey.Inspections, monkey.ModuloItems)
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

//func PlayKeepAwayMany(monkeys []*Monkey) int {
//	fmt.Println("Round 0")
//	printState(monkeys)
//	fmt.Println()
//	//for round := 1; round <= 10_000; round++ {
//	for round := 1; round <= 20; round++ {
//		debug := true //round%1000 == 0 || round == 100 || round == 20 || round == 1
//		if debug {
//			fmt.Printf("Round %2d\n", round)
//		}
//
//		for _, monkey := range monkeys {
//			for _, item := range monkey.Items {
//				worryLevel := monkey.Operation(item, monkey.OperationArg)
//				//worryLevel /= 3
//
//				var acceptingMonkey int
//				if worryLevel%monkey.Test == 0 {
//					acceptingMonkey = monkey.MonkeyTrue
//				} else {
//					acceptingMonkey = monkey.MonkeyFalse
//				}
//
//				fmt.Printf("%d (%2d) -> monkey %d\n", item, worryLevel, acceptingMonkey)
//				monkeys[acceptingMonkey].AddItem(worryLevel)
//				monkey.Inspections++
//			}
//			monkey.Items = []int{}
//		}
//
//		if debug {
//			printState(monkeys)
//			fmt.Println()
//		}
//	}
//
//	first, second := top2inspections(monkeys)
//
//	return first * second
//}

func PlayKeepAwayFast(monkeys []*Monkey) int {
	fmt.Println("Round 0")
	printModuloState(monkeys)
	fmt.Println()

	for round := 1; round <= 10_000; round++ {
		debug := round%1000 == 0 || round == 20 || round == 1
		if debug {
			fmt.Printf("Round %2d\n", round)
		}

		for _, monkey := range monkeys {
			for _, moduloItem := range monkey.ModuloItems {
				moduloItem.Eval(monkey.Operation, monkey.OperationArg)

				var acceptingMonkey int
				if moduloItem.Modulos[monkey.Test] == 0 {
					acceptingMonkey = monkey.MonkeyTrue
				} else {
					acceptingMonkey = monkey.MonkeyFalse
				}

				monkeys[acceptingMonkey].AddModuloItem(moduloItem)
				monkey.Inspections++
			}
			monkey.ModuloItems = []ModuloNumber{}
		}

		if debug {
			printModuloState(monkeys)
			fmt.Println()
		}
	}

	first, second := top2inspections(monkeys)

	return first * second
}

func ParseInput(r io.Reader) []*Monkey {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	modulos := make(map[int]struct{})

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

		// aggregate all possible modulos
		modulos[test] = struct{}{}
	}

	// convert all items into ModuloNumber items
	for _, monkey := range monkeys {
		var moduloItems []ModuloNumber
		for _, item := range monkey.Items {
			moduloItems = append(moduloItems, NewModuloNumber(item, modulos))
		}
		monkey.ModuloItems = moduloItems
	}

	return monkeys
}
