package fixed

import (
	"math"
	"math/bits"
)

const (
	fracBits        = 56
	fracMask        = oneValue - 1
	halfValue       = oneValue >> 1
	oneValue  int64 = 1 << fracBits
)

const ln2 = int64(0xb17217f7d1cf78)    // logₑ(2)
const invLog2E = ln2                   // log₂(2)/log₂(e) = logₑ(2) => 1/log₂(e) = logₑ(2)
const log2E = int64(0x171547652b82fe0) // log₂(e)
const invLn2 = log2E                   // logₑ(e)/logₑ(2) = log₂(e) => 1/logₑ(2) = log₂(e)

func fixed56(v int64) int64 {
	return v << fracBits
}

func mul56(x, y int64) int64 {
	hi2, lo2 := bits.Mul64(uint64(x), uint64(y))
	lo, c := bits.Add64(lo2, uint64(halfValue), 0)
	hi := hi2 - uint64((x>>63)&y) - uint64((y>>63)&x) + c
	if hi>>63 != hi>>(fracBits-1)&1 {
		panic(ErrOverflow)
	}
	return int64(hi<<(64-fracBits) | (lo >> fracBits))
}

func mul56u(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	return int64(hi<<(64-fracBits) | (lo >> fracBits))
}

func div56(x, y int64) int64 {
	xs := x >> 63
	ys := y >> 63
	a := uint64((x ^ xs) - xs) // abs
	b := uint64((y ^ ys) - ys) // abs
	hi, lo := a>>(64-fracBits), a<<fracBits
	// will panic when divides by zero or occurs overflow
	v, rem := bits.Div64(hi, lo, b)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ {0,1}
	// round to near
	v, carry := bits.Add64(v, (rem<<1)/b, 0)
	if carry != 0 {
		panic(ErrOverflow)
	}
	return int64(v) * ((xs^ys)*2 + 1)
}

func floor56(x int64) int64 {
	return x & ^fracMask
}

func float56(x int64) float64 {
	xs := x >> 63
	a := uint64((x ^ xs) - xs)
	l := bits.Len64(a)
	if a != 0 {
		if l > 52 { //float64 significat bits count
			a = a >> (l - 52 - 1)
			l = l - fracBits - 1
		} else {
			a = a << (52 - l + 1)
			l -= fracBits + 1
		}
		a &= ^(uint64(1) << 52)
		a |= uint64(1023+l) << 52
		a |= uint64(xs) & (uint64(1) << 63) // sign
	}
	f := math.Float64frombits(a)
	return f
}
