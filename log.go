package fixed

import (
	"math/bits"
)

func log2(x int64) int64 {
	var N int
	var a uint64
	var b int64
	// log2(x) = log2(a*(2^b)) = log2(a) + b
	if x < oneValue {
		if x <= 0 {
			// required x > 0
			panic(ErrOverflow)
		}
		N = fracBits - bits.Len64(uint64(x)) + 1
		a = uint64(x << N)
		b = -(int64(N) << fracBits)
	} else {
		if x == oneValue {
			return 0
		}
		N = bits.Len64(uint64(x)) - (fracBits + 1)
		a = uint64(x >> N)
		b = int64(N) << fracBits
	}
	// 1 <= a < 2, а = α^t, 0 <= t < 1 = > a*a = α^(t+t)
	for i := 0; i < fracBits; i++ { // && a != 0 will slowdown code
		hi, lo := bits.Mul64(a, a)
		a = (hi << (64 - fracBits)) | (lo >> fracBits)
		// c = (t+t).floor()&1 = (α^(t+t)>>1)&1 => 0 or 1
		c := int64(a >> (fracBits + 1))
		// c == 1 when t+t >= 1
		a = a >> c // a/2
		b += (oneHalf >> i) * c
	}
	return b
}

func Log2(x Fixed) Fixed {
	return Fixed{log2(x.int64)}
}

func log(x int64) int64 {
	return mul54(log2(x), 0x2c5c85fdf473de)
}

func Log(x Fixed) Fixed {
	return Fixed{log(x.int64)}
}
