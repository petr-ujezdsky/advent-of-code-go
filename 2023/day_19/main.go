package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strconv"
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

func (p Part) Sum() int {
	return p.Ratings[0] + p.Ratings[1] + p.Ratings[2] + p.Ratings[3]
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

func (c Condition) Evaluate(part Part) bool {
	switch c.Operand {
	case '<':
		return part.Ratings[c.Category] < c.Amount
	case '>':
		return part.Ratings[c.Category] > c.Amount
	}

	panic("Unknown operand")
}

type World struct {
	Workflows map[string]*Workflow
	Start     *Workflow
	Parts     []Part
}

func (w *Workflow) Resolve(part Part, path []*Workflow, fillPath bool) (bool, []*Workflow) {
	if fillPath {
		path = append(path, w)
	}

	switch w.Type {
	case TypeAccepts:
		return true, path
	case TypeRejects:
		return false, path
	case TypeNormal:
		// try conditions
		for _, condition := range w.Conditions {
			if condition.Evaluate(part) {
				return condition.Next.Resolve(part, path, fillPath)
			}
		}

		// none matched -> fallback
		return w.Fallback.Resolve(part, path, fillPath)
	}

	panic("Unknown type " + strconv.Itoa(int(w.Type)))
}

func DoWithInputPart01(world World) int {
	results := utils.ProcessParallel(world.Parts, func(part Part, i int) int {
		if accepted, _ := world.Start.Resolve(part, nil, false); accepted {
			return part.Sum()
		}

		return 0
	})

	sum := 0

	for result := range results {
		sum += result.Value
	}

	return sum
}

func DoWithInputPart02(world World) int {
	visited := make(map[string]*Workflow)

	count := VisitAll(world.Start, visited)

	fmt.Printf("Visited %d / %d, count %d\n", len(visited), len(world.Workflows), count)

	return 0
	//results := utils.ProcessParallel(world.Parts, func(part Part, i int) int {
	//	accepted, path := world.Start.Resolve(part, nil, true)
	//
	//	fmt.Printf("#%3d (index %3d) recursions: %2d accepted: %v\n", i+574, i, len(path), accepted)
	//
	//	if accepted {
	//		return part.Sum()
	//	}
	//
	//	return 0
	//})
	//
	//sum := 0
	//
	//for result := range results {
	//	sum += result.Value
	//}
	//
	//return sum
}

func VisitAll(workflow *Workflow, visited map[string]*Workflow) int {
	//if _, ok := visited[workflow.Name]; ok {
	//	return
	//}
	visited[workflow.Name] = workflow
	count := 0

	switch workflow.Type {
	case TypeAccepts:
	case TypeRejects:
		count++
	case TypeNormal:
		for _, condition := range workflow.Conditions {
			next := condition.Next
			count += VisitAll(next, visited)
		}

		if workflow.Fallback != nil {
			count += VisitAll(workflow.Fallback, visited)
		}
	}

	return count
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
