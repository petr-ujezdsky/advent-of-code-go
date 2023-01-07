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

func TestBitSet_And(t *testing.T) {
	bitSet1 := NewEmptyBitSet[uint8]()
	bitSet2 := NewEmptyBitSet[uint8]()

	bitSet1.Push(1)
	bitSet1.Push(2)
	bitSet1.Push(3)

	bitSet2.Push(2)

	bitSet := bitSet1.And(bitSet2)
	assert.Equal(t, false, bitSet.Contains(1))
	assert.Equal(t, true, bitSet.Contains(2))
	assert.Equal(t, false, bitSet.Contains(3))
}

func TestBitSet_Or(t *testing.T) {
	bitSet1 := NewEmptyBitSet[uint8]()
	bitSet2 := NewEmptyBitSet[uint8]()

	bitSet1.Push(1)
	bitSet1.Push(2)
	bitSet1.Push(3)

	bitSet2.Push(4)

	bitSet := bitSet1.Or(bitSet2)
	assert.Equal(t, true, bitSet.Contains(1))
	assert.Equal(t, true, bitSet.Contains(2))
	assert.Equal(t, true, bitSet.Contains(3))
	assert.Equal(t, true, bitSet.Contains(4))
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
