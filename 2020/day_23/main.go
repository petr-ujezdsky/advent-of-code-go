package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Cup struct {
	Label          int
	Previous, Next *Cup
}

func (c *Cup) String() string {
	return c.StringHighlighted(-1)
}

func (c *Cup) StringHighlighted(highlightedLabel int) string {
	result := &strings.Builder{}
	cup := c
	for {
		if cup != c {
			result.WriteString(" ")
		}

		if cup.Label == highlightedLabel {
			result.WriteString(fmt.Sprintf("(%d)", cup.Label))
		} else {
			result.WriteString(fmt.Sprintf("%d", cup.Label))
		}
		cup = cup.Next
		if cup == c {
			break
		}
	}

	return result.String()
}

func (c *Cup) CoupleAsNext(next *Cup) {
	c.Next = next
	next.Previous = c
}

type World struct {
	FirstCup    *Cup
	MaxLabel    int
	CupsByLabel map[int]*Cup
}

func DoWithInputPart01(world World, movesCount int) string {
	playMoves(world, movesCount)

	return describeCups(world.CupsByLabel)
}

func playMoves(world World, movesCount int) {
	currentCup := world.FirstCup

	for i := 0; i < movesCount; i++ {
		threeCups := [3]*Cup{currentCup.Next, currentCup.Next.Next, currentCup.Next.Next.Next}
		destinationLabel := findDestinationLabel(currentCup.Label, threeCups[0].Label, threeCups[1].Label, threeCups[2].Label, world.MaxLabel)
		destinationCup := world.CupsByLabel[destinationLabel]

		// remove three cups
		threeCups[0].Previous.CoupleAsNext(threeCups[2].Next)

		// insert three cups
		threeCups[2].CoupleAsNext(destinationCup.Next)
		destinationCup.CoupleAsNext(threeCups[0])

		//cupForLog := currentCup
		//for j := 0; j < i; j++ {
		//	cupForLog = cupForLog.Previous
		//}
		//fmt.Printf("-- move %d --\n", i+1)
		//fmt.Printf("-- cups %s\n", cupForLog.StringHighlighted(currentCup.Label))

		// select new current cup
		currentCup = currentCup.Next
	}
}

func findDestinationLabel(currentCupLabel, next1, next2, next3, maxLabel int) int {
	label := currentCupLabel

	for {
		label--

		if label == 0 {
			label = maxLabel
		}

		if label != currentCupLabel && label != next1 && label != next2 && label != next3 {
			return label
		}
	}
}

func describeCups(cups map[int]*Cup) string {
	result := &strings.Builder{}

	cup1 := cups[1]
	cup := cup1.Next

	for {
		result.WriteString(strconv.Itoa(cup.Label))
		cup = cup.Next
		if cup == cup1 {
			break
		}
	}

	return result.String()
}

func DoWithInputPart02(world World) int {
	playMoves(world, 10_000_000)

	cupA := world.CupsByLabel[1].Next
	cupB := cupA.Next

	return cupA.Label * cupB.Label
}

func EnlargeWorld(world *World, upToLabel int) {
	previousCup := world.FirstCup.Previous

	for label := len(world.CupsByLabel) + 1; label <= upToLabel; label++ {
		cup := &Cup{Label: label}

		// connect current and previous
		previousCup.CoupleAsNext(cup)

		// store to index
		world.CupsByLabel[label] = cup

		previousCup = cup
	}

	// connect first and last cup
	previousCup.CoupleAsNext(world.FirstCup)

	// update max label
	world.MaxLabel = upToLabel
}

func ParseInput(labels string) World {
	var firstCup, previousCup *Cup
	cupsByLabel := make(map[int]*Cup)

	for _, char := range labels {
		label := int(char - '0')
		cup := &Cup{Label: label}

		// connect current and previous
		if previousCup != nil {
			previousCup.CoupleAsNext(cup)
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
	previousCup.CoupleAsNext(firstCup)

	return World{
		FirstCup:    firstCup,
		MaxLabel:    9,
		CupsByLabel: cupsByLabel,
	}
}
