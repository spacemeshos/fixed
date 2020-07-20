package fixed

import (
	"fmt"
	"math"
	"testing"
)

func TestFixedMul(t *testing.T) {
	for i, tc := range mulTestCases {
		x := From(tc.x)
		y := From(tc.y)
		want, found := tc.z[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no float representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else if z := float64(x.int64) * float64(y.int64) / (1 << (fracBits * 2)); z != want {
			t.Errorf("tc.x=%v, tc.y=%v: z: got %v, want %v", tc.x, tc.y, z, want)
			continue
		}
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no string representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else {
			if s == "overflow" {
				func() {
					defer func() { recover() }()
					got := x.Mul(y)
					t.Errorf("tc.x=%v: Mul: got %q, want overflow", tc.x, got)
				}()
			} else if got, want := x.Mul(y).String(), s; got != want {
				t.Errorf("tc.x=%v: Mul: got %q, want %q", tc.x, got, want)
			}
		}
	}
}

func TestFixedMulByOneMinusIota(t *testing.T) {
	oneMinusIota := I(1<<fracBits - 1)
	oneMinusIotaF := float64(oneMinusIota.int64) / (1 << fracBits)

	for _, neg := range []bool{false, true} {
		for i := 0; i < totalBits; i++ {
			x := I(1 << i)
			if neg {
				x.int64 = -x.int64
			} else if i == totalBits-1 {
				// A signed int32 can't represent 1<<31.
				// A signed int64 can't represent 1<<63.
				continue
			}

			// want equals x * oneMinusIota, rounded to nearest.
			want := Fixed{}
			if -1<<fracBits < x.int64 && x.int64 < 1<<fracBits {
				// (x * oneMinusIota) isn't exactly representable as an
				// Fixed. Calculate the rounded value using float64 math.
				xF := float64(x.int64) / (1 << fracBits)
				wantF := xF * oneMinusIotaF * (1 << fracBits)
				want.int64 = int64(math.Floor(wantF + 0.5))
			} else {
				// (x * oneMinusIota) is exactly representable.
				want.int64 = oneMinusIota.int64 << (i - fracBits)
				if neg {
					want.int64 = -want.int64
				}
			}

			if got := x.Mul(oneMinusIota); got != want {
				t.Errorf("neg=%t, i=%d, x=%v, Mul: got %v, want %v", neg, i, x, got, want)
			}
		}
	}
}

func (x Fixed) refMul(y Fixed) Fixed {
	xw, yw := x.int64>>fracBits, y.int64>>fracBits // Whole part of x and y.
	xf, yf := x.int64&fracMask, y.int64&fracMask   // Fractional part of x and y.
	ret := (xw * yw) << fracBits                   // Multiply whole part.
	ret += xw*yf + yw*xf                           // Multiply each whole by other fraction.
	ret += (xf * yf) >> fracBits                   // Multiply fractions by each other.
	ret += (xf * yf) >> (fracBits - 1) & 1         // Round to nearest, instead of rounding down.
	return Fixed{ret}
}

func BenchmarkFixed_MulRef(b *testing.B) {
	var r Fixed
	for i := 0; i < b.N; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.refMul(y)
	}
	Result = r
}

func BenchmarkFixed_Mul(b *testing.B) {
	var r Fixed
	for i := 0; i < b.N; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.Mul(y)
	}
	Result = r
}

func BenchmarkFixed_MulUnsafe(b *testing.B) {
	var r Fixed
	for i := 0; i < b.N; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.UnsafeMul(y)
	}
	Result = r
}
