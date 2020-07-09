package fixed

const eValue = int64(2718281) * (1 << fracBits) / 1000000

// Exp calculate e^x
func Exp(x Fixed) Fixed {
	return iexp(x).Mul(fexp(x))
}

// fexp calculates e^x for fraction part
func fexp(x Fixed) Fixed {
	return Fixed{expApprox[x.int64&fracMask]}
}

// iexp calculates e^x for integer part
func iexp(x Fixed) Fixed {
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
			r = r.Mul(Fixed{epow[i]})
		}
		n = n >> 1
	}
	if n > 0 {
		panic(ErrOverflow)
	}
	return r
}
