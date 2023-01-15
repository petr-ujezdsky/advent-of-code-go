package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseExpression(t *testing.T) {
	expression, pos := ParseExpression(0, "1 + 2")
	fmt.Println(expression)
	fmt.Println(pos)
}

//func TestParseOperation(t *testing.T) {
//	operation, _ := ParseOperation(0, "2 + 3 * 5")
//	fmt.Println(operation)
//	fmt.Println(operation.Evaluate())
//}

func TestParseExpressions(t *testing.T) {
	expression, pos := ParseExpressions(0, "2 + 3 * 5")
	fmt.Println(expression)
	fmt.Println(expression.Evaluate())

	assert.Equal(t, 25, expression.Evaluate())
	assert.Equal(t, 9, pos)
}

//func TestParseOperationEvaluate(t *testing.T) {
//	operation, pos := ParseOperation(0, "2 + 3")
//	assert.Equal(t, 5, operation.Evaluate())
//	assert.Equal(t, 5, pos)
//
//	operation, pos = ParseOperation(0, "2 * 3")
//	assert.Equal(t, 6, operation.Evaluate())
//	assert.Equal(t, 5, pos)
//}

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	assert.Equal(t, 0, len(expressions))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	result := DoWithInputPart01(expressions)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	result := DoWithInputPart01(expressions)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	result := DoWithInputPart02(expressions)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	expressions := ParseInput(reader)

	result := DoWithInputPart02(expressions)
	assert.Equal(t, 0, result)
}
