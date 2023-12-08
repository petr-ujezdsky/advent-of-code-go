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

var cardStrengthPart1 = map[rune]int{
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

var cardStrengthPart2 = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
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

	CompareStringPart1 string
	CompareStringPart2 string
	CardCounts         map[rune]int
	HandType           HandType
}

type World struct {
	Hands []*Hand
}

func DoWithInputPart01(world World) int {
	// sort hands
	sort.Slice(world.Hands, func(i, j int) bool {
		return world.Hands[i].CompareStringPart1 < world.Hands[j].CompareStringPart1
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
	// sort hands
	sort.Slice(world.Hands, func(i, j int) bool {
		return world.Hands[i].CompareStringPart2 < world.Hands[j].CompareStringPart2
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

func getHandType(cardCounts map[rune]int) HandType {
	switch len(cardCounts) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, count := range cardCounts {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, count := range cardCounts {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
		return HighCard
	}

	panic("Could not determine hand type")
}

func calculateHandType(cards string) HandType {
	cardCounts := make(map[rune]int)

	for _, card := range cards {
		cardCounts[card]++
	}

	return getHandType(cardCounts)
}

func findHandTypeJoker(cards string) HandType {
	if !strings.ContainsRune(cards, 'J') {
		return calculateHandType(cards)
	}

	best := HighCard

	for card := range cardStrengthPart2 {
		if card == 'J' {
			continue
		}

		replaced := strings.Replace(cards, "J", string(card), 1)
		handType := findHandTypeJoker(replaced)

		if handType > best {
			best = handType
		}
	}

	return best
}

func createCompareString(cards string, handType HandType, cardStrength map[rune]int) string {
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
		compareStringPart1 := createCompareString(cards, handType, cardStrengthPart1)

		handTypeJoker := findHandTypeJoker(cards)
		compareStringPart2 := createCompareString(cards, handTypeJoker, cardStrengthPart2)

		return &Hand{
			Cards:              cards,
			Bid:                bid,
			CardCounts:         cardCounts,
			HandType:           handType,
			CompareStringPart1: compareStringPart1,
			CompareStringPart2: compareStringPart2,
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Hands: items}
}
