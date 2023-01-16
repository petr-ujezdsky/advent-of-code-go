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

	assert.Equal(t, 6, len(world.Validators))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 2, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 239, result)
}

//func Test_02_example2(t *testing.T) {
//	reader, err := os.Open("data-00-example-2.txt")
//	assert.Nil(t, err)
//
//	world := ParseInput(reader)
//
//	result := DoWithInputPart01(world)
//	assert.Equal(t, 3, result)
//}
//
//func Test_02_example3(t *testing.T) {
//	reader, err := os.Open("data-00-example-3.txt")
//	assert.Nil(t, err)
//
//	world := ParseInput(reader)
//
//	result := DoWithInputPart01(world)
//	assert.Equal(t, 12, result)
//}
//
//func Test_02_example3(t *testing.T) {
//	reader, err := os.Open("data-00-example-3.txt")
//	assert.Nil(t, err)
//
//	world := ParseInput(reader)
//
//	result := DoWithInputPart01(world)
//	assert.Equal(t, 12, result)
//}

func TestIsValid2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	validator := ParseInput(reader).Validator

	type args struct {
		message   string
		validator MessageValidator
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa", validator}, false},
		{"", args{"bbabbbbaabaabba", validator}, true},
		{"", args{"babbbbaabbbbbabbbbbbaabaaabaaa", validator}, false},
		{"", args{"aaabbbbbbaaaabaababaabababbabaaabbababababaaa", validator}, false},
		{"", args{"bbbbbbbaaaabbbbaaabbabaaa", validator}, false},
		{"", args{"bbbababbbbaaaaaaaabbababaaababaabab", validator}, false},
		{"", args{"ababaaaaaabaaab", validator}, true},
		{"", args{"ababaaaaabbbaba", validator}, true},
		{"", args{"baabbaaaabbaaaababbaababb", validator}, false},
		{"", args{"abbbbabbbbaaaababbbbbbaaaababb", validator}, false},
		{"", args{"aaaaabbaabaaaaababaa", validator}, false},
		{"", args{"aaaabbaaaabbaaa", validator}, false},
		{"", args{"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa", validator}, false},
		{"", args{"babaaabbbaaabaababbaabababaaab", validator}, false},
		{"", args{"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba", validator}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsValid(tt.args.message, tt.args.validator), "IsValid(%v, %v)", tt.args.message, tt.args.validator)
		})
	}
}

func TestIsValid3(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	validator := ParseInput(reader).Validator

	type args struct {
		message   string
		validator MessageValidator
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa", validator}, false},
		{"", args{"bbabbbbaabaabba", validator}, true},
		{"", args{"babbbbaabbbbbabbbbbbaabaaabaaa", validator}, true},
		{"", args{"aaabbbbbbaaaabaababaabababbabaaabbababababaaa", validator}, true},
		{"", args{"bbbbbbbaaaabbbbaaabbabaaa", validator}, true},
		{"", args{"bbbababbbbaaaaaaaabbababaaababaabab", validator}, true},
		{"", args{"ababaaaaaabaaab", validator}, true},
		{"", args{"ababaaaaabbbaba", validator}, true},
		{"", args{"baabbaaaabbaaaababbaababb", validator}, true},
		{"", args{"abbbbabbbbaaaababbbbbbaaaababb", validator}, true},
		{"", args{"aaaaabbaabaaaaababaa", validator}, true},
		{"", args{"aaaabbaaaabbaaa", validator}, false},
		{"", args{"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa", validator}, true},
		{"", args{"babaaabbbaaabaababbaabababaaab", validator}, false},
		{"", args{"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba", validator}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsValid(tt.args.message, tt.args.validator), "IsValid(%v, %v)", tt.args.message, tt.args.validator)
		})
	}
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}
