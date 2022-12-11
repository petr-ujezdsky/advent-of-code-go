package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Vector2i = utils.Vector2i

type Step struct {
	Dir       rune
	DirVector utils.Vector2i
	Amount    int
}

func moveTail(head, tail Vector2i, visited map[Vector2i]struct{}) Vector2i {
	for true {
		dir := head.Subtract(tail)
		ones := dir.Signum()

		// exactly 1 step far -> exit
		if dir == ones {
			return tail
		}

		tail = tail.Add(ones)

		if visited != nil {
			visited[tail] = struct{}{}
		}
	}

	panic("Should not happen")
}

func printState(head Vector2i, tails []Vector2i) {
	size := 35
	m := utils.NewMatrix2iPopulated(size, size, 0)
	offset := Vector2i{size / 2, size / 2}

	// origin
	m.SetV(offset, -1)

	// tails
	for i, tail := range utils.Reverse(tails) {
		m.SetV(tail.InvY().Add(offset), len(tails)-1-i+1)
	}

	// head
	m.SetV(head.InvY().Add(offset), 99)

	fmt.Println(m.StringFmt(utils.FmtBooleanCustom(utils.FmtConstant[int](" ."), utils.FmtFmt[int]("%2d"))))
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func printTrail(visited map[Vector2i]struct{}) {
	size := 35
	m := utils.NewMatrix2iPopulated(size, size, 0)
	offset := Vector2i{size / 2, size / 2}

	// origin
	m.SetV(offset, -1)

	for pos, _ := range visited {
		pos = pos.InvY().Add(offset)
		m.SetV(pos, 1)
	}

	fmt.Println(m.StringFmt(utils.FmtBoolean[int]))
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func DoWithInput(steps []Step, tailsCount int) int {
	head := Vector2i{0, 0}
	tails := make([]Vector2i, tailsCount)
	visited := make(map[Vector2i]struct{})
	visited[tails[0]] = struct{}{}

	printState(head, tails)
	for _, step := range steps {
		head = head.Add(step.DirVector)
		headTail := head

		fmt.Printf("Step: %v %v\n", string(step.Dir), step.Amount)

		for i, tail := range tails {
			if i == len(tails)-1 {
				tails[i] = moveTail(headTail, tail, visited)
			} else {
				tails[i] = moveTail(headTail, tail, nil)
			}

			headTail = tails[i]
		}

		printState(head, tails)
		printTrail(visited)

	}

	printTrail(visited)

	return len(visited)
}

func dir2vec(dir rune, amount int) Vector2i {
	switch dir {
	case 'U':
		return Vector2i{0, amount}
	case 'R':
		return Vector2i{amount, 0}
	case 'D':
		return Vector2i{0, -amount}
	case 'L':
		return Vector2i{-amount, 0}
	}
	panic("Unknown dir " + string(dir))
}

func ParseInput(r io.Reader) []Step {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var steps []Step
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		dir := rune(parts[0][0])
		amount := utils.ParseInt(parts[1])

		step := Step{
			Dir:       dir,
			DirVector: dir2vec(dir, amount),
			Amount:    amount,
		}

		steps = append(steps, step)
	}

	return steps
}
