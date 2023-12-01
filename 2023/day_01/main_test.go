package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 4, len(world.Rows))
}

func Test_extractNumber(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"1abc2"}, 12},
		{"", args{"pqr3stu8vwx"}, 38},
		{"", args{"a1b2c3d4e5f"}, 15},
		{"", args{"treb7uchet"}, 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, extractNumber(tt.args.row), "extractNumber(%v)", tt.args.row)
		})
	}
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 142, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 54940, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}
