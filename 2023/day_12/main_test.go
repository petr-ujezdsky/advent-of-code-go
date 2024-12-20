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

	record := world.Records[0]
	assert.Equal(t, "???.###", record.ConditionsRaw)
	assert.Equal(t, []int{1, 1, 3}, record.GroupSizes)

	record = world.Records[1]
	assert.Equal(t, ".??..??...?##.", record.ConditionsRaw)
	assert.Equal(t, []int{1, 1, 3}, record.GroupSizes)

}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	ParseInput(reader)
	//
	//result := DoWithInputPart01(world)
	//assert.Equal(t, 21, result)
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
	assert.Equal(t, 3920437278260, result)
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
		{name: "", args: args{ParseRecord("??????#???#??? 1,8")}, want: 14},
		{name: "", args: args{ParseRecord("?.##????#?.?# 1,4,1,1")}, want: 1},
		{name: "", args: args{ParseRecord("?????.??.???. 1,1,1")}, want: 68},
		{name: "", args: args{ParseRecord("????????????? 1,1,1,2")}, want: 126},
		{name: "", args: args{ParseRecord("??????.??..? 2,1,2")}, want: 6},

		{name: "", args: args{Unfold(ParseRecord("???.### 1,1,3"), 5)}, want: 1},
		{name: "", args: args{Unfold(ParseRecord(".??..??...?##. 1,1,3"), 5)}, want: 16384},
		{name: "", args: args{Unfold(ParseRecord("?#?#?#?#?#?#?#? 1,3,1,6"), 5)}, want: 1},
		{name: "", args: args{Unfold(ParseRecord("????.#...#... 4,1,1"), 5)}, want: 16},
		{name: "", args: args{Unfold(ParseRecord("????.######..#####. 1,6,5"), 5)}, want: 2500},
		{name: "", args: args{Unfold(ParseRecord("?###???????? 3,2,1"), 5)}, want: 506250},
		{name: "", args: args{Unfold(ParseRecord("?#???#???????#????? 5,2,1,5"), 5)}, want: 4487214},
		// slow ones
		{name: "", args: args{Unfold(ParseRecord("????????????? 1,1,1,2"), 3)}, want: 17383860},
		{name: "", args: args{Unfold(ParseRecord(".?.?????#?????.???? 1,6,1,2,1"), 3)}, want: 116190},
		{name: "", args: args{Unfold(ParseRecord("?????#????#??????? 5,1,1,1,1,1"), 3)}, want: 47016},
		{name: "", args: args{Unfold(ParseRecord("?????????.???? 1,1,3,1"), 3)}, want: 1477765},
		{name: "", args: args{Unfold(ParseRecord("??.??????.??#?#????? 1,2,1,5,1,1"), 3)}, want: 1096784},
		{name: "", args: args{Unfold(ParseRecord(".???????.#.????????? 4,1,1,1,3,1"), 3)}, want: 350590},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculateArrangementsCount(tt.args.record), "calculateArrangementsCount(%v)", tt.args.record)
		})
	}
}

func TestUnfoldN(t *testing.T) {
	type args struct {
		record Record
		count  int
	}
	tests := []struct {
		name string
		args args
		want Record
	}{
		{name: "", args: args{record: ParseRecord("???.### 1,1,3"), count: 5}, want: ParseRecord("???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3")},
		{name: "", args: args{record: ParseRecord("???.### 1,1,3"), count: 2}, want: ParseRecord("???.###????.### 1,1,3,1,1,3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Unfold(tt.args.record, tt.args.count), "Unfold(%v,%v)", tt.args.record, tt.args.count)
		})
	}
}

func Test_calculateArrangementsCountUnfolded(t *testing.T) {
	type args struct {
		i      int
		record Record
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{i: 0, record: ParseRecord("???.### 1,1,3")}, want: 1},
		{name: "", args: args{i: 0, record: ParseRecord(".??..??...?##. 1,1,3")}, want: 16384},
		{name: "", args: args{i: 0, record: ParseRecord("?#?#?#?#?#?#?#? 1,3,1,6")}, want: 1},
		{name: "", args: args{i: 0, record: ParseRecord("????.#...#... 4,1,1")}, want: 16},
		{name: "", args: args{i: 0, record: ParseRecord("????.######..#####. 1,6,5")}, want: 2500},
		{name: "", args: args{i: 0, record: ParseRecord("?###???????? 3,2,1")}, want: 506250},
		{name: "", args: args{i: 0, record: ParseRecord("?#???#???????#????? 5,2,1,5")}, want: 4487214},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculateArrangementsCountUnfolded(tt.args.record, tt.args.i), "calculateArrangementsCountUnfolded(%v, %v)", tt.args.record, tt.args.i)
		})
	}
}

func Benchmark_calculateArrangementsCount_1(b *testing.B) {
	record := Unfold(ParseRecord("????????????? 1,1,1,2"), 1)
	for i := 0; i < b.N; i++ {
		assert.Equal(b, 126, calculateArrangementsCount(record))
	}
}

func Benchmark_calculateArrangementsCount_2(b *testing.B) {
	record := Unfold(ParseRecord("????????????? 1,1,1,2"), 2)
	for i := 0; i < b.N; i++ {
		assert.Equal(b, 43758, calculateArrangementsCount(record))
	}
}

func Benchmark_calculateArrangementsCount_3(b *testing.B) {
	record := Unfold(ParseRecord("????????????? 1,1,1,2"), 3)
	for i := 0; i < b.N; i++ {
		assert.Equal(b, 17383860, calculateArrangementsCount(record))
	}
}
