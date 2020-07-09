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
		want:= math.Exp(Fixed{a.int64}.Float())
		epsilon := got / 1000
		if got < want-epsilon || got > want+epsilon {
			t.Errorf("exp(%v) => %v: got %v, want %v", a.Float(), y, got, want)
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
