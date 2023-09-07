package fixed

import (
	"math"
	"math/rand"
	"testing"
)

func TestRange_Log(t *testing.T) {
	step := oneValue >> 12
	acc := accuracy{Epsilon: 1e-15}
	for i := step; i < 63*oneValue; i += step {
		a := randomFixed(int64(i))
		y := log(a)
		want := math.Log(float56(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v|%v) => %v|%v, want %v|%v", rawfixed(a), float56(a), y, y.float(), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_Log2(t *testing.T) {
	step := oneValue >> 12
	acc := accuracy{Epsilon: 1e-14}
	for i := step; i < 63*oneValue; i += step {
		a := randomFixed(int64(i))
		y := rawfixed(log2(a, fracBits))
		want := math.Log2(float56(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v|%v) => %v|%v, want %v|%v", rawfixed(a), float56(a), y, y.float(), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_iLog(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	acc := accuracy{Epsilon: 1e-15}
	maxarg := int64(100000)
	step := maxarg / 200
	for i := int64(0); i < 1000; i += step {
		a := i + rng.Int63n(step)
		y := rawfixed(ilog56(a))
		want := math.Log(float64(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v) => %v|%v, want %v|%v", a, y, y.float(), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_iLog2(t *testing.T) {
	acc := accuracy{Epsilon: 1e-14}
	maxarg := int64(100000)
	step := maxarg / 200
	for i := int64(0); i < 1000; i += step {
		a := i + rand.Int63n(step)
		y := rawfixed(log2(a, 0))
		want := math.Log2(float64(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v) => %v|%v, want %v|%v", a, y, y.float(), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}
