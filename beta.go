package fixed

// BetaReg calculates regularized incomplete beta function Iₓ(a,b)
func BetaReg(a, b, x Fixed) Fixed {
	return Fixed{incomplete(a.int64, b.int64, x.int64)}
}

func incomplete(a, b, x int64) int64 {
	// Iₓ(a,b) = (xᵃ*(1-x)ᵇ)/(a*B(a,b)) * (1/(1+(d₁/(1+(d₂/(1+...))))))
	// (xᵃ*(1-x)ᵇ)/B(a,b) = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + a*log(x) + b*log(1-x))
	// d_{2m+1} = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))
	// d_{2m}   = m(b-m)x/((a+2m-1)(a+2m))
	bt := int64(0)
	if 0 < x && x < oneValue {
		q := add128(
				sub128(lgamma128(a+b), lgamma128(a), lgamma128(b)),
				alogx128(x,a),
				alogx128(oneValue-x, b))
		// q obviously must be small enough
		bt = exp(q.fixed())
		//bt = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + alogx(x, a) + alogx(oneValue-x, b))
	} else if x < 0 || x > oneValue {
		panic(ErrOverflow)
	}

	if x >= div(a+oneValue, a+b+oneValue+oneValue) {
		// symmetry transform
		return oneValue - mul( bcf(oneValue-x, b, a), div(bt, b) )
		//return oneValue - mulDiv(bt, bcf(oneValue-x, b, a), b)
	}
	return mul( bcf(x, a, b), div(bt, a) )
	//return mulDiv(bt, bcf(x, a, b), a)
}

func bcf(x, a, b int64) int64 {
	const iters = 1300
	const epsilon = int64(1)

	nonzero := func(z int64) int64 {
		const minval = 1 //oneValue >> (totalBits-fracBits)
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

		// x <= oneValue, a + b < max fixed value
		// => d_{2m} and d_{2m+1} < max fixed value

		// d_{2m} = m(b-m)x/((a+2m-1)(a+2m)) = m/(a+2m) * (b-m)/(a+2m-1) * x
		n := mul(div(mul(m,x),a_m2), div(b-m,a_m2-oneValue))
		//n := mulDiv(mul(x, b-m), m, mul(a_m2-oneValue, a_m2))
		d = inv(nonzero(oneValue + mul(n, d)))
		c = nonzero(oneValue + div(n, c))
		h = mul(mul(h, d), c)

		// d_{2m+1} = -(a+m)(a+b+m)x/((a+2m)(a+2m+1)) = (a+m)/(a+2m) * (a+b+m)/(a+2m+1) * -x
		n = mul(div(mul(a+m,-x),a_m2), div(a+b+m,a_m2+oneValue))
		//n = mulDiv(mul(x, a+b+m), -a-m, mul(a_m2, a_m2+oneValue))
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
