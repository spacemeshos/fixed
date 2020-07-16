package fixed

import "math/bits"

func div(x, y int64) int64 {
	xs := x >> 63
	ys := y >> 63
	a := uint64((x ^ xs) - xs)              // abs
	b := uint64((y ^ ys) - ys)              // abs
	hi, lo := a>>(64-fracBits), a<<fracBits // а*frac
	// will panic when divides by zero or occurs overflow
	v, rem := bits.Div64(hi, lo, b)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ [0,1]
	// round to near
	v, carry := bits.Add64(v, (rem<<1)/b, 0)
	if carry != 0 {
		panic(ErrOverflow)
	}
	return int64(v) * ((xs^ys)*2 + 1)
}

// Div returns x/y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Div(y Fixed) Fixed {
	return Fixed{div(x.int64, y.int64)}
}

// UsafeDiv returns x/y in fixed-point arithmetic.
// Does not have overflow check (but bits.Div64 has it's own)
func (x Fixed) UnsafeDiv(y Fixed) Fixed {
	xs := x.int64 >> 63
	ys := y.int64 >> 63
	a := uint64((x.int64 ^ xs) - xs)        // abs
	b := uint64((y.int64 ^ ys) - ys)        // abs
	hi, lo := a>>(64-fracBits), a<<fracBits // а*frac
	// will panic when divides by zero or occurs overflow
	v, rem := bits.Div64(hi, lo, b)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ [0,1]
	// round to near
	v, _ = bits.Add64(v, (rem<<1)/b, 0)
	return Fixed{int64(v) * ((xs^ys)*2 + 1)}
}

func inv(x int64) int64 {
	xs := x >> 63
	b := uint64((x ^ xs) - xs)                                                // abs
	hi, lo := uint64(oneValue)>>(63-fracBits), uint64(oneValue)<<(fracBits+1) // 1*(frac+1)
	// will panic when divides by zero or occurs overflow
	v, _ := bits.Div64(hi, lo, b)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ [0,1]
	// round to near
	v = (v + 1) >> 1
	return int64(v) * (xs*2 + 1)
}

func (x Fixed) Inv() Fixed {
	return Fixed{inv(x.int64)}
}
