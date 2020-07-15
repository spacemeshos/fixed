package fixed

func BetaReg(a, b, x Fixed) Fixed {
	bt := int64(0)
	if x.int64 > 0 && x.int64 < oneValue {
		invBetaLn := GammaLn(Fixed{a.int64 + b.int64}).int64 -
			GammaLn(Fixed{a.int64}).int64 -
			GammaLn(Fixed{b.int64}).int64
		xab := Log(x).Mul(a).int64 + Log(Fixed{oneValue - x.int64}).Mul(b).int64
		bt = Exp(Fixed{invBetaLn + xab}).int64
	} else if x.int64 == 0 || x.int64 == oneValue {
		return x
	}

	symm := x.int64 >= div(a.int64+oneValue, a.int64+b.int64+2)
	if symm {
		a, b, x = b, a, Fixed{oneValue - x.int64}
	}

	eps := int64(1)
	fpmin := oneValue >> 3

	qab := a.int64 + b.int64
	qap := a.int64 + oneValue
	qam := a.int64 - oneValue
	c := oneValue
	d := oneValue - mulDiv(qab, x.int64, qap)

	if abs(d) < fpmin {
		d = fpmin
	}
	d = div(oneValue, d)
	h := d

	for m := oneValue; m < int64(31)<<fracBits; m += oneValue {
		m2 := m << 1
		aa := mulDiv(mul(m, (b.int64-m)), x.int64, mul(qam+m2, a.int64+m2))
		d = oneValue + mul(aa, d)
		if abs(d) < fpmin {
			d = fpmin
		}
		c = oneValue + div(aa, c)
		if abs(c) < fpmin {
			c = fpmin
		}
		d = div(oneValue, d)
		h = mul(mul(h, d), c)
		aa = mulDiv(mul(-a.int64-m, qab+m), x.int64, mul(a.int64+m2, qap+m2))
		d = oneValue + mul(aa, d)
		if abs(d) < fpmin {
			d = fpmin
		}
		c = oneValue + div(aa, c)
		if abs(c) < fpmin {
			c = fpmin
		}
		d = div(oneValue, d)
		del := mul(d, c)
		h = mul(h, del)
		if e := abs(del - oneValue); e <= eps {
			break
		}
	}

	if symm {
		return Fixed{oneValue - mulDiv(bt, h, a.int64)}
	}
	return Fixed{mulDiv(bt, h, a.int64)}
}
