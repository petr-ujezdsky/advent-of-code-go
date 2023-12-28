package collections

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
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

func newBitSet[T UInteger](values ...int) BitSet[T] {
	bitSet := newBitSetInitialized[T](0)
	for _, value := range values {
		bitSet.Push(value)
	}
	return bitSet
}

func NewBitSet8(values ...int) BitSet8 {
	return newBitSet[uint8](values...)
}

func NewBitSet64(values ...int) BitSet64 {
	return newBitSet[uint64](values...)
}

func NewFullBitSet8() BitSet8 {
	return newBitSetInitialized[uint8](math.MaxUint8)
}

func NewFullBitSet16() BitSet16 {
	return newBitSetInitialized[uint16](math.MaxUint16)
}

func NewFullBitSet32() BitSet32 {
	return newBitSetInitialized[uint32](math.MaxUint32)
}

func NewFullBitSet64() BitSet64 {
	return newBitSetInitialized[uint64](math.MaxUint64)
}

func newBitSetInitialized[T UInteger](mask T) BitSet[T] {
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

func (s *BitSet[T]) Invert(value int) {
	s.mask ^= 1 << value
}

func (s *BitSet[T]) Clone() BitSet[T] {
	return BitSet[T]{s.mask}
}

func (s *BitSet[T]) PushAll(s2 BitSet[T]) {
	s.mask = s.mask | s2.mask
}

func (s *BitSet[T]) And(s2 BitSet[T]) BitSet[T] {
	return BitSet[T]{mask: s.mask & s2.mask}
}

func (s *BitSet[T]) Or(s2 BitSet[T]) BitSet[T] {
	return BitSet[T]{mask: s.mask | s2.mask}
}

func (s *BitSet[T]) GetMask() T {
	return s.mask
}

func (s *BitSet[T]) Empty() bool {
	return s.mask == 0
}

func (s *BitSet[T]) String() string {
	bitsCount := int(unsafe.Sizeof(*s)) * 8
	format := "%." + strconv.Itoa(bitsCount) + "b"
	return string(slices.Reverse([]rune(fmt.Sprintf(format, s.mask))))
}
