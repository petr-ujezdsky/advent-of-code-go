package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type Evaluator = func(a, b int) int

var evaluators = map[rune]Evaluator{
	'+': func(a, b int) int { return a + b },
	'*': func(a, b int) int { return a * b },
}

type PriorityMerger = func(operand rune, last, current *Expression) *Expression

type Operation struct {
	Operand     rune
	Evaluator   Evaluator
	Left, Right *Expression
}

func (o Operation) Evaluate() int {
	left, right := o.Left.Evaluate(), o.Right.Evaluate()
	return o.Evaluator(left, right)
}

type Expression struct {
	Value     int
	Operation *Operation
}

func (e Expression) Evaluate() int {
	if e.Operation != nil {
		return e.Operation.Evaluate()
	}

	return e.Value
}

func ParseExpression(pos int, str string, merger PriorityMerger) (*Expression, int) {
	expressionStr := &strings.Builder{}
	for pos < len(str) {
		char := str[pos]

		if char == '(' {
			return ParseExpressions(pos+1, str, merger)
		}

		if char < '0' || char > '9' {
			break
		}

		expressionStr.WriteRune(rune(char))
		pos++
	}

	expression := &Expression{
		Value:     utils.ParseInt(expressionStr.String()),
		Operation: nil,
	}

	return expression, pos
}

func SamePriorityMerger(operand rune, last, current *Expression) *Expression {
	expression := &Expression{
		Value: 0,
		Operation: &Operation{
			Operand:   operand,
			Evaluator: evaluators[operand],
			Left:      last,
			Right:     current,
		},
	}

	return expression
}

func DifferentPriorityMerger(operand rune, last, current *Expression) *Expression {
	if operand == '*' {
		return SamePriorityMerger(operand, last, current)
	}

	if last.Operation == nil {
		return SamePriorityMerger(operand, last, current)
	}

	expression := &Expression{
		Value: 0,
		Operation: &Operation{
			Operand:   operand,
			Evaluator: evaluators[operand],
			Left:      last.Operation.Right,
			Right:     current,
		},
	}

	last.Operation.Right = expression

	return last
}

func ParseExpressions(pos int, str string, merger PriorityMerger) (*Expression, int) {
	var operand rune

	var last *Expression
	for pos < len(str) {
		var current *Expression
		current, pos = ParseExpression(pos, str, merger)

		if last != nil {
			last = merger(operand, last, current)
		} else {
			last = current
		}

		if pos >= len(str) {
			break
		}

		if str[pos] == ')' {
			pos++
			break
		}

		if str[pos] != ' ' {
			panic("Error")
		}

		pos++
		operand = rune(str[pos])
		pos += 2
	}

	return last, pos
}

func EvaluateExpression(str string, merger PriorityMerger) int {
	expression, _ := ParseExpressions(0, str, merger)
	return expression.Evaluate()
}

func DoWithInputPart01(expressions []string) int {
	sum := 0

	for _, expression := range expressions {
		sum += EvaluateExpression(expression, SamePriorityMerger)
	}

	return sum
}

func DoWithInputPart02(expressions []string) int {
	sum := 0

	for _, expression := range expressions {
		sum += EvaluateExpression(expression, DifferentPriorityMerger)
	}

	return sum
}

func ParseInput(r io.Reader) []string {
	return parsers.ParseToStrings(r)
}
