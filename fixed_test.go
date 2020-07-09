// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fixed

import (
	"fmt"
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

//noinspection GoUnusedGlobalVariable
var Result Fixed

func BenchmarkFixed_AddUnsafe(b *testing.B) {
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

func BenchmarkFixed_SubUnsafe(b *testing.B) {
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
