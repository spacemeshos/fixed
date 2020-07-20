package fixed

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
	//(int64(1) << 56) / 6402373705728000, // 1/15!
}

// max possible power of E for defined fractional size
var maxEpow = floor(log(fixed(1 << (64 - fracBits - 2))))

func exp(x int64) int64 {
	// exp(x) = 1/exp(-x) when x < 0
	// k = floor(x/ln(2)) => e^x = e^(ln(2)*k + (x - ln(2)*k)) = 2^k * e^(x-ln(2)*k)
	// y = x-ln(2)*k => e^y = 1 + y + y^2/2 + y^3/6 + y^4/24 + y^5/120

	if y := floor(x); y > maxEpow {
		panic(ErrOverflow)
	} else if y < -maxEpow {
		return 0
	}

	xs := x >> 63
	a := fixed56((x ^ xs) - xs) // abs(x) -> 56-bit
	t := mul56u(a, invLn2)
	k := floor56(t)
	u := mul56u(k, ln2)
	k >>= 56 // truncate to integer
	y := a - u
	ey := (oneValue << (56 - fracBits)) + y
	py := y

	for _, j := range invX {
		py = mul56u(py, y)
		ey += mul56u(py, j)
	}

	if int(k) < 56-fracBits {
		ey >>= (56 - fracBits) - int(k)
	} else {
		ey <<= int(k) - (56 - fracBits)
	}

	if xs != 0 {
		return inv(ey)
	}

	return ey
}

// Exp calculates eˣ
func Exp(x Fixed) Fixed {
	return Fixed{exp(x.int64)}
}
