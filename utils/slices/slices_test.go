package slices

import (
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	type args[T any] struct {
		slice []T
		count int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{name: "", args: args[int]{slice: []int{1, 2, 3, 4}, count: 3}, want: []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}},
		{name: "", args: args[int]{slice: []int{1, 2, 3, 4}, count: 1}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.slice, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
