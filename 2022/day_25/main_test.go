package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseSNAFU(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"0"}, 0},
		{"", args{"1"}, 1},
		{"", args{"2"}, 2},
		{"", args{"1="}, 3},
		{"", args{"1-"}, 4},
		{"", args{"10"}, 5},
		{"", args{"11"}, 6},
		{"", args{"12"}, 7},
		{"", args{"2="}, 8},
		{"", args{"2-"}, 9},
		{"", args{"20"}, 10},

		{"", args{"1=0"}, 15},
		{"", args{"1-0"}, 20},
		{"", args{"1=11-2"}, 2022},
		{"", args{"1-0---0"}, 12345},
		{"", args{"1121-1110-1=0"}, 314159265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseSNAFU(tt.args.str), "ParseSNAFU(%v)", tt.args.str)
		})
	}
}

func TestCreateSNAFU(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{0}, "0"},
		{"", args{1}, "1"},
		{"", args{2}, "2"},
		{"", args{3}, "1="},
		{"", args{4}, "1-"},
		{"", args{5}, "10"},
		{"", args{6}, "11"},
		{"", args{7}, "12"},
		{"", args{8}, "2="},
		{"", args{9}, "2-"},
		{"", args{10}, "20"},

		{"", args{15}, "1=0"},
		{"", args{20}, "1-0"},
		{"", args{2022}, "1=11-2"},
		{"", args{12345}, "1-0---0"},
		{"", args{314159265}, "1121-1110-1=0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateSNAFU(tt.args.n), "CreateSNAFU(%v)", tt.args.n)
		})
	}
}

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	assert.Equal(t, 0, len(snafuNumbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}
