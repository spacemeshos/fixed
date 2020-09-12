package fixed

import (
	"testing"
)

func TestFixedMul(t *testing.T) {
	for i, tc := range mulTestCases[6:] {
		x := From(tc.x)
		y := From(tc.y)
		s := tc.s
		if got, want := x.Mul(y).String(), s; got != want {
			t.Errorf("%d: %v*%v =>  %q, want %q", i, tc.x, tc.y, got, want)
		}
	}
}
