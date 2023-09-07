package fixed

import (
	"math"
	"testing"
)

func TestRange_NegExp(t *testing.T) {
	maxEpow := fixed56(34)
	acc := accuracy{Epsilon: float56(1 << max64(int64(fracBits-46), 2))}
	step := maxEpow / 10000
	t.Logf("max possible exponent argument is %g, step %g\n", float56(maxEpow), float56(step))
	for i := -maxEpow; i <= 0; i += step {
		a := randomFixed(i)
		y := exp56(a)
		got := y.float()
		want := math.Exp(float56(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("exp(%v) => got %v|%v, want %v|%v", float56(a), y, got, From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_Exp(t *testing.T) {
	maxEpow := fixed56(33) // fixed56(88)
	acc := accuracy{}
	step := maxEpow / 1000
	t.Logf("max possible exponent argument is %g, step %g\n", float56(maxEpow), float56(step))
	for i := fixed56(0); i < maxEpow; i += step {
		a := randomFixed(i)
		y := exp56(a)
		got := y.float()
		want := math.Exp(float56(a))
		acc.Epsilon = float56(1 << min64(max64(int64(y.bitlen()-24), 2), 62))
		if ok := acc.update(y, want); !ok {
			t.Errorf("exp(%v) => got %v|%v, want %v|%v, eps: %v", float56(a), y, got, From(want), want, acc.Epsilon)
			t.FailNow()
		}
	}
	t.Log(acc)
}
