package fixed

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
)

// ErrOverflow the error for overflow panic
var ErrOverflow = errors.New("overflow")

const (
	fracBits   int    = 48 // 16.48
	totalBits  int    = 64
	fracMask   int64  = (int64(1) << fracBits) - 1
	roundValue uint64 = uint64(1) << (fracBits - 1)
	oneValue   int64  = int64(1) << fracBits
	oneHalf    int64  = int64(1) << (fracBits - 1)

	testEpsilonBits int = 3
)

// 0.0126157505

// convert integer to fixed with fraction defined as fracBits
func fixed(i int) int64 {
	if abs(int64(i)) >= (int64(1)<<(63-fracBits))-1 {
		panic(ErrOverflow)
	}
	return int64(i) << fracBits
}

// convert fixed with fraction defined as fracBits to integer
func integer(x int64) int {
	return int(x >> fracBits)
}

// floor fixed with fraction defined as fracBits
func floor(x int64) int64 {
	return x &^ fracMask
}

// cail fixed with fraction defined as fracBits
func ceil(x int64) int64 {
	return (x + oneValue - 1) &^ fracMask
}

// round fixed to the near integer value
func round(x int64) int64 {
	return (x + int64(roundValue)) &^ fracMask
}

// convert float to fixed with fraction defined as fracBits
func from(f float64) int64 {
	if abs(int64(f)) >= (int64(1)<<(63-fracBits))-1 {
		panic(ErrOverflow)
	}
	a := math.Float64bits(f)
	if a != 0 {
		v := int64(a & ((uint64(1) << 52) - 1))
		s := (int64(a)>>63)*2 + 1
		e := int((a>>52)&((1<<11)-1)) - 1023 - 52 + fracBits
		v |= int64(1) << 52
		if e < 0 {
			v >>= -e
		} else {
			v <<= e
		}
		return s * v
	}
	return 0
}

// convert fixed with fraction defined as fracBits to float
func float(x int64) float64 {
	xs := x >> 63
	a := uint64((x ^ xs) - xs)
	l := bits.Len64(a)
	if a != 0 {
		if l > 52 { //float64 significat bits count
			a = a >> (l - 52 - 1)
			l = l - fracBits - 1
		} else {
			a = a << (52 - l + 1)
			l -= fracBits + 1
		}
		a &= ^(uint64(1) << 52)
		a |= uint64(1023+l) << 52
		a |= uint64(xs) & (uint64(1) << 63) // sign
	}
	f := math.Float64frombits(a)
	return f
}

var formatString = fmt.Sprintf("%%s%%d+%%0%dx/%d", fracBits/4, fracBits)

// returns a human-readable representation of a fixed-point number.
func format(x int64) string {
	xs := x >> 63
	v := (x ^ xs) - xs
	sign := "-"[:xs&1]
	return fmt.Sprintf(formatString, sign, v>>fracBits, v&fracMask)
}

// absolute value of fixed
func abs(x int64) int64 {
	xs := x >> 63
	return (x ^ xs) - xs
}

// sign of fixed
func sign(x int64) int {
	xs := x >> 63
	return int(xs)*2 + 1
}

// max value of two fixeds
func max(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
