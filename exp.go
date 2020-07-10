package fixed

const eValue = int64(2718281) * (1 << fracBits) / 1000000

// Exp calculates e^x
func Exp(x Fixed) Fixed {
	if x.int64 < 0 {
		// a = -x, e^a = 2^(log2(E)*a) = 2^(-i+f) = 2^f>>i
		n := Fixed{log2_E}.Mul(x.Neg()).int64
		f := oneValue - n&fracMask
		n += oneValue
		return Fixed{int64(pow2Approx[f]) >> (n >> fracBits)}
	}
	// e^x = e^(i+f) = e^i*e^f => iexp(x)*fexp(x)
	return iexp(x).Mul(Fixed{int64(expApprox[x.int64&fracMask])})
}

// iexp calculates e^x for a positive integer part
func iexp(x Fixed) Fixed {
	if x.int64 > 144796 {
		panic(ErrOverflow)
	}
	n := x.int64 >> fracBits
	r := Fixed{oneValue}
	if n == 0 {
		return r
	}
	e := Fixed{eValue}
	if n == 1 {
		return e
	}
	for i := 1; n > 0 && i < len(epow); i++ {
		if n&1 == 1 {
			r = r.UnsafeMul(Fixed{epow[i]})
		}
		n = n >> 1
	}
	return r
}
