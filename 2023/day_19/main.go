package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strings"
)

type WorkFlowType int

const (
	TypeNormal WorkFlowType = iota
	TypeAccepts
	TypeRejects
)

type Category int

const (
	CategoryX Category = iota
	CategoryM
	CategoryA
	CategoryS
)

type Part struct {
	Ratings [4]int
}

type Workflow struct {
	Name       string
	Conditions []Condition
	Fallback   *Workflow
	Type       WorkFlowType
}

type Condition struct {
	Category Category
	Operand  rune
	Amount   int
	Next     *Workflow
}

type World struct {
	Workflows map[string]*Workflow
	Start     *Workflow
	Parts     []Part
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func toCategory(str string) Category {
	switch str {
	case "x":
		return CategoryX
	case "m":
		return CategoryM
	case "a":
		return CategoryA
	case "s":
		return CategoryS

	}

	panic("Unknown category " + str)
}

var partRegex = regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

func ParsePart(str string) Part {
	// {x=5,m=71,a=885,s=445}
	ratings := partRegex.FindStringSubmatch(str)

	return Part{Ratings: [4]int{
		utils.ParseInt(ratings[1]),
		utils.ParseInt(ratings[2]),
		utils.ParseInt(ratings[3]),
		utils.ParseInt(ratings[4]),
	}}
}

var conditionRegex = regexp.MustCompile(`([xmas])([<>])(\d+):(.+)`)

func ParseCondition(str string, workflows map[string]*Workflow) Condition {
	parts := conditionRegex.FindStringSubmatch(str)

	category := toCategory(parts[1])
	operand := rune(parts[2][0])
	amount := utils.ParseInt(parts[3])
	next := getOrCreateWorkflow(parts[4], workflows)

	return Condition{
		Category: category,
		Operand:  operand,
		Amount:   amount,
		Next:     next,
	}
}

func ParseWorkFlow(str string, workflows map[string]*Workflow) {
	mainParts := strings.Split(str, "{")

	name := mainParts[0]
	conditionsRaw := strings.Split(mainParts[1][:len(mainParts[1])-1], ",")
	// proper conditions
	conditions := slices.Map(conditionsRaw[:len(conditionsRaw)-1], func(s string) Condition {
		return ParseCondition(s, workflows)
	})
	// fallback
	fallback := getOrCreateWorkflow(conditionsRaw[len(conditionsRaw)-1], workflows)

	workflow := getOrCreateWorkflow(name, workflows)

	workflow.Conditions = conditions
	workflow.Fallback = fallback
}

func getOrCreateWorkflow(name string, workflows map[string]*Workflow) *Workflow {
	if workflow, ok := workflows[name]; ok {
		return workflow
	}

	workFlowType := TypeNormal
	switch name {
	case "A":
		workFlowType = TypeAccepts
	case "R":
		workFlowType = TypeRejects
	}

	workflow := &Workflow{
		Name: name,
		Type: workFlowType,
	}

	workflows[name] = workflow

	return workflow
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// workflows
	workflows := make(map[string]*Workflow)
	for scanner.Scan() && len(scanner.Text()) > 0 {
		ParseWorkFlow(scanner.Text(), workflows)
	}

	// parts
	var parts []Part
	for scanner.Scan() {
		parts = append(parts, ParsePart(scanner.Text()))
	}

	return World{
		Workflows: workflows,
		Start:     workflows["in"],
		Parts:     parts,
	}
}
