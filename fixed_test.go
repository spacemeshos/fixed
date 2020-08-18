// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fixed

import (
	"fmt"
	"math/bits"
	"math/rand"
	"testing"
)

func I(i int) Fixed {
	return Fixed{int64(i)}
}

func TestFixed(t *testing.T) {
	one := Raw(oneValue)
	for i, tc := range testCases {
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if s == "overflow" {
			func() {
				defer func() { recover() }()
				x := From(tc.x)
				t.Errorf("tc.x=%v: From: got %v, want overflow", tc.x, x)
			}()
		} else {
			x := From(tc.x)
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
}

func TestFixed_From(t *testing.T) {
	f := func(v float64) int64 { return int64(v * float64(oneValue)) }
	if got, want := From(1).int64, f(1); got != want {
		t.Errorf("got {%v} != want {%v}", Fixed{got}, Fixed{want})
	}
	if got, want := From(0.5).int64, f(0.5); got != want {
		t.Errorf("got {%v} != want {%v}", Fixed{got}, Fixed{want})
	}
	if got, want := From(1.000000000000313).int64, f(1.000000000000313); got != want {
		t.Errorf("got {%v} != want {%v}", Fixed{got}, Fixed{want})
	}
	if got, want := From(3).int64, f(3); got != want {
		t.Errorf("got {%v} != want {%v}", Fixed{got}, Fixed{want})
	}
	if got, want := From(300).int64, f(300); got != want {
		t.Errorf("got {%v} != want {%v}", Fixed{got}, Fixed{want})
	}
	if got, want := From(3.128608483393691e-13).int64, f(3.128608483393691e-13); got != want {
		t.Errorf("got {%v} != want {%v}", Fixed{got}, Fixed{want})
	}
	func() {
		defer func() { recover() }()
		x := From(8388607.9677419355)
		t.Errorf("got %v, want overflow", x)
	}()
}

func TestFixed_Bytes(t *testing.T) {
	rand.Seed(42)
	for i := 0; i < 100; i++ {
		v := rand.Int63()
		if rand.Float64() < 0.5 {
			v = -v
		}
		b := Fixed{v}.Bytes()
		a := FromBytes(b)
		if a.int64 != v {
			t.Errorf("got {%v} != want {%v}", a, Fixed{v})
			t.FailNow()
		}
		a = FracFromBytes(b)
		if a.int64 != v&fracMask {
			t.Errorf("got {%v} != want {%v}", a, Fixed{v & fracMask})
			t.FailNow()
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

type accuracy struct {
	min, max int
	emax, emin, eavg float64
	avg, count float64
	Epsilon    int64
}

func (acc *accuracy) update(got Fixed, fwant float64) bool {
	epsilon := int64(0)
	if acc.Epsilon == 0 {
		epsilon = int64(1) << testEpsilonBits
	} else {
		epsilon = acc.Epsilon
	}
	want := From(fwant)
	if sign(got.int64) != sign(want.int64) {
		return false
	}
	x, y := uint64(abs(got.int64)), uint64(abs(want.int64))
	a, b := bits.Len64(x), bits.Len64(y)
	if a != b {
		if b > a {
			a = b
		}
	}
	t := 0
	for i := a - 1; i >= 0; i-- {
		if (x>>i)&1 != (y>>i)&1 {
			break
		}
		t++
	}
	if a < fracBits {
		t = fracBits - a + t
	}
	if abs(got.int64-want.int64) > epsilon {
		return false
	}
	if t > acc.max {
		acc.max = t
	}
	if acc.count == 0 || acc.min > t {
		acc.min = t
	}
	e := float64(abs(got.int64-want.int64))/float64(oneValue)
	if e > acc.emax {
		acc.emax = e
	}
	if acc.count == 0 || acc.emin > e {
		acc.emin = e
	}
	acc.eavg += e
	acc.avg += float64(t)
	acc.count += 1
	return true
}

func (acc accuracy) String() string {
	return fmt.Sprintf("matched { bits max: %v min: %v, avg: %.2f || epsilon max: %.8g min: %.8g, avg: %.8g }",
		acc.max, acc.min, acc.avg/acc.count,
		acc.emax, acc.emin, acc.eavg/acc.count)
}

func randomFixed(i int64) int64 {
	return i + rand.Int63n(oneValue)>>(fracBits-12)
}
