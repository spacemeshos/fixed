package fixed

import (
	"math/bits"
)

const log2_E = int64(1442695) * (1 << fracBits) / int64(1000000) // log2(E)
const log_2 = int64(693147) * (1 << fracBits) / int64(1000000)   // log(2)

// Log calculates the natural logarithm of x
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

// lg2 calculates log2(x)
func lg2(x int64) Fixed {
	// required x > 0
	// 1 <= fixed(x>>N or x<<N) < 2
	// x >= 1 => log2(x) = log2(2^N * x>>N) = N + log2(x>>N) = N + (log(x>>N)/log(2))
	// x < 1 => log2(x) = log2(x<<N / 2^N) = log2(x<<N) - N = (log(x<<N)/log(s)) - N
	if x < oneValue { // x < 1, but x > 0 as required
		N := fracBits - bits.Len64(uint64(x)) + 1
		x <<= N // 1 <= x < 2
		r := Fixed{logApprox[x-oneValue]}
		// N + (log(x>>n)/log(2))
		return r.Div(Fixed{log_2}).Sub(Fixed{int64(N) << fracBits})
	} else { // x >= 1
		N := bits.Len64(uint64(x)) - (fracBits + 1)
		x >>= N // 1 <= x < 2
		r := Fixed{logApprox[x-oneValue]}
		// N + (log(x>>n)/log(2))
		return r.Div(Fixed{log_2}).Add(Fixed{int64(N) << fracBits})
	}
}
