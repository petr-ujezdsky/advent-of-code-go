package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Rule struct {
	Name           string
	Range1, Range2 utils.IntervalI
}

func (r Rule) IsValid(v int) bool {
	return r.Range1.Contains(v) || r.Range2.Contains(v)
}

type Ticket []int

func (t Ticket) IsValid(rules []Rule) (bool, int) {
	for _, value := range t {
		invalidCount := 0
		for _, rule := range rules {
			if !rule.IsValid(value) {
				invalidCount++
			}
		}

		if invalidCount == len(rules) {
			return false, value
		}
	}

	return true, 0
}

type World struct {
	Rules        []Rule
	MyTicket     Ticket
	OtherTickets []Ticket
}

func DoWithInputPart01(world World) int {
	invalidValuesSum := 0

	for _, ticket := range world.OtherTickets {
		if ok, value := ticket.IsValid(world.Rules); !ok {
			invalidValuesSum += value
		}
	}

	return invalidValuesSum
}

func DoWithInputPart02(world World) int {
	return 0
}

func parseInterval(str string) utils.IntervalI {
	numbers := utils.ExtractInts(str, false)
	return utils.IntervalI{
		Low:  numbers[0],
		High: numbers[1],
	}
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// rules
	var rules []Rule
	for scanner.Scan() && scanner.Text() != "" {
		mainParts := strings.Split(scanner.Text(), ": ")
		name := mainParts[0]

		intervals := strings.Split(mainParts[1], " or ")

		rules = append(rules, Rule{
			Name:   name,
			Range1: parseInterval(intervals[0]),
			Range2: parseInterval(intervals[1]),
		})
	}

	// my ticket
	scanner.Scan()
	scanner.Scan()
	myTicket := Ticket(utils.ExtractInts(scanner.Text(), false))

	// other tickets
	scanner.Scan()
	scanner.Scan()
	var otherTickets []Ticket
	for scanner.Scan() {
		otherTickets = append(otherTickets, utils.ExtractInts(scanner.Text(), false))
	}

	return World{
		Rules:        rules,
		MyTicket:     myTicket,
		OtherTickets: otherTickets,
	}
}
