package fixed

// x 64-bit is an integer value
func lgamma(x int64) Fixed {
	switch {
	case x <= 0:
		panic(ErrOverflow)

	case x <= 2:
		return fixed(0)

	case x < 8:
		z := int64(1)
		x -= 1
		for ; x >= 2; x -= 1 {
			z *= x
		}
		return rawfixed(ilog56(z))

	default:
		const gW0 = int64(0x6b3f8e4325f5a4) // 0.4189385332046727
		const gW1 = int64(0x1555555555553b) // 0.08333333333333297
		const gW2 = int64(-0xb60b60b58172)  // -0.0027777777772877554
		const gW3 = int64(0x34033f319e71)   // 0.0007936505586430196
		const gW4 = int64(-0x270197181fce)  // -0.00059518755745034
		const gW5 = int64(0x36cf7499b5ab)   // 0.0008363399189962821
		const gW6 = int64(-0x6ae2742e790f)  // -0.0016309293409657527

		t := ilog56(x) - oneValue
		z := div56(1, x) // yes, it uses 1 instead oneValue because x is an integer value here
		y := mul56(z, z)
		w := gW0 + mul56(z, gW1+mul56(y, gW2+mul56(y, gW3+mul56(y, gW4+mul56(y, gW5+mul56(y, gW6))))))
		a := sub(fixed(x), rawfixed(halfValue))
		return add(mul(a, rawfixed(t)), rawfixed(w))
	}
}
