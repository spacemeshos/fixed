package fixed

import (
	"math"
	"testing"
)

var gammaE = (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>5)))))

func Test_Lgamma(t *testing.T) {
	for i := oneValue >> 1; i < 1000*oneValue; i += oneValue >> 3 {
		a := Fixed{i}
		y := Lgamma(a)
		got := y.Float()
		want, _ := math.Lgamma(a.Float())
		gammaEpsilon := 1 / gammaE
		if math.Abs(got) > 1 {
			gammaEpsilon = math.Max(got/gammaE, gammaEpsilon)
		}
		if got < want-gammaEpsilon || got > want+gammaEpsilon {
			t.Errorf("gamma(%v) => %v: got %v, want %v, eps: %v", a.Float(), y, got, want, gammaEpsilon)
			t.FailNow()
		}
	}
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
