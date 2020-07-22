package fixed

import (
	"math/bits"
)

func mul(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	hi = hi - uint64((x>>63)&y) - uint64((y>>63)&x)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	if carry != 0 || int64(hi)>>63 != int64(hi)>>(fracBits-1) {
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

func mulDiv(a, b, c int64) int64 {
	return div(mul(a, b), c)
}

func (x Fixed) MulDiv(y, d Fixed) Fixed {
	return Fixed{mulDiv(x.int64, y.int64, d.int64)}
}

func mul2div(a, b, x, y int64) int64 {
	// a * b
	hi, li := bits.Mul64(uint64(a), uint64(b))
	hi = hi - uint64((a>>63)&b) - uint64((b>>63)&a)

	// x * y
	hj, lj := bits.Mul64(uint64(x), uint64(y))
	hj = hj - uint64((x>>63)&y) - uint64((y>>63)&x)

	// signs
	si := int64(hi) >> 63 // sign i
	sj := int64(hj) >> 63 // sign j

	// abs(a*b)
	li, borrow := bits.Sub64(li^uint64(si), uint64(si), 0)
	hi, _ = bits.Sub64(hi^uint64(si), uint64(si), borrow)

	// use all most significant bits of result
	z := bits.Len64(uint64((int64(hj) ^ sj) - sj)) // z < 64
	if z <= fracBits {
		z = fracBits
	} else {
		li = hi<<(64-z+fracBits) | li>>(z-fracBits)
		hi = hi >> (z - fracBits)
	}

	// abs(x*y)
	d := uint64(abs(int64(hj<<(64-z) | lj>>z)))
	v, rem := bits.Div64(hi, li, d)

	// TODO
	v, _ = bits.Add64(v, (rem<<1)/d, 0)

	return int64(v) * ((si^sj)*2 + 1)
}
