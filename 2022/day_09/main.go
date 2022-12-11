package main

import (
	"bufio"
	_ "embed"
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

//
//func tailLocation(head, tail Vector2i) Vector2i {
//	dir := tail.Subtract(head)
//	dirAbs := dir.Abs()
//	ones := dir.Signum()
//
//	// move diagonally
//	if dirAbs.X == dirAbs.Y {
//		return ones
//	}
//
//	// find smaller part
//	i, _ := dirAbs.ArgMin()
//
//	// erase smaller part
//	return ones.Change(i, 0)
//}

func moveTail(head, tail Vector2i, visited map[Vector2i]struct{}) Vector2i {
	for true {
		dir := head.Subtract(tail)
		ones := dir.Signum()

		// exactly 1 step far -> exit
		if dir == ones {
			return tail
		}

		tail = tail.Add(ones)
		visited[tail] = struct{}{}
	}

	panic("Should not happen")
}

func DoWithInput(steps []Step) int {
	head := Vector2i{0, 0}
	tail := Vector2i{0, 0}
	visited := make(map[Vector2i]struct{})
	visited[tail] = struct{}{}

	for _, step := range steps {
		head = head.Add(step.DirVector)
		tail = moveTail(head, tail, visited)
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
