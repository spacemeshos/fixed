package fixed

const eValue = int64(2718281) * (1 << fracBits) / 1000000

// Exp calculates e^x
func Exp(x Fixed) Fixed {
	// e^x = e^(i+f) = e^i*e^f => iexp(x)*fexp(x)
	return iexp(x).Mul(fexp(x))
}

// fexp calculates e^x for fraction part
func fexp(x Fixed) Fixed {
	return Fixed{expApprox[x.int64&fracMask]}
}

// iexp calculates e^x for integer part
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

// FastExp calculates e^x fast but with lower precision
//func FastExp(x Fixed) Fixed {
//	// e^x = 2^(log2(e)*x)
//	// log2(e)*x = i + f
//	// e^x = 2^(i+f) = 2^f<<i
//	n := Fixed{log2_E}.Mul(x)
//	i := 62 - fracBits - int(n.int64>>fracBits)
//	if i < 0 {
//		panic(ErrOverflow)
//	}
//	return Fixed{pow2[n.int64&fracMask] >> i }
//}
