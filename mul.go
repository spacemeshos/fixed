package fixed

import "math/bits"

// Mul returns x*y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Mul(y Fixed) Fixed {
	hi, lo := bits.Mul64(uint64(x.int64), uint64(y.int64))
	hi = hi - uint64((x.int64>>63)&y.int64) - uint64((y.int64>>63)&x.int64)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	if carry != 0 || hi>>63 != hi>>(fracBits-1)&1 {
		panic(ErrOverflow)
	}
	return Fixed{int64(hi<<(64-fracBits) | (lo >> fracBits))}
}

// UnsafeMul returns x*y in fixed-point arithmetic.
// Does not have overflow check
func (x Fixed) UnsafeMul(y Fixed) Fixed {
	hi, lo := bits.Mul64(uint64(x.int64), uint64(y.int64))
	hi = hi - uint64((x.int64>>63)&y.int64) - uint64((y.int64>>63)&x.int64)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return Fixed{int64(hi<<(64-fracBits) | (lo >> fracBits))}
}
