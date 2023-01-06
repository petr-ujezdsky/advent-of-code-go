package utils_test

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"os"
	"testing"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseToInts(t *testing.T) {
	reader, err := os.Open("test-data/parse-to-ints.txt")
	assert.Nil(t, err)

	parsed, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	expected := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	assert.Equal(t, expected, parsed)
}

func TestExtractInts(t *testing.T) {
	assert.Equal(t, []int{34, 60, 18, 25}, utils.ExtractInts("Hi there, I'm 34 years old and 60in tall. Today should be 18-25 degrees Celsius", false))
	assert.Equal(t, []int{34, 60, 18, -25}, utils.ExtractInts("Hi there, I'm 34 years old and 60in tall. Today should be 18-25 degrees Celsius", true))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, utils.Abs(5))
	assert.Equal(t, 5, utils.Abs(-5))
	assert.Equal(t, 0, utils.Abs(0))
	// does not work - overflows
	// assert.Equal(t, -math.MinInt, utils.Abs(math.MinInt))
}

func TestSignum(t *testing.T) {
	assert.Equal(t, -1, utils.Signum(-600))
	assert.Equal(t, -1, utils.Signum(-1))
	assert.Equal(t, 0, utils.Signum(0))
	assert.Equal(t, 1, utils.Signum(1))
	assert.Equal(t, 1, utils.Signum(20))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 1, utils.Max(0, 1))
	assert.Equal(t, 5, utils.Max(5, -1))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 0, utils.Min(0, 1))
	assert.Equal(t, -1, utils.Min(5, -1))
}

func TestArgMin(t *testing.T) {
	index, min := utils.ArgMin(5, 4, 3)
	assert.Equal(t, 2, index)
	assert.Equal(t, 3, min)

	index, min = utils.ArgMin([]int{5, -300, 80, 500}...)
	assert.Equal(t, 1, index)
	assert.Equal(t, -300, min)
}

func TestSumNtoM(t *testing.T) {
	assert.Equal(t, 1, utils.SumNtoM(0, 1))
	assert.Equal(t, 0, utils.SumNtoM(-50, 50))
	assert.Equal(t, 5050, utils.SumNtoM(0, 100))
	assert.Equal(t, 5050, utils.SumNtoM(1, 100))
	assert.Equal(t, -5050, utils.SumNtoM(-100, 0))
}

func TestClamp(t *testing.T) {
	assert.Equal(t, 0, utils.Clamp(-1, 0, 10))
	assert.Equal(t, 0, utils.Clamp(0, 0, 10))
	assert.Equal(t, 5, utils.Clamp(5, 0, 10))
	assert.Equal(t, 10, utils.Clamp(10, 0, 10))
	assert.Equal(t, 10, utils.Clamp(30, 0, 10))
}

func TestShallowCopy(t *testing.T) {
	data := []int{1, 2, 3}
	clone := slices.Clone(data)

	// modify original data
	data[0] = 9

	assert.Equal(t, 9, data[0])
	assert.Equal(t, 1, clone[0])
}

func TestCopy(t *testing.T) {
	source := []int{1, 2, 3}
	target := make([]int, 3)

	slices.Copy(source, target)

	assert.Equal(t, []int{1, 2, 3}, target)
}

func TestSubstring(t *testing.T) {
	type args struct {
		str  string
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"text", 1, 3}, "ex"},
		{"", args{"text", 1, 2}, "e"},
		{"", args{"text", 1, 1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, utils.Substring(tt.args.str, tt.args.from, tt.args.to), "Substring(%v, %v, %v)", tt.args.str, tt.args.from, tt.args.to)
		})
	}
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, slices.Reverse([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 1}, slices.Reverse([]int{1, 2, 3, 4}))
}

func TestRemoveUnordered(t *testing.T) {
	assert.Equal(t, []int{3, 2}, slices.RemoveUnordered([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{1, 3}, slices.RemoveUnordered([]int{1, 2, 3}, 1))
	assert.Equal(t, []int{1, 2}, slices.RemoveUnordered([]int{1, 2, 3}, 2))
}

func TestParseBinary8(t *testing.T) {
	type args struct {
		onesAndZeros string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"", args{"0"}, 0},
		{"", args{"00000000"}, 0},
		{"", args{"00000001"}, 1},
		{"", args{"00000010"}, 2},
		{"", args{"00000100"}, 4},
		{"", args{"00001000"}, 8},
		{"", args{"00010000"}, 16},
		{"", args{"00100000"}, 32},
		{"", args{"01000000"}, 64},
		{"", args{"10000000"}, 128},
		{"", args{"11111111"}, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, utils.ParseBinary8(tt.args.onesAndZeros), "ParseBinary8(%v)", tt.args.onesAndZeros)
		})
	}
}

func TestModFloor(t *testing.T) {
	type args struct {
		value int
		size  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{2, 10}, 2},
		{"", args{12, 10}, 2},
		{"", args{-2, 10}, 8},
		{"", args{-12, 10}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, utils.ModFloor(tt.args.value, tt.args.size), "ModFloor(%v, %v)", tt.args.value, tt.args.size)
		})
	}
}
