package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"sort"
	"strconv"
	"strings"
)

var cardStrength = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards string
	Bid   int

	CompareString string
	CardCounts    map[rune]int
	HandType      HandType
}

type World struct {
	Hands []*Hand
}

//func byHandTypeThenByCardsStrength(left, right *Hand) bool {
//	if left.HandType != left.HandType {
//		return left.HandType < left.HandType
//	}
//
//	for i, leftCard := range left.Cards {
//		rightCard := rune(right.Cards[i])
//
//		leftStrength := cardStrength[leftCard]
//		rightStrength := cardStrength[rightCard]
//
//		if leftStrength != rightStrength {
//			return leftStrength < rightStrength
//		}
//	}
//
//	panic("Hands are equal")
//}

func DoWithInputPart01(world World) int {
	// sort hands
	sort.Slice(world.Hands, func(i, j int) bool {
		//return byHandTypeThenByCardsStrength(world.Hands[i], world.Hands[j])
		return world.Hands[i].CompareString < world.Hands[j].CompareString
	})

	sum := 0
	for i, hand := range world.Hands {
		// calculate rank
		rank := i + 1

		// aggregate total winning
		sum += rank * hand.Bid
	}

	return sum
}

func DoWithInputPart02(world World) int {
	return 0
}

func getHandType(cardCounts map[rune]int) HandType {
	if len(cardCounts) == 5 {
		return HighCard
	}

	if len(cardCounts) == 4 {
		return OnePair
	}

	if len(cardCounts) == 3 {
		for _, count := range cardCounts {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	}

	if len(cardCounts) == 2 {
		for _, count := range cardCounts {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}

	if len(cardCounts) == 1 {
		return FiveOfAKind
	}

	panic("Could not determine hand type")
}

func createCompareString(cards string, handType HandType) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(int(handType)))

	for _, card := range cards {
		strength := cardStrength[card]

		sb.WriteRune(rune('a' + strength))
	}

	return sb.String()
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) *Hand {
		parts := strings.Split(str, " ")

		cards := parts[0]
		bid := utils.ParseInt(parts[1])

		cardCounts := make(map[rune]int)

		for _, card := range cards {
			cardCounts[card]++
		}

		handType := getHandType(cardCounts)
		compareString := createCompareString(cards, handType)

		return &Hand{
			Cards:         cards,
			Bid:           bid,
			CardCounts:    cardCounts,
			HandType:      handType,
			CompareString: compareString,
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Hands: items}
}
