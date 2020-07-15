package fixed

func BinCDF(n, p, x Fixed) Fixed {
	if x.int64 < 0 {
		return Fixed{}
	} else if x.int64 >= n.int64 {
		return One
	} else {
		k := Fixed{floor(x.int64)}
		return BetaReg(n.Sub(k), k.Add(One), One.Sub(p))
	}
}
