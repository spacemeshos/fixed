package fixed

import (
	"math"
	"math/bits"
	"math/rand"
	"testing"
)

func Test_Lgamma1(t *testing.T) {
	acc := accuracy{Epsilon: 64}
	rand.Seed(43)
	step := oneValue >> 4
	for i := step; i < 2000*oneValue; i += step {
		a := randomFixed(i)
		y := Lgamma(Fixed{a})
		got := y.Float()
		want, _ := math.Lgamma(float(a))
		acc.Epsilon = 1 << max(int64(bits.Len64(uint64(abs(y.int64)))-fracBits+2), 2)
		if ok := acc.update(y, want); !ok {
			t.Errorf("gamma(%v) => got %v|%v, want %v|%v", float(a), y, got, From(want), want)
			t.FailNow()
		}
	}
	t.Log(acc)
}

func BenchmarkFixed_LogGamma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = Lgamma(Fixed{(int64(i) + oneValue) % (100 * oneValue)})
	}
}

func BenchmarkFixed_RefLogGamma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = From(math.Log(math.Gamma(Fixed{(int64(i) + oneValue) % (100 * oneValue)}.Float())))
	}
}
