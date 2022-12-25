package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strings"
)

type Vector2i = utils.Vector2i

type Step struct {
	Dir       rune
	DirVector utils.Vector2i
	Amount    int
}

func moveTails(rope []Vector2i, iHead int, visited map[Vector2i]struct{}) {
	// at the end
	if iHead == len(rope)-1 {
		return
	}

	for true {
		head := rope[iHead]
		tail := rope[iHead+1]

		dir := head.Subtract(tail)
		ones := dir.Signum()

		// exactly 1 step far -> exit
		if dir == ones {
			return
		}

		tail = tail.Add(ones)
		rope[iHead+1] = tail

		if iHead == len(rope)-2 {
			visited[tail] = struct{}{}
		}

		// after just 1 step, move remaining tails
		moveTails(rope, iHead+1, visited)
	}

	panic("Should not happen")
}

func printState(head Vector2i, tails []Vector2i) {
	size := 35
	m := utils.NewMatrixInt(size, size)
	offset := Vector2i{size / 2, size / 2}

	// origin
	m.SetV(offset, -1)

	// tails
	for i, tail := range slices.Reverse(tails) {
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
	m := utils.NewMatrixInt(size, size)
	offset := Vector2i{size / 2, size / 2}

	// origin
	m.SetV(offset, -1)

	for pos := range visited {
		pos = pos.InvY().Add(offset)
		m.SetV(pos, 1)
	}

	fmt.Println(m.StringFmt(utils.FmtBoolean[int]))
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func DoWithInput(steps []Step, tailsCount int) int {
	rope := make([]Vector2i, tailsCount)
	visited := make(map[Vector2i]struct{})
	visited[rope[0]] = struct{}{}

	//printState(head, tails)
	for _, step := range steps {
		fmt.Printf("Step: %v %v\n", string(step.Dir), step.Amount)
		rope[0] = rope[0].Add(step.DirVector)

		moveTails(rope, 0, visited)

		//printState(head, tails)
		//printTrail(visited)

	}

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
