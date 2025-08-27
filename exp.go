package fixed

func exp(x Fixed) Fixed {
	if x.integer() < -38 {
		return fixed(0)
	}
	return exp56(x.fixed56())
}

func exp56(x int64) Fixed {
	ey, k := exp_(x)
	z := rawfixed(ey)

	if x < 0 {
		if z.iszero() {
			panic(ErrOverflow)
		}
		z = z.inv()
		k = -k
	}

	if k < 0 {
		z = z.shr(int(-k))
	} else {
		z = z.shl(int(k))
	}

	return z
}

func exp_(x int64) (ey, k int64) {
	// exp(x) = 1/exp(-x) when x < 0
	// k = floor(x/ln(2)) => e^x = e^(ln(2)*k + (x - ln(2)*k)) = 2^k * e^(x-ln(2)*k)
	// y = x-ln(2)*k => e^y = 1 + y + y^2/2 + y^3/6 + y^4/24 + y^5/120 + ...

	xs := x >> 63
	a := (x ^ xs) - xs // abs(x)
	t := mul56u(a, invLn2)
	k = floor56(t)
	u := mul56u(k, ln2)
	k >>= fracBits // truncate to integer
	y := a - u
	ey = oneValue + y
	py := y

	for _, j := range invX {
		py = mul56u(py, y)
		ey += mul56u(py, j)
	}

	return
}

var invX = []int64{
	(int64(1) << 56) / 2,           // 1/2!
	(int64(1) << 56) / 6,           // 1/3!
	(int64(1) << 56) / 24,          // 1/4!
	(int64(1) << 56) / 120,         // 1/5!
	(int64(1) << 56) / 720,         // 1/6!
	(int64(1) << 56) / 5040,        // 1/7!
	(int64(1) << 56) / 40320,       // 1/8!
	(int64(1) << 56) / 362880,      // 1/9!
	(int64(1) << 56) / 3628800,     // 1/10!
	(int64(1) << 56) / 39916800,    // 1/11!
	(int64(1) << 56) / 479001600,   // 1/12!
	(int64(1) << 56) / 6227020800,  // 1/13!
	(int64(1) << 56) / 87178291200, // 1/14!
	(int64(1) << 56) / 1307674368000, // 1/15!
	(int64(1) << 56) / 20922789888000, // 1/16!
	//(int64(1) << 56) / 355687428096000, // 1/17!
}
