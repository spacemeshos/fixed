package fixed

func lgamma(x int64) int64 {
	switch {
	case x <= 2*oneValue:
		switch {
		case x <= 0:
			panic(ErrOverflow)
		case x == oneValue, x == 2*oneValue:
			return 0
		default:
			return lgamma(x+oneValue) - log(x)
		}
	case x < 8*oneValue:
		const gR1 = int64(0x1645a762c4ab740) // 1.3920053346762105
		const gR2 = int64(0xb8d0c49e9ee6e0)  // 0.7219355475671381
		const gR3 = int64(0x2c03db99f7be4e)  // 0.17193386563280308
		const gR4 = int64(0x4c5fa9d0bb51d)   // 0.01864591917156529
		const gR5 = int64(0x32fbb5948352)    // 0.0007779424963818936
		const gR6 = int64(0x7aebde96ce)      // 7.326684307446256e-06
		const gS0 = int64(-0x13c467e37db0c8) // -0.07721566490153287
		const gS1 = int64(0x3709166dc410f0)  // 0.21498241596060885
		const gS2 = int64(0x53663d3c4e7d64)  // 0.325778796408931
		const gS3 = int64(0x2577397dcbe5ee)  // 0.14635047265246445
		const gS4 = int64(0x6d2071fa4e658)   // 0.02664227030336386
		const gS5 = int64(0x789ad9cda3c8)    // 0.0018402845140733772
		const gS6 = int64(0x217fd9ba2fd)     // 3.194753265841009e-05

		i := floor(x)
		y := fixed56(x - i)
		p := mul56(y, gS0+mul56(y, gS1+mul56(y, gS2+mul56(y, gS3+mul56(y, gS4+mul56(y, gS5+mul56(y, gS6)))))))
		q := oneValue56 + mul56(y, gR1+mul56(y, gR2+mul56(y, gR3+mul56(y, gR4+mul56(y, gR5+mul56(y, gR6))))))
		g := y/2 + div56(p, q)
		z := oneValue
		i -= oneValue
		for ; i >= 2*oneValue; i -= oneValue {
			z = mul(z, y>>(56-fracBits)+i)
		}
		g += log56(z)
		return g >> (56 - fracBits)

	default:
		const gW0 = int64(0x6b3f8e4325f5a4) // 0.4189385332046727
		const gW1 = int64(0x1555555555553b) // 0.08333333333333297
		const gW2 = int64(-0xb60b60b58172)  // -0.0027777777772877554
		const gW3 = int64(0x34033f319e71)   // 0.0007936505586430196
		const gW4 = int64(-0x270197181fce)  // -0.00059518755745034
		const gW5 = int64(0x36cf7499b5ab)   // 0.0008363399189962821
		const gW6 = int64(-0x6ae2742e790f)  // -0.0016309293409657527

		// calculate with 56 bit precision independent of fraction part size
		t := log56(x) - oneValue56
		z := inv56(x)
		y := mul56(z, z)
		w := gW0 + mul56(z, gW1+mul56(y, gW2+mul56(y, gW3+mul56(y, gW4+mul56(y, gW5+mul56(y, gW6))))))

		// normalize w to fit fixed fraction part
		w >>= 56 - fracBits
		return mul56(x-oneHalf, t) + w
	}
}

// Lgamma calculates logₑΓ(x)
//
// defined only for x ∈ (0,2^(64-fracBits-1)]
func Lgamma(x Fixed) Fixed {
	return Fixed{lgamma(x.int64)}
}
