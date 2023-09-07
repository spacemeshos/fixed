package fixed

// a,b  - integer values
// x - 8/56 fixed point value
func incomplete(a, b, x int64) Fixed {
	// Iₓ(a,b) = (xᵃ*(1-x)ᵇ)/(a*B(a,b)) * (1/(1+(d₁/(1+(d₂/(1+...))))))
	// (xᵃ*(1-x)ᵇ)/B(a,b) = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + a*log(x) + b*log(1-x))
	// d_{2m+1} = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))
	// d_{2m}   = m(b-m)x/((a+2m-1)(a+2m))

	if a > int64(1)<<30 || b > int64(1)<<30 {
		panic(ErrOverflow)
	}

	bt := fixed(0)

	if 0 < x && x < oneValue {
		bt = exp(addx(subx(lgamma(a+b), lgamma(a), lgamma(b)), alogx(x, a), alogx(oneValue-x, b)))
	} else if x < 0 || x > oneValue {
		panic(ErrOverflow)
	}

	bcfx := func() Fixed {
		if bt.iszero() {
			return bt
		}
		h := bcf(x, a, b)
		return div(mul(bt, h), fixed(a))
	}

	if x > div(fixed(a+1), fixed(a+b+2)).fixed56() {
		// symmetry transform
		// 1 - bt/b*bcf(1-x,b,a)
		x, a, b = oneValue-x, b, a
		return sub(fixedOne, bcfx())
	}
	return bcfx()
}

var bcfEpsilon = from(1e-14)

func bcf(x, a, b int64) Fixed {
	const iters = 300
	xx := rawfixed(x)

	nonzero := func(z Fixed) Fixed {
		if z.iszero() {
			return rawfixed(1)
		}
		return z
	}

	c := fixedOne
	// d = 1/(nonzero(1-x*(a+b)/(a+1)))
	d := nonzero(sub(fixedOne, div(mul(xx, fixed(a+b)), fixed(a+1)))).inv()

	h := d
	del := fixed(0)

	for m := int64(1); m < iters; m++ {
		// fm := fixed(m)
		// amm := fixed(a + m + m)

		// d_{2m} = n = m(b-m)x/((a+2m-1)(a+2m))
		// n := div(mulx(fm, fixed(b-m), xx), mul(fixed(a+m+m-1), amm))
		n := div(mul(xx, fixed(m*(b-m))), fixed((a+m+m-1)*(a+m+m)))

		// d = 1/(nonzero(1+n*d))
		d = nonzero(muladd1(n, d)).inv()
		// c = nonzero(1 + n/c)
		c = nonzero(divadd1(n, c))
		// h = h*d*c
		h = mulx(h, d, c)

		// d_{2m+1} = n = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))
		// n = div(mulx(fixed(-a-m), fixed(a+b+m), xx), mul(amm, fixed(a+m+m+1)))
		n = div(mul(fixed((-a-m)*(a+b+m)), xx), fixed((a+m+m)*(a+m+m+1)))
		// d = 1/(nonzero(1+n*d))
		d = nonzero(muladd1(n, d)).inv()
		// c = nonzero(1 + n/c)
		c = nonzero(divadd1(n, c))

		del = mul(d, c)
		// fmt.Println(del.Float())
		h = mul(h, del)

		if sub(del, fixedOne).less(bcfEpsilon) {
			return h
		}
	}
	// panic(ErrOverflow)
	return h
}
