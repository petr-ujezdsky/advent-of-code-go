package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitSet128_All(t *testing.T) {
	bitSet := NewBitSet128()

	for i := 0; i < 128; i++ {
		assert.Equal(t, false, bitSet.Contains(i))
	}

	// low bit mask
	bitSet.Push(7)

	for i := 0; i < 128; i++ {
		if i == 7 {
			assert.Equal(t, true, bitSet.Contains(7))
		} else {
			assert.Equal(t, false, bitSet.Contains(i))
		}
	}
	assert.Equal(t, "00000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", bitSet.String())

	bitSet.Remove(7)
	assert.Equal(t, false, bitSet.Contains(7))

	// high bit mask
	bitSet.Push(127)

	for i := 0; i < 127; i++ {
		assert.Equal(t, false, bitSet.Contains(i), "i=%v", i)
	}
	assert.Equal(t, true, bitSet.Contains(127))
	assert.Equal(t, "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001", bitSet.String())

	bitSet.Remove(127)
	assert.Equal(t, false, bitSet.Contains(127))
}

func TestNewFullBitSet128(t *testing.T) {
	bitSet := NewFullBitSet128()
	assert.Equal(t, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", bitSet.String())
}

func TestBitSet128_Invert(t *testing.T) {
	bitSet := NewBitSet128()

	bitSet.Push(3)
	assert.True(t, bitSet.Contains(3))

	bitSet.Invert(3)
	assert.False(t, bitSet.Contains(3))

	assert.False(t, bitSet.Contains(123))
	bitSet.Invert(123)
	assert.True(t, bitSet.Contains(123))
}
