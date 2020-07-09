// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fixed

import (
	"fmt"
	"math"
	"testing"
)

var testCases = []struct {
	x     float64
	s     map[string]string
	floor int
	round int
	ceil  int
}{{
	x: 0,
	s: map[string]string{
		"26_6":  "0+00/64",
		"40_24": "0+00000000/16777216",
		"52_12": "0+0000/4096",
	},
	floor: 0,
	round: 0,
	ceil:  0,
}, {
	x: 1,
	s: map[string]string{
		"26_6":  "1+00/64",
		"40_24": "1+00000000/16777216",
		"52_12": "1+0000/4096",
	},
	floor: 1,
	round: 1,
	ceil:  1,
}, {
	x: 1.25,
	s: map[string]string{
		"26_6":  "1+16/64",
		"40_24": "1+04194304/16777216",
		"52_12": "1+1024/4096",
	},
	floor: 1,
	round: 1,
	ceil:  2,
}, {
	x: 2.5,
	s: map[string]string{
		"26_6":  "2+32/64",
		"40_24": "2+08388608/16777216",
		"52_12": "2+2048/4096",
	},
	floor: 2,
	round: 3,
	ceil:  3,
}, {
	x: 63 / 64.0,
	s: map[string]string{
		"26_6":  "0+63/64",
		"40_24": "0+16515072/16777216",
		"52_12": "0+4032/4096",
	},
	floor: 0,
	round: 1,
	ceil:  1,
}, {
	x: -0.5,
	s: map[string]string{
		"26_6":  "-0+32/64",
		"40_24": "-0+08388608/16777216",
		"52_12": "-0+2048/4096",
	},
	floor: -1,
	round: +0,
	ceil:  +0,
}, {
	x: -4.125,
	s: map[string]string{
		"26_6":  "-4+08/64",
		"40_24": "-4+02097152/16777216",
		"52_12": "-4+0512/4096",
	},
	floor: -5,
	round: -4,
	ceil:  -4,
}, {
	x: -7.75,
	s: map[string]string{
		"26_6":  "-7+48/64",
		"40_24": "-7+12582912/16777216",
		"52_12": "-7+3072/4096",
	},
	floor: -8,
	round: -8,
	ceil:  -7,
}}

func I(i int) Fixed {
	return Fixed{int64(i)}
}

func TestFixed(t *testing.T) {
	one := Raw(oneValue)
	for i, tc := range testCases {
		x := From(tc.x)
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no string representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else if got, want := x.String(), s; got != want {
			t.Errorf("tc.x=%v: String: got %q, want %q", tc.x, got, want)
		}
		if got, want := x.Floor(), tc.floor; got != want {
			t.Errorf("tc.x=%v: Floor: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Round(), tc.round; got != want {
			t.Errorf("tc.x=%v: Round: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Ceil(), tc.ceil; got != want {
			t.Errorf("tc.x=%v: Ceil: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Mul(one), x; got != want {
			t.Errorf("tc.x=%v: Mul by one: got %v, want %v", tc.x, got, want)
		}
	}
}

var mulTestCases = []struct {
	x float64
	y float64
	z map[string]float64
	s map[string]string
}{{
	x: 0,
	y: 1.5,
	z: map[string]float64{
		"26_6":  0,
		"40_24": 0,
		"52_12": 0,
	},
	s: map[string]string{
		"26_6":  "0+00/64",
		"40_24": "0+00000000/16777216",
		"52_12": "0+0000/4096",
	},
}, {
	x: +1.25,
	y: +4,
	z: map[string]float64{
		"26_6":  +5,
		"40_24": +5,
		"52_12": +5,
	},
	s: map[string]string{
		"26_6":  "5+00/64",
		"40_24": "5+00000000/16777216",
		"52_12": "5+0000/4096",
	},
}, {
	x: +1.25,
	y: -4,
	z: map[string]float64{
		"26_6":  -5,
		"40_24": -5,
		"52_12": -5,
	},
	s: map[string]string{
		"26_6":  "-5+00/64",
		"40_24": "-5+00000000/16777216",
		"52_12": "-5+0000/4096",
	},
}, {
	x: -1.25,
	y: +4,
	z: map[string]float64{
		"26_6":  -5,
		"40_24": -5,
		"52_12": -5,
	},
	s: map[string]string{
		"26_6":  "-5+00/64",
		"40_24": "-5+00000000/16777216",
		"52_12": "-5+0000/4096",
	},
}, {
	x: -1.25,
	y: -4,
	z: map[string]float64{
		"26_6":  +5,
		"40_24": +5,
		"52_12": +5,
	},
	s: map[string]string{
		"26_6":  "5+00/64",
		"40_24": "5+00000000/16777216",
		"52_12": "5+0000/4096",
	},
}, {
	x: 1.25,
	y: 1.5,
	z: map[string]float64{
		"26_6":  1.875,
		"40_24": 1.875,
		"52_12": 1.875,
	},
	s: map[string]string{
		"26_6":  "1+56/64",
		"40_24": "1+14680064/16777216",
		"52_12": "1+3584/4096",
	},
}, {
	x: 1234.5,
	y: -8888.875,
	z: map[string]float64{
		"26_6":  -10973316.1875,
		"40_24": -10973316.1875,
		"52_12": -10973316.1875,
	},
	s: map[string]string{
		"26_6":  "-10973316+12/64",
		"40_24": "-10973316+03145728/16777216",
		"52_12": "-10973316+0768/4096",
	},
}, {
	x: 1.515625, // 1 + 33/64 = 97/64
	y: 1.531250, // 1 + 34/64 = 98/64
	z: map[string]float64{
		"26_6":  2.32080078125, // 2 + 1314/4096 = 9506/4096
		"40_24": 2.32080078125, // 2 + 1314/4096 = 9506/4096
		"52_12": 2.32080078125, // 2 + 1314/4096 = 9506/4096
	},
	s: map[string]string{
		"26_6":  "2+21/64",             // 2.32812500000, which is closer than 2+20/64 (in decimal, 2.3125)
		"40_24": "2+05382144/16777216", // 2.32080078125
		"52_12": "2+1314/4096",         // 2.32080078125
	},
}, {
	x: 0.500244140625, // 2049/4096, approximately 32/64
	y: 0.500732421875, // 2051/4096, approximately 32/64
	z: map[string]float64{
		"26_6":  0.25,               // 4194304/16777216, or 1024/4096
		"40_24": 0.2504884600639343, // 4202499/16777216
		"52_12": 0.2504884600639343, // 4202499/16777216
	},
	s: map[string]string{
		"26_6":  "0+16/64",             // 0.25000000000
		"40_24": "0+04202499/16777216", // 0.2504884600639343
		"52_12": "0+1026/4096",         // 0.25048828125, which is closer than 0+1027/4096 (in decimal, 0.250732421875)
	},
}, {
	x: 0.015625,       // 1/64
	y: 0.000244140625, // 1/4096, approximately 0/64
	z: map[string]float64{
		"26_6":  0.0,                  // 0
		"40_24": 0.000003814697265625, // 1/262144
		"52_12": 0.000003814697265625, // 1/262144
	},
	s: map[string]string{
		"26_6":  "0+00/64",             // 0
		"40_24": "0+00000064/16777216", // 0.000003814697265625
		"52_12": "0+0000/4096",         // 0, which is closer than 0+0001/4096 (in decimal, 0.000244140625)
	},
}, {
	// Round the Fixed calculation down.
	x: 1.44140625, // 1 + 1808/4096 = 5904/4096, approximately 92/64
	y: 1.44140625, // 1 + 1808/4096 = 5904/4096, approximately 92/64
	z: map[string]float64{
		"26_6":  2.06640625,         // 2 +  272/4096 = 8464/4096
		"40_24": 2.0776519775390625, // 2 +  318/4096 +  256/16777216 = 34857216/16777216
		"52_12": 2.0776519775390625, // 2 +  318/4096 +  256/16777216 = 34857216/16777216
	},
	s: map[string]string{
		"26_6":  "2+04/64",             // 2.06250000000, which is closer than 2+05/64     (in decimal, 2.078125000000)
		"40_24": "2+01302784/16777216", // 2.0776519775390625
		"52_12": "2+0318/4096",         // 2.07763671875, which is closer than 2+0319/4096 (in decimal, 2.077880859375)
	},
}, {
	// Round the Fixed calculation up.
	x: 1.44140625,     // 1 + 1808/4096 = 5904/4096, approximately 92/64
	y: 1.441650390625, // 1 + 1809/4096 = 5905/4096, approximately 92/64
	z: map[string]float64{
		"26_6":  2.06640625,         // 2 +  272/4096 = 8464/4096
		"40_24": 2.0780038833618164, // 2 +  319/4096 + 2064/16777216 = 34863120/16777216
		"52_12": 2.0780038833618164, // 2 +  319/4096 + 2064/16777216 = 34863120/16777216
	},
	s: map[string]string{
		"26_6":  "2+04/64",             // 2.06250000000, which is closer than 2+05/64     (in decimal, 2.078125000000)
		"40_24": "2+01308688/16777216", // 2.0780038833618164
		"52_12": "2+0320/4096",         // 2.07812500000, which is closer than 2+0319/4096 (in decimal, 2.077880859375)
	},
}}

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
		} else if got, want := x.Mul(y).String(), s; got != want {
			t.Errorf("tc.x=%v: Mul: got %q, want %q", tc.x, got, want)
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
	Fixed_Div_(t, func(a, b Fixed) Fixed { return a.UnsafeDiv(b) })
}

//noinspection GoUnusedGlobalVariable
var Result Fixed

func (a Fixed) refMul(b Fixed) Fixed {
	x := a.int64
	y := b.int64
	xw, yw := x>>fracBits, y>>fracBits     // Whole part of x and y.
	xf, yf := x&fracMask, y&fracMask       // Fractional part of x and y.
	ret := (xw * yw) << fracBits           // Multiply whole part.
	ret += xw*yf + yw*xf                   // Multiply each whole by other fraction.
	ret += (xf * yf) >> fracBits           // Multiply fractions by each other.
	ret += (xf * yf) >> (fracBits - 1) & 1 // Round to nearest, instead of rounding down.
	return Fixed{ret}
}

func BenchmarkFixed_refMul(b *testing.B) {
	var r Fixed
	for i := 0; i < b.N; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.refMul(y)
	}
	Result = r
}

func (x Fixed) refDiv(y Fixed) Fixed {
	ret := x.int64 / y.int64 << fracBits                          // Calculate whole part.
	ret += x.int64 % y.int64 << fracBits / y.int64                // Add fractional part.
	ret += x.int64 % y.int64 << fracBits % y.int64 << 1 / y.int64 // Round to nearest, instead of rounding down.
	return Fixed{ret}
}

func BenchmarkFixed_refDiv(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.refDiv(y)
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

func BenchmarkFixed_UnsafeMul(b *testing.B) {
	var r Fixed
	for i := 0; i < b.N; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.UnsafeMul(y)
	}
	Result = r
}

func BenchmarkFixed_UnsafeDiv(b *testing.B) {
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

func BenchmarkFixed_UnsafeAdd(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.UnsafeAdd(y)
	}
	Result = r
}

func BenchmarkFixed_Add(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i<<(fracBits-3)))
		r = x.Add(y)
	}
	Result = r
}

func BenchmarkFixed_UnsafeSub(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i))
		r = x.UnsafeSub(y)
	}
	Result = r
}

func BenchmarkFixed_Sub(b *testing.B) {
	var r Fixed
	for i := 1; i < b.N+1; i++ {
		x, y := Raw(int64(3*i)), Raw(int64(i))
		r = x.Sub(y)
	}
	Result = r
}
