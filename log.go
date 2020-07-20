package fixed

import (
	"math/bits"
)

func log2(x int64) int64 {
	var N int
	var a uint64
	var b int64
	// log₂(x) = log₂(a*(2ᵇ)) = log₂(a) + b
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
	for i := 0; i < fracBits; i++ {
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

// Log2 returns log₂(x)
func Log2(x Fixed) Fixed {
	return Fixed{log2(x.int64)}
}

// returns regular fixed from regular fixed argument
func log(x int64) int64 {
	return mul56(log2(x), invLog2E)
}

// returns 56-bit precision fixed from regular fixed argument
func log56(x int64) int64 {
	return mul(log2(x), invLog2E)
}

// regular a * log(x) with internal 56-bit precision multiplication
func alogx(x int64, a int64) int64 {
	return mul56(mul(log2(x), invLog2E), a)
}

// Log returns Ln(x)
func Log(x Fixed) Fixed {
	return Fixed{log(x.int64)}
}
