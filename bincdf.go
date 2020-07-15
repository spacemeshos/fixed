package fixed

func BinCDF(n, p, x Fixed) Fixed {
	if x.int64 < 0 {
		return Fixed{}
	} else if x.int64 >= n.int64 {
		return One
	} else {
		k := floor(x.int64)
		return Fixed{incomplete(n.int64-k, k+oneValue, oneValue-p.int64)}
	}
}
