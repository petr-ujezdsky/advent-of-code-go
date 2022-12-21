package utils

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

type UInteger interface {
	uint8 | uint16 | uint32 | uint64
}

type BitSet[T UInteger] struct {
	mask T
}

type BitSet8 = BitSet[uint8]
type BitSet16 = BitSet[uint16]
type BitSet32 = BitSet[uint32]
type BitSet64 = BitSet[uint64]

func NewEmptyBitSet[T UInteger]() BitSet[T] {
	return NewBitSetInitialized[T](0)
}

func NewFullBitSet8() BitSet8 {
	return NewBitSetInitialized[uint8](math.MaxUint8)
}

func NewFullBitSet16() BitSet16 {
	return NewBitSetInitialized[uint16](math.MaxUint16)
}

func NewFullBitSet32() BitSet32 {
	return NewBitSetInitialized[uint32](math.MaxUint32)
}

func NewFullBitSet64() BitSet64 {
	return NewBitSetInitialized[uint64](math.MaxUint64)
}

func NewBitSetInitialized[T UInteger](mask T) BitSet[T] {
	return BitSet[T]{mask}
}

func (s *BitSet[T]) Contains(value int) bool {
	return (s.mask & (1 << value)) != 0
}

func (s *BitSet[T]) Push(value int) {
	s.mask |= 1 << value
}

func (s *BitSet[T]) Remove(value int) {
	s.mask &= ^(1 << value)
}

func (s *BitSet[T]) Clone() BitSet[T] {
	return BitSet[T]{s.mask}
}

func (s *BitSet[T]) String() string {
	bitsCount := int(unsafe.Sizeof(*s)) * 8
	format := "%." + strconv.Itoa(bitsCount) + "b"
	return string(Reverse([]rune(fmt.Sprintf(format, s.mask))))
}
