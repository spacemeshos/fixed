package fixed

import (
	"math"
	"math/rand"
	"testing"
)

func TestRange_Log(t *testing.T) {
	rand.Seed(42)
	step := oneValue >> 12
	acc := accuracy{Epsilon: 1 << 3}
	for i := oneValue; i < 27*oneValue; i += step {
		a := randomFixed(int64(i))
		y := Log(Fixed{a})
		want := math.Log(float(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log(%v|%v) => %v: got %v, want %v|%v", Fixed{a}, float(a), y, y.Float(), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func TestRange_Log2(t *testing.T) {
	rand.Seed(42)
	//step := oneValue >> 12
	acc := accuracy{Epsilon: 1 << max(int64(fracBits-46), 2)}
	for i := 1; i < 1000000; i++ {
		a := randomFixed(int64(i))
		y := Log2(Fixed{a})
		want := math.Log2(float(a))
		if ok := acc.update(y, want); !ok {
			t.Errorf("log2(%v|%v) => %v: got %v, want %v|%v", Fixed{a}, float(a), y, y.Float(), From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func BenchmarkFixed_Log(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Log(Fixed{int64(i)})
	}
}

func BenchmarkFixed_LogRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Log(Fixed{int64(i)}.Float()))
	}
}
