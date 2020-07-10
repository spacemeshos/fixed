package fixed

import (
	"math/bits"
)

const log2_E = int64(1442695) * (1 << fracBits) / int64(1000000) // log2(E)
const log_2 = int64(693147) * (1 << fracBits) / int64(1000000)   // log(2)
var invLog_2 = Fixed{oneValue}.Div(Fixed{log_2})
var invLog2_E = Fixed{oneValue}.Div(Fixed{log2_E})

// Log calculates the natural logarithm of x
func Log(x Fixed) Fixed {
	// required x > 0
	if x.int64 <= 0 {
		panic(ErrOverflow)
	}
	var lg2 Fixed
	if x.int64 < oneValue { // 0 < x < 1
		// log2(x) = log2(x<<N / 2^N) = log2(x<<N) - N = (log(x<<N)/log(2)) - N
		N := fracBits - bits.Len64(uint64(x.int64)) + 1
		a := x.int64 << N // 1 <= a < 2
		// (log(a<<n)/log(2)) - N
		lg2 = Fixed{int64(logApprox[a-oneValue])}.
			Mul(invLog_2). // instead Div(Fixed{log_2})
			Sub(Fixed{int64(N) << fracBits})
	} else { // x >= 1
		// log2(x) = log2(2^N * x>>N) = N + log2(x>>N) = N + (log(x>>N)/log(2))
		N := bits.Len64(uint64(x.int64)) - (fracBits + 1)
		a := x.int64 >> N // 1 <= a < 2
		// N + (log(a>>n)/log(2))
		lg2 = Fixed{int64(logApprox[a-oneValue])}.
			Mul(invLog_2). // instead Div(Fixed{log_2})
			Add(Fixed{int64(N) << fracBits})
	}
	// log(x) = log2(x)/log2(E)
	return lg2.Mul(invLog2_E) // instead lg2.Div(Fixed{log2_E})
}
