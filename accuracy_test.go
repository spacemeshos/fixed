package fixed

import (
	"fmt"
	"math"
	"math/rand"
)

const testEpsilonBits int = 3

type accuracy struct {
	min, max         int
	emax, emin, eavg float64
	avg, count       float64
	Epsilon          float64
}

func (acc *accuracy) update(got Fixed, fwant float64) bool {
	epsilon := float64(0)
	if acc.Epsilon == 0 {
		epsilon = float56(int64(1) << testEpsilonBits)
	} else {
		epsilon = acc.Epsilon
	}
	want := From(fwant)
	e := math.Abs(got.float() - want.float())
	if e > epsilon {
		return false
	}

	a, b := got.bitlen(), want.bitlen()
	if a != b {
		if b > a {
			a = b
		}
	}
	t := 0
	if a > fracBits {
		a = fracBits
	}
	for i := a; i >= 0; i-- {
		if got.bit(i) != want.bit(i) {
			break
		}
		t++
	}
	if t > acc.max {
		acc.max = t
	}
	if acc.count == 0 || acc.min > t {
		acc.min = t
	}
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
	return fmt.Sprintf("matched { bits max: %v min: %v, avg: %.2f || epsilon max: %.3g min: %.3g, avg: %.3g }",
		acc.max, acc.min, acc.avg/acc.count,
		acc.emax, acc.emin, acc.eavg/acc.count)
}

func randomFixed(i int64) int64 {
	return i + rand.Int63n(oneValue)>>(fracBits-12)
}
