package fixed

import (
	"math/bits"
)

// 72/56 fixed-point value
type Fixed struct {
	lo, hi uint64
}

var fixedRawval1 = rawfixed(1)
var fixedOne = rawfixed(oneValue)

const unsignMask = uint64(1)<<63 - 1
const signMask = uint64(1) << 63

func sign_(x int64) uint64 {
	return uint64(x) & signMask
}

func abs_(x int64) uint64 {
	xs := x >> 63
	return uint64((x ^ xs) - xs)
}

func rawfixed(x int64) Fixed {
	r := Fixed{lo: abs_(x), hi: sign_(x)}
	return r
}

func (x Fixed) fixed56() int64 {
	if (x.hi&unsignMask) != 0 || (x.lo&signMask) != 0 {
		panic(ErrOverflow)
	}
	return int64(x.lo) * x.sign()
}

func fixed(x int64) Fixed {
	v, s := abs_(x), sign_(x)
	return Fixed{lo: v << 56, hi: (v >> 8) | s}
}

func ufixed(x uint64) Fixed {
	return Fixed{lo: x << 56, hi: x >> 8}
}

func fixed_(x uint64, sign uint64) Fixed {
	return Fixed{lo: x, hi: sign}
}

func (x Fixed) integer() int64 {
	if x.bitlen() > 56+63 {
		panic(ErrOverflow)
	}
	sign := (int64(x.hi)>>63)*2 + 1
	return int64(x.hi<<8|x.lo>>56) * sign
}

func (x Fixed) fraction() uint64 {
	return x.lo & uint64(fracMask)
}

func (x Fixed) round() int64 {
	if x.bitlen() > 56+63 {
		panic(ErrOverflow)
	}
	sign := (int64(x.hi)>>63)*2 + 1
	v := int64(x.hi<<8|x.lo>>56)*sign + int64((x.lo>>55)&1)*sign
	return v
}

func (x Fixed) floor() int64 {
	if x.bitlen() > 56+63 {
		panic(ErrOverflow)
	}
	sign := (int64(x.hi)>>63)*2 + 1
	v := int64(x.hi<<8|x.lo>>56) * sign
	if sign < 0 && x.lo&uint64(fracMask) != 0 {
		return v - 1
	}
	return v
}

func (x Fixed) ceil() int64 {
	if x.bitlen() > 56+63 {
		panic(ErrOverflow)
	}
	sign := (int64(x.hi)>>63)*2 + 1
	v := int64(x.hi<<8|x.lo>>56) * sign
	if sign > 0 && x.lo&uint64(fracMask) != 0 {
		v += 1
	}
	return v
}

func (x Fixed) setSignAs(y Fixed) Fixed {
	x.hi = x.hi&unsignMask | y.sign_()
	return x
}

func (x Fixed) sign() int64 {
	return (int64(x.hi)>>63)*2 + 1
}

func (x Fixed) sign_() uint64 {
	return x.hi & signMask
}

func (x Fixed) neg() Fixed {
	x.hi ^= signMask
	return x
}

func (x Fixed) abs() Fixed {
	return Fixed{x.lo, x.hi & unsignMask}
}

func (x Fixed) bit(i int) bool {
	switch {
	case i > 127:
		return false
	case i >= 64:
		return (x.hi>>(i-64))&1 != 0
	case i >= 0:
		return (x.hi>>i)&1 != 0
	}
	return false
}

func (x Fixed) bitlen() int {
	if v := x.hi & unsignMask; v != 0 {
		return bits.Len64(v) + 64
	}
	return bits.Len64(x.lo)
}

func (x Fixed) shlmax(m int) (Fixed, int) {
	if n := bits.LeadingZeros64(x.hi) - 1; n < m {
		m = n
	}
	return x.shl_(m), m
}

func (x Fixed) shl_(n int) Fixed {
	lo, hi := x.lo, x.hi
	if n > 64 {
		hi, lo = lo, 0
		n -= 64
	}
	hi, lo = hi<<n|lo>>(64-n), lo<<n
	return Fixed{lo: lo, hi: hi}
}

func (x Fixed) shl(n int) Fixed {
	lo, hi := x.lo, x.hi&unsignMask
	if n > 64 {
		if hi != 0 {
			panic(ErrOverflow)
		}
		hi, lo = lo, 0
		n -= 64
	}
	if hi>>(64-n)|((hi<<n)&signMask) != 0 {
		panic(ErrOverflow)
	}
	hi, lo = hi<<n|lo>>(64-n), lo<<n
	return Fixed{lo: lo, hi: hi | (x.hi & signMask)}
}

func (x Fixed) shr(n int) Fixed {
	lo, hi := x.lo, x.hi&unsignMask
	if n > 64 {
		hi, lo = 0, hi
		n -= 64
	}
	hi, lo = hi>>n, hi<<(64-n)|lo>>n
	return Fixed{lo: lo, hi: hi | (x.hi & signMask)}
}

func (x Fixed) iszero() bool {
	return x.hi&unsignMask|x.lo == 0
}

func (x Fixed) less(y Fixed) bool {
	return x.hi < y.hi || (x.hi == y.hi && x.lo < y.lo)
}

func (x Fixed) greater(y Fixed) bool {
	return x.hi > y.hi || (x.hi == y.hi && x.lo > y.lo)
}

func (x Fixed) equal(y Fixed) bool {
	return x.hi == y.hi && x.lo == y.lo
}
