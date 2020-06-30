// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fixed

import (
	"fmt"
	"math"
	"math/rand"
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

func TestFixed(t *testing.T) {
	const one = Fixed(1 << fracBits)
	for i, tc := range testCases {
		x := Fixed(tc.x * (1 << fracBits))
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
		if totalBits != 32 {
			continue
		}
		if got, want := x.mul(one), x; got != want {
			t.Errorf("tc.x=%v: mul by one: got %v, want %v", tc.x, got, want)
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
		"26_6":  "2+21/64",     // 2.32812500000, which is closer than 2:20 (in decimal, 2.3125)
		"40_24": "2+05382144/16777216",
		"52_12": "2+1314/4096", // 2.32080078125
	},
}, {
	x: 0.500244140625, // 2049/4096, approximately 32/64
	y: 0.500732421875, // 2051/4096, approximately 32/64
	z: map[string]float64{
		"26_6":  0.25,               // 4194304/16777216, or 1024/4096
		"40_24": 0.2504884600639343,
		"52_12": 0.2504884600639343, // 4202499/16777216
	},
	s: map[string]string{
		"26_6":  "0+16/64",     // 0.25000000000
		"40_24": "0+04202499/16777216",
		"52_12": "0+1026/4096", // 0.25048828125, which is closer than 0:1027 (in decimal, 0.250732421875)
	},
}, {
	x: 0.015625,       // 1/64
	y: 0.000244140625, // 1/4096, approximately 0/64
	z: map[string]float64{
		"26_6":  0.0,                  // 0
		"40_24": 3.814697265625e-06,
		"52_12": 0.000003814697265625, // 1/262144
	},
	s: map[string]string{
		"26_6":  "0+00/64",     // 0
		"40_24": "0+00000064/16777216",
		"52_12": "0+0000/4096", // 0, which is closer than 0:0001 (in decimal, 0.000244140625)
	},
}, {
	// Round the Fixed calculation down.
	x: 1.44140625, // 1 + 1808/4096 = 5904/4096, approximately 92/64
	y: 1.44140625, // 1 + 1808/4096 = 5904/4096, approximately 92/64
	z: map[string]float64{
		"26_6":  2.06640625,         // 2 +  272/4096 = 8464/4096
		"40_24": 2.0776519775390625,
		"52_12": 2.0776519775390625, // 2 +  318/4096 +  256/16777216 = 34857216/16777216
	},
	s: map[string]string{
		"26_6":  "2+04/64",     // 2.06250000000, which is closer than 2:05   (in decimal, 2.078125000000)
		"40_24": "2+01302784/16777216",
		"52_12": "2+0318/4096", // 2.07763671875, which is closer than 2:0319 (in decimal, 2.077880859375)
	},
}, {
	// Round the Fixed calculation up.
	x: 1.44140625,     // 1 + 1808/4096 = 5904/4096, approximately 92/64
	y: 1.441650390625, // 1 + 1809/4096 = 5905/4096, approximately 92/64
	z: map[string]float64{
		"26_6":  2.06640625,         // 2 +  272/4096 = 8464/4096
		"40_24": 2.0780038833618164,
		"52_12": 2.0780038833618164, // 2 +  319/4096 + 2064/16777216 = 34863120/16777216
	},
	s: map[string]string{
		"26_6":  "2+04/64",     // 2.06250000000, which is closer than 2:05   (in decimal, 2.078125000000)
		"40_24": "2+01308688/16777216",
		"52_12": "2+0320/4096", // 2.07812500000, which is closer than 2:0319 (in decimal, 2.077880859375)
	},
}}

func TestFixedMul(t *testing.T) {
	for i, tc := range mulTestCases {
		x := Fixed(tc.x * (1 << fracBits))
		y := Fixed(tc.y * (1 << fracBits))
		want, found := tc.z[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no float representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else if z := float64(x) * float64(y) / (1 << (fracBits * 2)); z != want {
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
	const (
		oneMinusIota  = Fixed(1<<fracBits) - 1
		oneMinusIotaF = float64(oneMinusIota) / (1 << fracBits)
	)

	for _, neg := range []bool{false, true} {
		for i := uintptr(0); i < totalBits; i++ {
			x := Fixed(1 << i)
			if neg {
				x = -x
			} else if i == totalBits-1 {
				// A signed int32 can't represent 1<<31.
				// A signed int64 can't represent 1<<63.
				continue
			}

			// want equals x * oneMinusIota, rounded to nearest.
			want := Fixed(0)
			if -1<<fracBits < x && x < 1<<fracBits {
				// (x * oneMinusIota) isn't exactly representable as an
				// Fixed. Calculate the rounded value using float64 math.
				xF := float64(x) / (1 << fracBits)
				wantF := xF * oneMinusIotaF * (1 << fracBits)
				want = Fixed(math.Floor(wantF + 0.5))
			} else {
				// (x * oneMinusIota) is exactly representable.
				want = oneMinusIota << (i - fracBits)
				if neg {
					want = -want
				}
			}

			if got := x.Mul(oneMinusIota); got != want {
				t.Errorf("neg=%t, i=%d, x=%v, Mul: got %v, want %v", neg, i, x, got, want)
			}
			if totalBits != 32 {
				continue
			}
			if got := x.mul(oneMinusIota); got != want {
				t.Errorf("neg=%t, i=%d, x=%v, mul: got %v, want %v", neg, i, x, got, want)
			}
		}
	}
}

func TestFixedMulVsMul(t *testing.T) {
	if totalBits != 32 {
		return
	}
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 10000; i++ {
		u := Fixed(rng.Uint32())
		v := Fixed(rng.Uint32())
		Mul := u.Mul(v)
		mul := u.mul(v)
		if Mul != mul {
			t.Errorf("u=%#08x, v=%#08x: Mul=%#08x and mul=%#08x differ",
				uint32(u), uint32(v), uint32(Mul), uint32(mul))
		}
	}
}

func TestMuli32(t *testing.T) {
	rng := rand.New(rand.NewSource(2))
	for i := 0; i < 10000; i++ {
		u := int32(rng.Uint32())
		v := int32(rng.Uint32())
		lo, hi := muli32(u, v)
		got := uint64(lo) | uint64(hi)<<32
		want := uint64(int64(u) * int64(v))
		if got != want {
			t.Errorf("u=%#08x, v=%#08x: got %#016x, want %#016x", uint32(u), uint32(v), got, want)
		}
	}
}

func TestMulu32(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	for i := 0; i < 10000; i++ {
		u := rng.Uint32()
		v := rng.Uint32()
		lo, hi := mulu32(u, v)
		got := uint64(lo) | uint64(hi)<<32
		want := uint64(u) * uint64(v)
		if got != want {
			t.Errorf("u=%#08x, v=%#08x: got %#016x, want %#016x", u, v, got, want)
		}
	}
}

// mul (with a lower case 'm') is an alternative implementation of Fixed.Mul
// (with an upper case 'M'). It has the same structure as the Fixed.Mul
// implementation, but Fixed.mul is easier to test since Go has built-in
// 64-bit integers.
func (x Fixed) mul(y Fixed) Fixed {
	const M, N = 26, 6
	lo, hi := muli32(int32(x), int32(y))
	ret := Fixed(hi<<M | lo>>N)
	ret += Fixed((lo >> (N - 1)) & 1) // Round to nearest, instead of rounding down.
	return ret
}

// muli32 multiplies two int32 values, returning the 64-bit signed integer
// result as two uint32 values.
//
// muli32 isn't used directly by this package, but it has the same structure as
// muli64, and muli32 is easier to test since Go has built-in 64-bit integers.
func muli32(u, v int32) (lo, hi uint32) {
	const (
		s    = 16
		mask = 1<<s - 1
	)

	u1 := uint32(u >> s)
	u0 := uint32(u & mask)
	v1 := uint32(v >> s)
	v0 := uint32(v & mask)

	w0 := u0 * v0
	t := u1*v0 + w0>>s
	w1 := t & mask
	w2 := uint32(int32(t) >> s)
	w1 += u0 * v1
	return uint32(u) * uint32(v), u1*v1 + w2 + uint32(int32(w1)>>s)
}

// mulu32 is like muli32, except that it multiplies unsigned instead of signed
// values.
//
// This implementation comes from $GOROOT/src/runtime/softfloat64.go's mullu
// function, which is in turn adapted from Hacker's Delight.
//
// mulu32 (and its corresponding test, TestMulu32) isn't used directly by this
// package. It is provided in this test file as a reference point to compare
// the muli32 (and TestMuli32) implementations against.
func mulu32(u, v uint32) (lo, hi uint32) {
	const (
		s    = 16
		mask = 1<<s - 1
	)

	u0 := u & mask
	u1 := u >> s
	v0 := v & mask
	v1 := v >> s

	w0 := u0 * v0
	t := u1*v0 + w0>>s
	w1 := t & mask
	w2 := t >> s
	w1 += u0 * v1
	return u * v, u1*v1 + w2 + w1>>s
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

func TestFixed_Div(t *testing.T) {
	for i, tc := range divTestCases {
		x, y := I(tc.x), I(tc.y)
		z := x.Div(y)
		got := z.String()
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no string representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else if got != s {
			t.Errorf("got: %s, want: %s", got, s)
		}
	}
}

func BenchmarkFixed_Mul(b *testing.B) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		x, y := I(rng.Int()), I(rng.Int())
		x.Mul(y)
	}
}
