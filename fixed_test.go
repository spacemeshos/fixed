package fixed

import (
	"testing"
)

func max64(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func min64(x int64, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func TestFixed(t *testing.T) {
	for _, tc := range testCases {
		x := From(tc.x)
		if x.Float() != tc.x {
			t.Errorf("tc.x=%v: Float: got %v, want %v", tc.x, x.Float(), x)
		}
		if got, want := x.String(), tc.s; got != want {
			t.Errorf("tc.x=%v: String: got %q, want %q", tc.x, got, want)
		}
		if got, want := x.Floor(), int64(tc.floor); got != want {
			t.Errorf("tc.x=%v: Floor: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Round(), int64(tc.round); got != want {
			t.Errorf("tc.x=%v: Round: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Ceil(), int64(tc.ceil); got != want {
			t.Errorf("tc.x=%v: Ceil: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Mul(One()), x; got != want {
			t.Errorf("tc.x=%v: Mul by one: got %v, want %v", tc.x, got, want)
		}
	}
}
