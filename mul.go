package fixed

import "math/bits"

func mul(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	hi = hi - uint64((x>>63)&y) - uint64((y>>63)&x)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	if carry != 0 || hi>>63 != hi>>(fracBits-1)&1 {
		panic(ErrOverflow)
	}
	return int64(hi<<(64-fracBits) | (lo >> fracBits))
}

// Mul returns x*y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Mul(y Fixed) Fixed {
	return Fixed{mul(x.int64, y.int64)}
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

func mul54(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	hi = hi - uint64((x>>63)&y) - uint64((y>>63)&x)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	if carry != 0 || hi>>63 != hi>>(54-1)&1 {
		panic(ErrOverflow)
	}
	return int64(hi<<10 | lo>>54)
}

func mul54s(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	hi = hi - uint64((x>>63)&y) - uint64((y>>63)&x)
	return int64(hi<<10 | lo>>54)
}

func mul54u(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	return int64(hi<<10 | lo>>54)
}

func (x Fixed) Mul54(y int64) Fixed {
	return Fixed{mul54(x.int64, y)}
}

func mulDiv(a, b, c int64) int64 {
	hi, lo := bits.Mul64(uint64(a), uint64(b))
	hi = hi - uint64((a>>63)&b) - uint64((b>>63)&a)
	xs := int64(hi) >> 63
	ys := c >> 63
	y := uint64((c ^ ys) - ys)
	lo, borrow := bits.Sub64(lo^uint64(xs), uint64(xs), 0)
	hi, _ = bits.Sub64(hi^uint64(xs), uint64(xs), borrow)
	// will panic when divides by zero or occurs overflow
	v, rem := bits.Div64(hi, lo, y)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ [0,1]
	// round to near
	v, _ = bits.Add64(v, (rem<<1)/y, 0)
	return int64(v) * ((xs^ys)*2 + 1)
}

func (x Fixed) MulDiv(y, d Fixed) Fixed {
	return Fixed{mulDiv(x.int64, y.int64, d.int64)}
}
