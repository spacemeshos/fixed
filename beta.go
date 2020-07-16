package fixed

func BetaReg(a, b, x Fixed) Fixed {
	return Fixed{incomplete(a.int64, b.int64, x.int64)}
}

func incomplete(a, b, x int64) int64 {
	// I(x; a,b) = (xᵃ*(1-x)ᵇ)/(a*B(a,b)) * (1/(1+(d₁/(1+(d₂/(1+...))))))
	// (xᵃ*(1-x)ᵇ)/B(a,b) = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + a*log(x) + b*log(1-x))
	// d_{2m+1} = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))
	// d_{2m}   = m(b-m)x/((a+2m-1)(a+2m))
	bt := int64(0)
	if 0 < x && x < oneValue {
		bt = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + mul(log(x), a) + mul(log(oneValue-x), b))
	} else if x < 0 || x > oneValue {
		panic(ErrOverflow)
	}

	if x >= div(a+oneValue, a+b+oneValue+oneValue) {
		// continued fraction after symmetry transform.
		return oneValue - mulDiv(bt, bcf(oneValue-x, b, a), b)
	}
	return mulDiv(bt, bcf(x, a, b), a)
}

func bcf(x, a, b int64) int64 {
	const iters = 31
	const epsilon = int64(1)

	nonzero := func(z int64) int64 {
		const minval = oneValue >> 3
		if abs(z) < minval {
			return minval
		}
		return z
	}

	c := oneValue
	d := inv(nonzero(oneValue - mulDiv(a+b, x, a+oneValue)))
	h := d
	for m := oneValue; m < fixed(iters); m += oneValue {
		a_m2 := a + m + m

		// d_{2m}
		n := mulDiv(mul(m, b-m), x, mul(a_m2-oneValue, a_m2))
		d = inv(nonzero(oneValue + mul(n, d)))
		c = nonzero(oneValue + div(n, c))
		h = mul(mul(h, d), c)

		// d_{2m+1}
		n = mulDiv(mul(-a-m, a+b+m), x, mul(a_m2, a_m2+oneValue))
		d = inv(nonzero(oneValue + mul(n, d)))
		c = nonzero(oneValue + div(n, c))

		del := mul(d, c)
		h = mul(h, del)
		if abs(del-oneValue) <= epsilon {
			return h
		}
	}
	panic(ErrOverflow)
}
