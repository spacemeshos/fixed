package fixed

import (
	"math"
	"math/rand"
	"testing"
)

func Test_Lgamma1(t *testing.T) {
	acc := accuracy{Epsilon: 1e-10}
	rand.Seed(42)
	maxarg := int64(1000000)
	step := maxarg / 1000
	for i := int64(0); i < 10000; i += step {
		a := i + rand.Int63n(step)
		y := lgamma(a)
		got := y.Float()
		want, _ := math.Lgamma(float64(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("lgamma(%v) => got %v|%v, want %v|%v, eps: %v", float64(a), y, got, From(want), want, acc.Epsilon)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func Test_Lgamma2(t *testing.T) {
	acc := accuracy{Epsilon: 1e-11}
	for a := int64(1); a < 1000; a++ {
		y := lgamma(a)
		got := y.Float()
		want, _ := math.Lgamma(float64(a))
		//acc.Epsilon = float56(1 << min64(max64(int64(y.val.BitLen()-24), 2),62))
		if ok := acc.update(y, want); !ok {
			t.Errorf("lgamma(%v) => got %v|%v, want %v|%v, eps: %v", float64(a), y, got, From(want), want, acc.Epsilon)
			t.FailNow()
		}
	}
	t.Log(acc)
}
