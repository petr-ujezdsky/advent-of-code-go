package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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
	Name             string
	Conditions       []Condition
	Type             WorkFlowType
	ParentCondition  *Condition
	ParentConditions []*Condition
}

type Condition struct {
	Category          Category
	Operand           rune
	Amount            int
	Next              *Workflow
	Previous          *Condition
	Owner             *Workflow
	TrueBox, FalseBox utils.BoundingBoxN
}

func (c Condition) Evaluate(part Part) bool {
	return c.TrueBox.Intervals[c.Category].Contains(part.Ratings[c.Category])
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
	// look at counts
	//visited := make(map[string]*Workflow)
	//
	//count := VisitAll(world.Start, visited)
	//
	//fmt.Printf("Visited %d / %d, count %d\n", len(visited), len(world.Workflows), count)

	// print tree
	//treeStr := tree.PrintTree(world.Start, func(node *Workflow) (string, []*Workflow) {
	//	var children []*Workflow
	//
	//	for _, condition := range node.Conditions {
	//		children = append(children, condition.Next)
	//	}
	//
	//	return node.Name, children
	//})
	//
	//fmt.Printf("%v\n", treeStr)
	//
	//// walk reverse A -> in
	accepting := world.Workflows["A"]

	//WalkReverse(accepting)
	ratingIntervals := WalkReverseAndIntersect(accepting)

	combinations := CountCombinations(ratingIntervals)
	return combinations

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
	}

	return count
}

func WalkReverse(workflow *Workflow) {
	if workflow == nil {
		fmt.Printf("\n")
		return
	}

	fmt.Printf(" -> %s", workflow.Name)
	conditionTrue := workflow.ParentCondition
	if conditionTrue == nil {
		fmt.Printf("\n")
		return
	}

	conditionFalse := conditionTrue.Previous
	for conditionFalse != nil {

		conditionFalse = conditionFalse.Previous
	}

	WalkReverse(conditionTrue.Owner)
}

func CountCombinations(results []utils.BoundingBoxN) int {
	// total combinations
	combinations := 0
	for _, ratingIntervals := range results {
		combinations += ratingIntervals.Volume()
	}

	// subtract duplicates (intersection)
	var intersections []utils.BoundingBoxN
	duplicates := 0
	for i, box1 := range results {
		for _, box2 := range results[i+1:] {
			if intersection, ok := box1.Intersection(box2); ok {
				duplicates += intersection.Volume()
				intersections = append(intersections, intersection)
			}
		}
	}

	fmt.Printf("Total %d, duplicates %d\n", combinations, duplicates)

	return combinations - duplicates
}

func WalkReverseAndIntersect(acceptingWorkflow *Workflow) []utils.BoundingBoxN {
	var ratingIntervals []utils.BoundingBoxN

	for _, conditionTrue := range acceptingWorkflow.ParentConditions {
		fullBoundingBox := fullBox()

		intersection := walkReverseAndIntersectCondition(conditionTrue, fullBoundingBox)

		ratingIntervals = append(ratingIntervals, intersection)
	}

	return ratingIntervals
}

func walkReverseAndIntersectCondition(conditionTrue *Condition, ratingIntervals utils.BoundingBoxN) utils.BoundingBoxN {
	if conditionTrue == nil {
		return ratingIntervals
	}

	// fulfill the condition
	if intersection, ok := ratingIntervals.Intersection(conditionTrue.TrueBox); ok {
		ratingIntervals = intersection
	}

	// neglect the previous conditions
	conditionFalse := conditionTrue.Previous
	for conditionFalse != nil {
		if intersection, ok := ratingIntervals.Intersection(conditionFalse.FalseBox); ok {
			ratingIntervals = intersection
		}

		conditionFalse = conditionFalse.Previous
	}

	return walkReverseAndIntersectCondition(conditionTrue.Owner.ParentCondition, ratingIntervals)
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

func fullBox() utils.BoundingBoxN {
	fullInterval := utils.IntervalI{Low: 1, High: 4000}
	fullIntervals := []utils.IntervalI{fullInterval, fullInterval, fullInterval, fullInterval}
	return utils.BoundingBoxN{Intervals: fullIntervals}
}

func emptyBox() utils.BoundingBoxN {
	emptyInterval := utils.IntervalI{Low: -2, High: -1}
	emptyIntervals := []utils.IntervalI{emptyInterval, emptyInterval, emptyInterval, emptyInterval}
	return utils.BoundingBoxN{Intervals: emptyIntervals}
}

var conditionRegex = regexp.MustCompile(`([xmas])([<>])(\d+):(.+)`)

func ParseCondition(str string, workflows map[string]*Workflow, previous *Condition, owner *Workflow) Condition {
	parts := conditionRegex.FindStringSubmatch(str)

	category := toCategory(parts[1])
	operand := rune(parts[2][0])
	amount := utils.ParseInt(parts[3])
	next := getOrCreateWorkflow(parts[4], workflows)

	trueBox := fullBox()
	falseBox := fullBox()

	if operand == '>' {
		trueBox.Intervals[category] = utils.IntervalI{Low: amount + 1, High: 4000}
		falseBox.Intervals[category] = utils.IntervalI{Low: 1, High: amount}
	} else {
		trueBox.Intervals[category] = utils.IntervalI{Low: 1, High: amount - 1}
		falseBox.Intervals[category] = utils.IntervalI{Low: amount, High: 4000}
	}

	condition := Condition{
		Category: category,
		Operand:  operand,
		Amount:   amount,
		Next:     next,
		Previous: previous,
		Owner:    owner,
		TrueBox:  trueBox,
		FalseBox: falseBox,
	}

	next.ParentCondition = &condition
	// for leaf 'A' nodes - it has multiple parents
	next.ParentConditions = append(next.ParentConditions, &condition)

	return condition
}

func ParseAlwaysTrueCondition(nextName string, workflows map[string]*Workflow, previous *Condition, owner *Workflow) Condition {
	next := getOrCreateWorkflow(nextName, workflows)

	condition := Condition{
		Category: CategoryX,
		Operand:  '>',
		Amount:   -1,
		Next:     next,
		Previous: previous,
		Owner:    owner,
		TrueBox:  fullBox(),
		FalseBox: emptyBox(),
	}

	next.ParentCondition = &condition
	// for leaf 'A' nodes - it has multiple parents
	next.ParentConditions = append(next.ParentConditions, &condition)

	return condition
}

func ParseWorkFlow(str string, workflows map[string]*Workflow) {
	mainParts := strings.Split(str, "{")

	name := mainParts[0]
	workflow := getOrCreateWorkflow(name, workflows)

	conditionsRaw := strings.Split(mainParts[1][:len(mainParts[1])-1], ",")

	// proper conditions
	var previousCondition *Condition
	conditions := make([]Condition, len(conditionsRaw))
	for i, conditionRaw := range conditionsRaw {
		var condition Condition
		if i == len(conditionsRaw)-1 {
			condition = ParseAlwaysTrueCondition(conditionRaw, workflows, previousCondition, workflow)
		} else {
			condition = ParseCondition(conditionRaw, workflows, previousCondition, workflow)
		}

		previousCondition = &condition
		conditions[i] = condition
	}

	workflow.Conditions = conditions
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
