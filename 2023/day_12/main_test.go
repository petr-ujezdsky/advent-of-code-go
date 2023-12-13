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

	assert.Equal(t, 6, len(world.Records))
	assert.Equal(t, "???.###", world.Records[0].ConditionsRaw)
	assert.Equal(t, []int{1, 1, 3}, world.Records[0].GroupSizes)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 21, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 6935, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 525152, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_isValid(t *testing.T) {
	type args struct {
		conditions []rune
		groupSizes []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{conditions: []rune("#.#.###"), groupSizes: []int{1, 1, 3}}, want: true},
		{name: "", args: args{conditions: []rune("###.###"), groupSizes: []int{1, 1, 3}}, want: false},
		{name: "", args: args{conditions: []rune("#######"), groupSizes: []int{1, 1, 3}}, want: false},
		{name: "", args: args{conditions: []rune("......."), groupSizes: []int{1, 1, 3}}, want: false},
		{name: "", args: args{conditions: []rune("#.#.###.###"), groupSizes: []int{1, 1, 3}}, want: false},
		{name: "", args: args{conditions: []rune("#.#...."), groupSizes: []int{1, 1, 3}}, want: false},
		{name: "", args: args{conditions: []rune(".###.##.#..#"), groupSizes: []int{3, 2, 1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isValid(tt.args.conditions, tt.args.groupSizes), "isValid(%v, %v)", tt.args.conditions, tt.args.groupSizes)
		})
	}
}

func Test_calculateArrangementsCount(t *testing.T) {
	type args struct {
		record Record
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{ParseRecord("???.### 1,1,3")}, want: 1},
		{name: "", args: args{ParseRecord(".??..??...?##. 1,1,3")}, want: 4},
		{name: "", args: args{ParseRecord("?#?#?#?#?#?#?#? 1,3,1,6")}, want: 1},
		{name: "", args: args{ParseRecord("????.#...#... 4,1,1")}, want: 1},
		{name: "", args: args{ParseRecord("????.######..#####. 1,6,5")}, want: 4},
		{name: "", args: args{ParseRecord("?###???????? 3,2,1")}, want: 10},
		{name: "", args: args{ParseRecord("?#???#???????#????? 5,2,1,5")}, want: 4},

		{name: "", args: args{Unfold(ParseRecord("???.### 1,1,3"))}, want: 1},
		{name: "", args: args{Unfold(ParseRecord(".??..??...?##. 1,1,3"))}, want: 16384},
		{name: "", args: args{Unfold(ParseRecord("?#?#?#?#?#?#?#? 1,3,1,6"))}, want: 1},
		{name: "", args: args{Unfold(ParseRecord("????.#...#... 4,1,1"))}, want: 16},
		{name: "", args: args{Unfold(ParseRecord("????.######..#####. 1,6,5"))}, want: 2500},
		{name: "", args: args{Unfold(ParseRecord("?###???????? 3,2,1"))}, want: 506250},
		{name: "", args: args{Unfold(ParseRecord("?#???#???????#????? 5,2,1,5"))}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculateArrangementsCount(tt.args.record), "calculateArrangementsCount(%v)", tt.args.record)
		})
	}
}

func TestUnfold(t *testing.T) {
	type args struct {
		record Record
	}
	tests := []struct {
		name string
		args args
		want Record
	}{
		{name: "", args: args{record: ParseRecord("???.### 1,1,3")}, want: ParseRecord("???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Unfold(tt.args.record), "Unfold(%v)", tt.args.record)
		})
	}
}
