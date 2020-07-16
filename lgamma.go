package fixed

func lgamma(x int64) int64 {
	t := log(x)
	if x <= 2*oneValue {
		return lgamma(x+oneValue) - t
	}
	const g0 = 0x1acfe390c97d69 // 0.4189385332046727
	const g1 = 0x555555555554e  // 0.08333333333333297
	const g2 = -0x2d82d82d605c  // -0.0027777777772877554
	const g3 = 0xd00cfcc679c    // 0.0007936505586430196
	const g4 = -0x9c065c607f3   // -0.00059518755745034
	const g5 = 0xdb3dd266d6a    // 0.0008363399189962821
	const g6 = -0x1ab89d0b9e43  // -0.0016309293409657527
	z := inv(x)
	y := mul(z, z)
	w := (g0 + mul(z, g1+mul(y, g2+mul(y, g3+mul(y, g4+mul(y, g5+mul(y, g6))))))) >> (54 - fracBits)
	return mul(x-oneHalf, t-oneValue) + w
}

func Lgamma(x Fixed) Fixed {
	return Fixed{lgamma(x.int64)}
}
