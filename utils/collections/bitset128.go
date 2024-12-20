package collections

type BitSet128 struct {
	maskLow, maskHigh BitSet64
}

func NewBitSet128(values ...int) BitSet128 {
	bitSet := BitSet128{
		maskLow:  NewBitSet64(),
		maskHigh: NewBitSet64(),
	}

	for _, value := range values {
		bitSet.Push(value)
	}

	return bitSet
}

func NewFullBitSet128() BitSet128 {
	return BitSet128{
		maskLow:  NewFullBitSet64(),
		maskHigh: NewFullBitSet64(),
	}
}

func (s *BitSet128) getMask(value int) (*BitSet64, int) {
	if value < 64 {
		return &s.maskLow, 0
	}

	return &s.maskHigh, 64
}

func (s *BitSet128) Contains(value int) bool {
	mask, offset := s.getMask(value)
	return mask.Contains(value - offset)
}

func (s *BitSet128) Push(value int) {
	mask, offset := s.getMask(value)
	mask.Push(value - offset)
}

func (s *BitSet128) Remove(value int) {
	mask, offset := s.getMask(value)
	mask.Remove(value - offset)
}

func (s *BitSet128) Invert(value int) {
	mask, offset := s.getMask(value)
	mask.Invert(value - offset)
}

func (s *BitSet128) Clone() BitSet128 {
	return BitSet128{s.maskLow, s.maskHigh}
}

func (s *BitSet128) String() string {
	return s.maskLow.String() + s.maskHigh.String()
}
