package fixed

import (
	"math/bits"
)

const log2_E = int64(1442695) * (1 << fracBits) / int64(1000000) // log2(E)
const log_2 = int64(693147) * (1 << fracBits) / int64(1000000)   // log(2)

/*
Log returns the natural logarithm of x
*/
func Log(x Fixed) Fixed {
	if x.int64 <= 0 {
		panic(ErrOverflow)
	}
	if x.int64 == oneValue {
		return Fixed{0}
	}
	// log(x) = lg2(x)/lg2(E)
	return lg2(x.int64).Div(Fixed{log2_E})
}

func lg2(x int64) Fixed {
	// required x > 0
	// 1 < fixed(x>>N) < 2
	// log2(x) = log2(2^N * x>>N) = N + log2(x>>N) = N + (log(x>>N)/log(2))
	N := bits.Len64(uint64(x)) - (fracBits)
	if N < 0 {
		N = 0
	}
	x >>= N
	r := Fixed{logApprox[x]}
	// N + (log(x>>n)/log(2))
	return r.Div(Fixed{log_2}).Add(Fixed{int64(N) << fracBits})
}
