package fixed

import (
	"fmt"
	"testing"
)

var divTestCases = []struct {
	x, y int
	s    map[string]string
}{{
	x: 2,
	y: 3,
	s: map[string]string{
		"26_6":  "0+43/64",
		"40_24": "0+11184811/16777216",
		"52_12": "0+2731/4096",
	},
}, {
	x: 1,
	y: 3,
	s: map[string]string{
		"26_6":  "0+21/64",
		"40_24": "0+05592405/16777216",
		"52_12": "0+1365/4096",
	},
}, {
	x: 10,
	y: 7,
	s: map[string]string{
		"26_6":  "1+27/64",
		"40_24": "1+07190235/16777216",
		"52_12": "1+1755/4096",
	},
}, {
	x: 18,
	y: 7,
	s: map[string]string{
		"26_6":  "2+37/64",
		"40_24": "2+09586981/16777216",
		"52_12": "2+2341/4096",
	},
}, {
	x: (1 << (totalBits - fracBits - 1)) - 1,
	y: 31,
	s: map[string]string{
		"26_6":  "1082401+00/64",
		"40_24": "17734058512+08118008/16777216",
		"52_12": "72638703667266+0132/4096",
	},
}}

func Fixed_Div_(t *testing.T, div func(a, b Fixed) Fixed) {
	for i, tc := range divTestCases {
		x, y := New(tc.x), New(tc.y)
		z := x.Div(y)
		got := z.String()
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no string representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else if got != s {
			t.Errorf("got: %s, want: %s, %v/%v", got, s, tc.x, tc.y)
		}
	}
}

func TestFixed_UnsafeDiv(t *testing.T) {
	Fixed_Div_(t, func(a, b Fixed) Fixed { return a.UnsafeDiv(b) })
}

func TestFixed_Div(t *testing.T) {
	Fixed_Div_(t, func(a, b Fixed) Fixed { return a.Div(b) })
}

func (x Fixed) refDiv(y Fixed) Fixed {
	ret := x.int64 / y.int64 << fracBits                          // Calculate whole part.
	ret += x.int64 % y.int64 << fracBits / y.int64                // Add fractional part.
	ret += x.int64 % y.int64 << fracBits % y.int64 << 1 / y.int64 // Round to nearest, instead of rounding down.
	return Fixed{ret}
}

func BenchmarkFixed_DivRef(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.refDiv(y)
	}
	Result = r
}

func BenchmarkFixed_DivUnsafe(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.UnsafeDiv(y)
	}
	Result = r
}

func BenchmarkFixed_Div(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.Div(y)
	}
	Result = r
}
