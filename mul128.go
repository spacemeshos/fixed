package fixed

import "math/bits"

func mul(x, y Fixed) Fixed {
	a, b := x.hi&unsignMask, y.hi&unsignMask
	r := [4]uint64{}

	c := uint64(0)
	r[1], r[0] = bits.Mul64(x.lo, y.lo)

	h, l := bits.Mul64(a, y.lo)
	r[1], c = bits.Add64(r[1], l, 0)
	r[2], c = bits.Add64(r[2], h, c)
	r[3] += c

	h, l = bits.Mul64(x.lo, b)
	r[1], c = bits.Add64(r[1], l, 0)
	r[2], c = bits.Add64(r[2], h, c)
	r[3] += c

	h, l = bits.Mul64(a, b)
	r[2], c = bits.Add64(r[2], l, 0)
	r[3], c = bits.Add64(r[3], h, c)

	r[0] = r[0]>>56 | r[1]<<8
	r[1] = r[1]>>56 | r[2]<<8
	r[2] = r[2]>>56 | r[3]<<8
	r[3] = r[3]>>56 | 0

	if c|r[3]|r[2]|(r[1]&signMask) != 0 {
		panic(ErrOverflow)
	}

	return Fixed{lo: r[0], hi: r[1] | (x.hi&signMask ^ y.hi&signMask)}
}

func mul64_(x Fixed, y uint64) Fixed {
	b, a := bits.Mul64(x.lo, y)
	c0, d := bits.Mul64(x.hi, y)
	b, c1 := bits.Add64(b, d, 0)
	if c0|c1 != 0 {
		panic(ErrOverflow)
	}
	return Fixed{lo: a, hi: b}
}

func muladd1(x, y Fixed) Fixed {
	return add(mul(x, y), fixedOne)
}

func mulx(x Fixed, y ...Fixed) Fixed {
	for _, a := range y {
		x = mul(x, a)
	}
	return x
}
