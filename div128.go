package fixed

import "math/bits"

//go:noinline
func div64(u Fixed, v uint64) (Fixed, Fixed) {
	if u.hi < v {
		lo, r := bits.Div64(u.hi, u.lo, v)
		return Fixed{lo: lo}, Fixed{lo: r}
	} else {
		hi, r := bits.Div64(0, u.hi, v)
		lo, r := bits.Div64(r, u.lo, v)
		return Fixed{lo: lo, hi: hi}, Fixed{lo: r}
	}
}

func div128(u, v Fixed) (Fixed, Fixed) {
	if v.hi == 0 {
		return div64(u, v.lo)
	}

	n := bits.LeadingZeros64(v.hi)
	v1 := Fixed{lo: v.lo << n, hi: v.hi<<n | v.lo>>(64-n)}

	tq, _ := bits.Div64(u.hi, u.lo, v1.hi)
	tq >>= 64 - n
	tq -= -tq >> 63 //if tq != 0 { tq-- }

	q := Fixed{lo: tq}
	r := sub(u, mul64_(v, tq))

	if r.greater(v) {
		lo, c := bits.Add64(q.lo, 1, 0)
		q = Fixed{lo: lo, hi: c}
		r = sub(r, v)
	}

	return q, r
}

func udiv(u, v Fixed) Fixed {
	if u.hi & ^uint64(0x07f) == 0 {
		r, _ := div128(Fixed{lo: u.lo << 56, hi: u.hi<<56 | u.lo>>8}, v)
		return r
	}

	a, rem := div128(u, v)
	a = a.shl(56)
	b, n := rem.shlmax(56)
	v = v.shr(56 - n)
	b, _ = div128(b, v)
	a = add(a, b)
	return a
}

func div(x, y Fixed) Fixed {
	u, v := x.abs(), y.abs()
	a := udiv(u, v)
	a.hi |= x.hi&signMask ^ y.hi&signMask
	return a
}

func divadd1(x, y Fixed) Fixed {
	return add(div(x, y), fixedOne)
}

func (x Fixed) inv() Fixed {
	return div(fixedOne, x)
}
