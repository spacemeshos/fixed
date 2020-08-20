package fixed

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestFixed_BinCDF(t *testing.T) {
	fmt.Printf("%.10f\n", BinCDF(New(34), From(0.9), New(25)).Float())
}

func TestFixed_BinCDF1(t *testing.T) {
	fmt.Printf("%.10f\n", BinCDF(New(340), From(0.82), New(250)).Float())
}

func TestFixed_BinCDF2(t *testing.T) {
	const Max = 1<<(totalBits-fracBits-2)
	for k := 2; k < Max; k+=10 {
		BinCDF(New(k), From(.05), New(2))
	}
}

func TestFixed_BinCDF3(t *testing.T) {
	acc := accuracy{}
	for i, tc := range bincdfTestCases {
		s, found := tc.s[fmt.Sprintf("%d_%d", totalBits-fracBits, fracBits)]
		if !found {
			t.Logf("case #%d has no string representation for %d_%d", i, totalBits-fracBits, fracBits)
		} else {
			if s == "overflow" {
				func() {
					defer func() { recover() }()
					x, n, p := New(tc.x), New(tc.n), From(tc.p)
					got := BinCDF(n,p,x)
					t.Errorf("BinCDF: got %q, want overflow", got)
				}()
			} else {
				x, n, p := New(tc.x), New(tc.n), From(tc.p)
				got := BinCDF(n,p,x)
				acc.Epsilon = 1<<(bits.Len64(uint64(tc.n))+8)
				if ok := acc.update(got, tc.cdf); !ok {
					t.Errorf("BinCDF(%v,%v,%v) => got %v|%v, want %v|%v",  tc.n, tc.p, tc.x, got, got.Float(), From(tc.cdf), tc.cdf)
				}
			}
		}
	}
	t.Log(acc)
}
