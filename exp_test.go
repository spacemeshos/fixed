package fixed

import (
	"math"
	"testing"
)

func TestRange_Exp(t *testing.T) {
	for i := 1; i < 130000; i++ {
		a := Fixed{int64(i)}
		y := Exp(a)
		got := y.Float()
		want0 := math.Exp(a.Float())
		want1 := math.Exp(Fixed{a.int64 + 1}.Float())
		if want0 > want1 {
			want0, want1 = want1, want0
		}
		epsilon := got / 1000
		if got < want0-epsilon || got > want1+epsilon {
			t.Errorf("exp(%v) => %v: got %v, want [%v,%v]", a.Float(), y, got, want0, want1)
			t.FailNow()
		}
	}
}

func BenchmarkFixed_Exp(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Result = Exp(Fixed{int64(i % 130000)})
	}
}

func BenchmarkFixed_RefExp(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Result = From(math.Exp(Fixed{int64(i % 130000)}.Float()))
	}
}