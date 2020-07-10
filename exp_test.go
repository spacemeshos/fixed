package fixed

import (
	"math"
	"testing"
)

func TestRange_NegExp(t *testing.T) {
	for i := -144000; i <= 0; i++ {
		a := Fixed{int64(i)}
		y := Exp(a)
		got := y.Float()
		want := math.Exp(Fixed{a.int64}.Float())
		epsilon := 0.001
		if got < want-epsilon || got > want+epsilon {
			t.Errorf("exp(%v) => %v: got %v, want %v", a.Float(), y, got, want)
			t.FailNow()
		}
	}
}

func TestRange_Exp(t *testing.T) {
	for i := 0; i < 144000; i++ {
		a := Fixed{int64(i)}
		y := Exp(a)
		got := y.Float()
		want := math.Exp(Fixed{a.int64}.Float())
		epsilon := got / 2000
		if got < want-epsilon || got > want+epsilon {
			t.Errorf("exp(%v) => %v: got %v, want %v", a.Float(), y, got, want)
			t.FailNow()
		}
	}
}

func TestRange_Exp240(t *testing.T) {
	func() {
		defer func() { recover() }()
		x := Fixed{240000}
		Exp(x)
		t.Errorf("exp(%v) must panic by overflow", x)
	}()
}

func TestRange_ExpEpowMax(t *testing.T) {
	func() {
		defer func() { recover() }()
		x := Fixed{int64(1) << (fracBits + len(epow) + 1)}
		Exp(x)
		t.Errorf("exp(%v) must panic by overflow", x)
	}()
}

func BenchmarkFixed_Exp(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Exp(Fixed{int64(i % 130000)})
	}
}

func BenchmarkFixed_ExpRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Exp(Fixed{int64(i % 130000)}.Float()))
	}
}

func BenchmarkFixed_NegExp(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Exp(Fixed{-int64(i % 130000)})
	}
}

func BenchmarkFixed_NegExpRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Exp(Fixed{-int64(i % 130000)}.Float()))
	}
}

func BenchmarkFixed_FullExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = Exp(Fixed{int64(i)*(144000*2)/int64(b.N) - 144000 })
	}
}

func BenchmarkFixed_FullExpRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = From(math.Exp(Fixed{int64(i)*(144000*2)/int64(b.N) - 144000 }.Float()))
	}
}
