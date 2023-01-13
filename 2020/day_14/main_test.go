package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//func Test_merge(t *testing.T) {
//	addresses := []Address{
//		Address("1001"),
//		Address("1001"),
//		Address("0100"),
//	}
//
//	result := merge(addresses)
//	assert.Equal(t, Address("XX0X"), result)
//}

func TestAddress_And(t *testing.T) {
	type args struct {
		a2 Address
	}
	tests := []struct {
		name string
		a    Address
		args args
		want Address
	}{
		{"", Address("XX0X"), args{Address("1011")}, ""},
		{"", Address("XX0X"), args{Address("1001")}, Address("1001")},
		{"", Address("XX0X"), args{Address("10X1")}, Address("1001")},
		{"", Address("XX0X"), args{Address("1XX1")}, Address("1X01")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.a.And(tt.args.a2), "And(%v)", tt.args.a2)
		})
	}
}

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	assert.Equal(t, 4, len(items))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart01(items)
	assert.Equal(t, 165, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart01(items)
	assert.Equal(t, 9879607673316, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart02(items)
	assert.Equal(t, 208, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	items := ParseInput(reader)

	result := DoWithInputPart02(items)
	assert.Equal(t, 3435342392262, result)
}

// 3432511849446
// 3396163389894
// 2693975408037
// 8372322902
// 3435342392262
