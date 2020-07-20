package fixed

import "math/bits"

const (
	oneValue56   int64  = int64(1) << 56
	roundValue56 uint64 = uint64(1) << 55
)

// the constants are fixed-point numbers with 56-bit frac part
const ln2 = int64(0xb17217f7d1cf78)    // logₑ(2)
const invLog2E = ln2                   // log₂(2)/log₂(e) = logₑ(2) => 1/log₂(e) = logₑ(2)
const log2E = int64(0x171547652b82fe0) // log₂(e)
const invLn2 = log2E                   // logₑ(e)/logₑ(2) = log₂(e) => 1/logₑ(2) = log₂(e)

// converts regular fixed to fixed with 56-bit fraction
func fixed56(x int64) int64 {
	return x << (56 - fracBits)
}

// multiplies two fixed, at least one of them must have 56-bit precision
func mul56(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	hi = hi - uint64((x>>63)&y) - uint64((y>>63)&x)
	lo, carry := bits.Add64(lo, roundValue56, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	if carry != 0 || hi>>63 != hi>>(56-1)&1 {
		panic(ErrOverflow)
	}
	return int64(hi<<8 | lo>>56)
}

// multiplies two fixed, at least one of them must have 56-bit precision
// does not have overflow checks and sign compensation
func mul56u(x, y int64) int64 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	return int64(hi<<8 | lo>>56)
}

// divides one fixed by another, at least one must have 56-bit precision
func div56(x, y int64) int64 {
	xs := x >> 63
	ys := y >> 63
	a := uint64((x ^ xs) - xs) // abs
	b := uint64((y ^ ys) - ys) // abs
	hi, lo := a>>8, a<<56
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

// inverts regular fixed with 56-precision
func inv56(x int64) int64 {
	xs := x >> 63
	b := uint64((x ^ xs) - xs) // abs
	v, _ := bits.Div64(uint64(1<<(57+fracBits-64)), 0, b)
	v = (v + 1) >> 1
	return int64(v) * (xs*2 + 1)
}

// floor 56-precision fixed
func floor56(x int64) int64 {
	return x & ^(int64(1)<<56 - 1)
}
