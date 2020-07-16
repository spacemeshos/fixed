package fixed

var invX = []int64{
	(int64(1) << 54) / 2,    // 1/2!
	(int64(1) << 54) / 6,    // 1/3!
	(int64(1) << 54) / 24,   // 1/4!
	(int64(1) << 54) / 120,  // 1/5!
	(int64(1) << 54) / 720,  // 1/6!
	(int64(1) << 54) / 5040, // 1/7!
	//	(int64(1)<<54)/40320, // 1/8!
	//	(int64(1)<<54)/362880, // 1/9!
	//	(int64(1)<<54)/3628800, // 1/10!
	//	(int64(1)<<54)/39916800, // 1/11!
	//	(int64(1)<<54)/479001600, // 1/12!
}

func exp(x int64) int64 {
	// exp(x) = 1/exp(-x) when x < 0
	// k = floor(x/ln(2)) => e^x = e^(ln(2)*k + (x - ln(2)*k)) = 2^k * e^(x-ln(2)*k)
	// y = x-ln(2)*k => e^y = 1 + y + y^2/2 + y^3/6 + y^4/24 + y^5/120

	if y := floor(x); y >= fixed(27) {
		panic(ErrOverflow)
	} else if y <= fixed(-27) {
		return 0
	}

	xs := x >> 63
	a := (x ^ xs) - xs // abs(x)
	a <<= (54 - fracBits)
	invLog2 := int64(0x5c551d94ae0bf8)
	t := mul54u(a, invLog2)
	// k = floor(t) with 54-bit frac
	k := t &^ (int64(1)<<54 - 1)
	log2 := int64(0x2c5c85fdf473de)
	u := mul54u(k, log2)
	k >>= 54
	y := a - u
	ey := (oneValue << (54 - fracBits)) + y
	py := y

	for _, j := range invX {
		py = mul54u(py, y)
		ey += mul54u(py, j)
	}

	if int(k) < 54-fracBits {
		ey >>= (54 - fracBits) - int(k)
	} else {
		ey <<= int(k) - (54 - fracBits)
	}

	if xs != 0 {
		return inv(ey)
	}

	return ey
}

// Exp calculates e^x
func Exp(x Fixed) Fixed {
	return Fixed{exp(x.int64)}
}
