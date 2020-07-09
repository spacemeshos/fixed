package fixed

const eValue = int64(2718281) * (1 << fracBits) / 1000000

func Exp(x Fixed) Fixed {
	return iexp(x).Mul(fexp(x))
}

func fexp(x Fixed) Fixed {
	return Fixed{expApprox[x.int64&fracMask]}
}

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
	for n > 0 {
		if n&1 == 1 {
			r = r.Mul(e)
		}
		e = e.Mul(e)
		n = n >> 1
	}
	return r
}