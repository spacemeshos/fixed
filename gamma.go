package fixed

var gammaDK = []int64{
	0x00000068427088c7, // x >> 54-fracBits => 2.4857408913875355e-05
	0x00434a86fd9492ec, // x >> 54-fracBits => 1.0514237858172197
	-0x0dd3d5fbf0b5940, // x >> 54-fracBits => -3.4568709722201625
	0x0120c925de05f430, // x >> 54-fracBits => 4.512277094668948
	-0x0bee70d231ca448, // x >> 54-fracBits => -2.9828522532357664
	0x00439c02a5f4e264, // x >> 54-fracBits => 1.056397115771267
	-0x00c81e7af638da6, // x >> 54-fracBits => -0.19542877319164587
	0x0001181e3e500255, // x >> 54-fracBits => 0.01709705434044412
	-0x000095ed4fe1425, // x >> 54-fracBits => -0.0005719261174043057
	0x000000136fb6c5ed, // x >> 54-fracBits => 4.633994733599057e-06
	-0x000000002eba766, // x >> 54-fracBits => -2.7199490848860772e-09
}

const ln2SqrtEOverPi int64 = 0x9eeb95 // 0.6207822
const gammaR int64 = 0xae687d2        // 10.90051
const invE = 0x178b56362cef38         // x >> 54-fracBits => 1/E

func GammaLn(x Fixed) Fixed {
	if x.int64 < (oneValue >> 1) {
		panic(ErrOverflow)
	}

	s := gammaDK[0]
	for i, g := range gammaDK[1:] {
		s = s + div(g, x.int64+(int64(i)<<fracBits))
	}
	s >>= 54 - fracBits

	xh := x.int64 - (oneValue >> 1)
	lgGr := mul(log(mul54(xh+gammaR, invE)), xh)
	return Fixed{log(s) + ln2SqrtEOverPi + lgGr}
}
