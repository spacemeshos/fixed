package fixed

import (
	"math/bits"
)

func alogx(x int64, a int64) Fixed {
	return mul(log(x), fixed(a))
}

func log(x int64) Fixed {
	return rawfixed(log56(x))
}

func ilog(x int64) Fixed {
	return rawfixed(ilog56(x))
}

func log56(x int64) int64 {
	return mul56(log2(x, fracBits), invLog2E)
}

func ilog56(x int64) int64 {
	return mul56(log2(x, 0), invLog2E)
}

func log2(x int64, onew int) int64 {
	var N int
	var a uint64
	var b int64
	// log₂(x) = log₂(a*(2ᵇ)) = log₂(a) + b
	one := int64(1) << onew
	if x < one {
		if x <= 0 {
			// required x > 0
			panic(ErrOverflow)
		}
		N = (onew + 1) - bits.Len64(uint64(x))
		b = -(int64(N) << fracBits)
	} else {
		if x == one {
			return 0
		}
		N = bits.Len64(uint64(x)) - (onew + 1)
		b = int64(N) << fracBits
	}
	a = (uint64(x) << (64 - bits.Len64(uint64(x)))) >> 1
	// 1 <= a < 2, а = α^t, 0 <= t < 1 = > a*a = α^(t+t)
	for i := 0; i < fracBits; i++ {
		hi, lo := bits.Mul64(a, a)
		a = (hi << 2) | (lo >> 62)
		// c = (t+t).floor()&1 = (α^(t+t)>>1)&1 => 0 or 1
		c := int64(a >> 63)
		// a = a/2 if t+t >= 1
		a = a >> c
		// set fraction bit i to 1 if t+t >= 1
		b += (halfValue >> i) * c // TODO: multiplication can be replaced by shifting 'c'
	}
	return b
}
