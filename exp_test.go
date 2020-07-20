package fixed

import (
	"math"
	"math/bits"
	"math/rand"
	"testing"
)

func TestRange_NegExp(t *testing.T) {
	acc := accuracy{Epsilon: 1 << max(int64(fracBits-46), 2)}
	rand.Seed(42)
	step := oneValue >> 20
	t.Logf("max possible exponent argument is %g, step %g\n", float(maxEpow), float(step))
	for i := -maxEpow; i <= 0; i += step {
		a := randomFixed(i)
		y := Exp(Fixed{a})
		got := y.Float()
		want := math.Exp(float(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("exp(%v) => got %v|%v, want %v|%v", float(a), y, got, From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_Exp(t *testing.T) {
	acc := accuracy{}
	rand.Seed(43)
	step := oneValue >> 20
	t.Logf("max possible exponent argument is %g, step %g\n", float(maxEpow), float(step))
	for i := fixed(0); i < maxEpow; i += step {
		a := randomFixed(i)
		y := Exp(Fixed{a})
		got := y.Float()
		want := math.Exp(float(a))
		acc.Epsilon = 1 << max(int64(bits.Len64(uint64(abs(y.int64)))-fracBits+2), max(int64(fracBits-40), 2))
		if ok := acc.update(y, want); !ok {
			t.Errorf("exp(%v) => got %v|%v, want %v|%v", float(a), y, got, From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_Exp27(t *testing.T) {
	func() {
		defer func() { recover() }()
		x := New(27)
		Exp(x)
		t.Errorf("exp(%v) must panic by overflow", x)
	}()
}

func BenchmarkFixed_Exp(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Exp(Fixed{int64(i) % maxEpow})
	}
}

func BenchmarkFixed_ExpRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Exp(Fixed{int64(i) % maxEpow}.Float()))
	}
}

func BenchmarkFixed_NegExp(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Exp(Fixed{-int64(i) % maxEpow})
	}
}

func BenchmarkFixed_NegExpRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Exp(Fixed{-int64(i) % maxEpow}.Float()))
	}
}

func BenchmarkFixed_FullExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = Exp(Fixed{int64(i)*(maxEpow*2)/int64(b.N) - maxEpow})
	}
}

func BenchmarkFixed_FullExpRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = From(math.Exp(Fixed{int64(i)*(maxEpow*2)/int64(b.N) - maxEpow}.Float()))
	}
}
