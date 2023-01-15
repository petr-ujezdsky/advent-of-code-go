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

func ParseExpression(pos int, str string) (*Expression, int) {
	expressionStr := &strings.Builder{}
	for pos < len(str) {
		char := str[pos]

		if char == '(' {
			return ParseExpressions(pos+1, str)
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

func ParseExpressions(pos int, str string) (*Expression, int) {
	var operand rune

	var last *Expression
	for pos < len(str) {
		var current *Expression
		current, pos = ParseExpression(pos, str)

		if last != nil {
			expression := &Expression{
				Value: 0,
				Operation: &Operation{
					Operand:   operand,
					Evaluator: evaluators[operand],
					Left:      last,
					Right:     current,
				},
			}

			last = expression
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

func EvaluateExpression(str string) int {
	expression, _ := ParseExpressions(0, str)
	return expression.Evaluate()
}

func DoWithInputPart01(expressions []string) int {
	sum := 0

	for _, expression := range expressions {
		sum += EvaluateExpression(expression)
	}

	return sum
}

func DoWithInputPart02(expressions []string) int {
	return len(expressions)
}

func ParseInput(r io.Reader) []string {
	return parsers.ParseToStrings(r)
}
