package main

import (
	"fmt"
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

func TestIsValid(t *testing.T) {
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

func TestIsValidWithCycles(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	validator := world.Validator
	validator8 := world.Validators[8].Validator

	type args struct {
		message    string
		validator  MessageValidator
		validator8 MessageValidator
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa", validator, validator8}, false},
		{"", args{"bbabbbbaabaabba", validator, validator8}, true},
		{"", args{"babbbbaabbbbbabbbbbbaabaaabaaa", validator, validator8}, true},
		{"", args{"aaabbbbbbaaaabaababaabababbabaaabbababababaaa", validator, validator8}, true},
		{"", args{"bbbbbbbaaaabbbbaaabbabaaa", validator, validator8}, true},
		{"", args{"bbbababbbbaaaaaaaabbababaaababaabab", validator, validator8}, true},
		{"", args{"ababaaaaaabaaab", validator, validator8}, true},
		{"", args{"ababaaaaabbbaba", validator, validator8}, true},
		{"", args{"baabbaaaabbaaaababbaababb", validator, validator8}, true},
		{"", args{"abbbbabbbbaaaababbbbbbaaaababb", validator, validator8}, true},
		{"", args{"aaaaabbaabaaaaababaa", validator, validator8}, true},
		{"", args{"aaaabbaaaabbaaa", validator, validator8}, false},
		{"", args{"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa", validator, validator8}, true},
		{"", args{"babaaabbbaaabaababbaabababaaab", validator, validator8}, false},
		{"", args{"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba", validator, validator8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsValidWithCycles(tt.args.message, tt.args.validator, tt.args.validator8), "IsValid(%v, %v)", tt.args.message, tt.args.validator)
		})
	}
}

func TestPrint(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)
	visitor := func(variant string) { fmt.Println(variant) }

	//world.Validators[4].Validator.Print("", true, visitor)
	//world.Validators[15].Validator.Print("", true, visitor)
	//world.Validators[8].Validator.Print("", true, visitor)
	//world.Validators[42].Validator.Print("", true, visitor)
	world.Validators[31].Validator.Print("", true, visitor)
	//world.Validators[11].Validator.Print("", true, visitor)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 12, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-02.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 405, result)
}
