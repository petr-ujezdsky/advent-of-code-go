package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitSet_All(t *testing.T) {
	bitSet := NewEmptyBitSet[uint8]()

	for i := 0; i < 8; i++ {
		assert.Equal(t, false, bitSet.Contains(i))
	}

	bitSet.Push(7)

	for i := 0; i < 7; i++ {
		assert.Equal(t, false, bitSet.Contains(i))
	}
	assert.Equal(t, true, bitSet.Contains(7))
	assert.Equal(t, "00000001", bitSet.String())

	bitSet.Remove(7)
	assert.Equal(t, false, bitSet.Contains(7))
}

func TestNewFullBitSet8(t *testing.T) {
	bitSet := NewFullBitSet8()
	assert.Equal(t, "11111111", bitSet.String())
}

func TestNewFullBitSet16(t *testing.T) {
	bitSet := NewFullBitSet16()
	assert.Equal(t, "1111111111111111", bitSet.String())
}

func TestNewFullBitSet32(t *testing.T) {
	bitSet := NewFullBitSet32()
	assert.Equal(t, "11111111111111111111111111111111", bitSet.String())
}

func TestNewFullBitSet64(t *testing.T) {
	bitSet := NewFullBitSet64()
	assert.Equal(t, "1111111111111111111111111111111111111111111111111111111111111111", bitSet.String())
}
