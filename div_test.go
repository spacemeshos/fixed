package fixed

import (
	"fmt"
	"testing"
)

func Fixed_Div_(t *testing.T, div func(a, b Fixed) Fixed) {
	for i, tc := range divTestCases {
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no string representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else {
			if s == "overflow" {
				func() {
					defer func() { recover() }()
					x, y := New(tc.x), New(tc.y)
					got := x.Div(y)
					t.Errorf("Div: got %q, want overflow", got)
				}()
			} else {
				x, y := New(tc.x), New(tc.y)
				got := x.Div(y).String()
				if got != s {
					t.Errorf("x.Div(y): %s, want: %s, %v/%v", got, s, tc.x, tc.y)
				}
				got = Div64(int64(tc.x), int64(tc.y)).String()
				if got != s {
					t.Errorf("fixed.Div64(p,q): %s, want: %s, %v/%v", got, s, tc.x, tc.y)
				}
			}
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
