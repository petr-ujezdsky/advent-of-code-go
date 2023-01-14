package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
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

func isValidForPosition(rule Rule, position int, tickets []Ticket) bool {
	for _, ticket := range tickets {
		if !rule.IsValid(ticket[position]) {
			return false
		}
	}

	return true
}

func findExactPosition(rulesPositions map[string]map[int]struct{}) (string, int) {
	exactName := ""
	exactPosition := -1

	for name, positions := range rulesPositions {
		if len(positions) > 1 {
			continue
		}

		if exactName != "" {
			panic("Found more possible solutions")
		}

		exactName = name
		exactPosition = maps.FirstKey(positions)
	}

	if exactName == "" {
		panic("Found no solution")
	}

	return exactName, exactPosition
}

func findAllPositions(rulesPositions map[string]map[int]struct{}) map[string]int {
	exactPositions := make(map[string]int)

	for len(rulesPositions) > 0 {
		name, position := findExactPosition(rulesPositions)
		exactPositions[name] = position

		delete(rulesPositions, name)
		for _, positions := range rulesPositions {
			delete(positions, position)
		}
	}

	return exactPositions
}

func DoWithInputPart02(world World) int {
	// extract only valid tickets
	tickets := slices.Filter(world.OtherTickets, func(s Ticket) bool {
		valid, _ := s.IsValid(world.Rules)
		return valid
	})

	// add mine
	tickets = append(tickets, world.MyTicket)

	fieldsCount := len(world.MyTicket)

	// extract only relevant rules
	mainRules := slices.Filter(world.Rules, func(rule Rule) bool {
		return strings.Contains(rule.Name, "departure")
	})

	// find all possible positions
	rulesPossiblePositions := make(map[string]map[int]struct{})
	for _, rule := range world.Rules {
		possiblePositions := make(map[int]struct{})

		for position := 0; position < fieldsCount; position++ {
			if isValidForPosition(rule, position, tickets) {
				possiblePositions[position] = struct{}{}
			}
		}

		rulesPossiblePositions[rule.Name] = possiblePositions
	}

	for name, positions := range rulesPossiblePositions {
		fmt.Printf("%-19v: %v\n", name, maps.Keys(positions))
	}

	// determine positions
	positions := findAllPositions(rulesPossiblePositions)

	product := 1
	for _, rule := range mainRules {
		position := positions[rule.Name]
		product *= world.MyTicket[position]
	}

	return product
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
