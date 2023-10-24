package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type Cup struct {
	Label          int
	Previous, Next *Cup
}

func (c *Cup) String() string {
	result := &strings.Builder{}
	cup := c
	for {
		result.WriteString(strconv.Itoa(cup.Label))
		cup = cup.Next
		if cup == c {
			break
		}
	}

	return result.String()
}

type World struct {
	CupsByLabel map[int]*Cup
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(labels string) World {
	var firstCup, previousCup *Cup
	cupsByLabel := make(map[int]*Cup)

	for _, char := range labels {
		label := int(char - '0')
		cup := &Cup{Label: label}

		// connect current and previous
		if previousCup != nil {
			previousCup.Next = cup
			cup.Previous = previousCup
		}

		// store first cup
		if firstCup == nil {
			firstCup = cup
		}

		// store to index
		cupsByLabel[label] = cup

		previousCup = cup
	}

	// connect first and last cup
	firstCup.Previous = previousCup
	previousCup.Next = firstCup

	return World{CupsByLabel: cupsByLabel}
}
