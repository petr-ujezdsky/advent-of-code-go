package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseExpression(t *testing.T) {
	expression, pos := ParseExpression(0, "1 + 2", SamePriorityMerger)
	fmt.Println(expression)
	fmt.Println(pos)
	assert.Equal(t, 1, expression.Evaluate())

	expression, pos = ParseExpression(4, "1 + 2", SamePriorityMerger)
	fmt.Println(expression)
	fmt.Println(pos)
	assert.Equal(t, 2, expression.Evaluate())

	expression, pos = ParseExpression(0, "(1 + 2)", SamePriorityMerger)
	fmt.Println(expression)
	fmt.Println(pos)
	assert.Equal(t, 3, expression.Evaluate())
}

func TestParseExpressions(t *testing.T) {
	expression, pos := ParseExpressions(0, "2 + 3 * 5", SamePriorityMerger)
	fmt.Println(expression)
	fmt.Println(expression.Evaluate())

	assert.Equal(t, 25, expression.Evaluate())
	assert.Equal(t, 9, pos)
}

func TestEvaluateExpression(t *testing.T) {
	type args struct {
		str    string
		merger PriorityMerger
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"2", SamePriorityMerger}, 2},
		{"", args{"2 + 3", SamePriorityMerger}, 5},
		{"", args{"2 + 3 * 5", SamePriorityMerger}, 25},
		{"", args{"2 + (3 * 5)", SamePriorityMerger}, 17},

		{"", args{"1 + 2 * 3 + 4 * 5 + 6", SamePriorityMerger}, 71},
		{"", args{"1 + (2 * 3) + (4 * (5 + 6))", SamePriorityMerger}, 51},
		{"", args{"2 * 3 + (4 * 5)", SamePriorityMerger}, 26},
		{"", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)", SamePriorityMerger}, 437},
		{"", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", SamePriorityMerger}, 12240},
		{"", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", SamePriorityMerger}, 13632},

		{"", args{"2", DifferentPriorityMerger}, 2},
		{"", args{"2 + 3", DifferentPriorityMerger}, 5},
		{"", args{"2 + 3 * 5", DifferentPriorityMerger}, 25},
		{"", args{"2 + (3 * 5)", DifferentPriorityMerger}, 17},

		{"", args{"1 + 2 * 3 + 4 * 5 + 6", DifferentPriorityMerger}, 231},
		{"", args{"1 + (2 * 3) + (4 * (5 + 6))", SamePriorityMerger}, 51},
		{"", args{"2 * 3 + (4 * 5)", DifferentPriorityMerger}, 46},
		{"", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)", DifferentPriorityMerger}, 1445},
		{"", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", DifferentPriorityMerger}, 669060},
		{"", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", DifferentPriorityMerger}, 23340},

		// problematic
		{"", args{"(2 * 1 + 2) + 3", DifferentPriorityMerger}, 9},
		{"", args{"(2 + 4 * 9)", DifferentPriorityMerger}, 54},
		{"", args{"(6 + 9 * 8 + 6)", DifferentPriorityMerger}, 210},
		{"", args{"(6 + 9 * 8 + 6) + 6", DifferentPriorityMerger}, 216},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, EvaluateExpression(tt.args.str, tt.args.merger), "EvaluateExpression(%v)", tt.args.str)
		})
	}
}

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	assert.Equal(t, 377, len(expressions))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	result := DoWithInputPart01(expressions)
	assert.Equal(t, 5019432542701, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	result := DoWithInputPart02(expressions)
	assert.Equal(t, 70518821989947, result)
}
