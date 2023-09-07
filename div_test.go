package fixed

import (
	"testing"
)

func testDivfunc(t *testing.T, x, y, z float64, s string) {
	if s == "overflow" {
		func() {
			defer func() { recover() }()
			x, y := from(x), from(y)
			got := div(x, y)
			t.Errorf("Div: got %q, want overflow", got)
		}()
	} else {
		a, b := from(x), from(y)
		got := div(a, b)
		// t.Logf("* x.Div(y): %s|%v, want: %s|%v, %v|%v / %v|%v", got, got.float(), s, z, a, x, b, y)
		if got.format() != s {
			t.Errorf("x.Div(y): %s|%v, want: %s|%v, %v/%v", got, got.float(), s, z, x, y)
		}
	}
}

func TestFixed_Div1(t *testing.T) {
	for _, tc := range divTestCases {
		testDivfunc(t, float64(tc.x), float64(tc.y), tc.z, tc.s)
	}
}
