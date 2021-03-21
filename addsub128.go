package fixed

import "math/bits"

func add(x, y Fixed) Fixed {
	xs, ys := x.hi&signMask, y.hi&signMask
	if xs != ys {
		if ys != 0 {
			// x + -y => x - y
			return usub(x, y.abs())
		}
		// -x + y => y - x
		return usub(y, x.abs())
	}
	return uadd(x.abs(), y.abs()).setSignAs(x)
}

func addx(x Fixed, y ...Fixed) Fixed {
	for _, a := range y {
		x = add(x, a)
	}
	return x
}

func sub(x, y Fixed) Fixed {
	return add(x, y.neg())
}

func subx(x Fixed, y ...Fixed) Fixed {
	for _, a := range y {
		x = sub(x, a)
	}
	return x
}

func uadd(x, y Fixed) Fixed {
	r, c := Fixed{}, uint64(0)
	r.lo, c = bits.Add64(x.lo, y.lo, c)
	r.hi, c = bits.Add64(x.hi, y.hi, c)
	if c|r.hi&signMask != 0 {
		panic(ErrOverflow)
	}
	return r
}

func usub(x, y Fixed) (r Fixed) {
	c := uint64(0)
	if x.less(y) {
		r.lo, c = bits.Sub64(y.lo, x.lo, c)
		r.hi, _ = bits.Sub64(y.hi, x.hi, c)
		r.hi |= signMask
		return
	}
	r.lo, c = bits.Sub64(x.lo, y.lo, c)
	r.hi, _ = bits.Sub64(x.hi, y.hi, c)
	return r
}
