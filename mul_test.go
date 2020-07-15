package fixed

import (
	"fmt"
	"math"
	"testing"
)

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
