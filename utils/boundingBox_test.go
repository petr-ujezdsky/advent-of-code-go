package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoundingBox_Volume(t *testing.T) {
	type fields struct {
		XInterval IntervalI
		YInterval IntervalI
		ZInterval IntervalI
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"", fields{XInterval: IntervalI{}, YInterval: IntervalI{}, ZInterval: IntervalI{}}, 1},
		{"", fields{XInterval: IntervalI{High: 2}, YInterval: IntervalI{High: 3}, ZInterval: IntervalI{High: 4}}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b1 := BoundingBox{
				XInterval: tt.fields.XInterval,
				YInterval: tt.fields.YInterval,
				ZInterval: tt.fields.ZInterval,
			}
			assert.Equalf(t, tt.want, b1.Volume(), "Volume()")
		})
	}
}
