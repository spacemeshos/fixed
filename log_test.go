package fixed

import (
	"math"
	"math/rand"
	"testing"
)

const logEpsilon = 0.001

func TestRandom_Log(t *testing.T) {
	rand.Seed(42)
	for i := 0; i < 100000; i++ {
		a := (rand.Float64() + 1e-5) * 100000
		x := From(a)
		y := Log(x)
		got := y.Float()
		want0 := math.Log(x.Float())
		want1 := math.Log(Fixed{x.int64 + 1}.Float())
		if want0 > want1 {
			want0, want1 = want1, want0
		}
		if got < want0-logEpsilon || got > want1+logEpsilon {
			t.Errorf("log(%v) => %v: got %v, want [%v,%v]", a, y, got, want0, want1)
			t.FailNow()
		}
	}
}

func TestRange_Log(t *testing.T) {
	for i := 1; i < 1000000; i++ {
		a := Fixed{int64(i)}
		y := Log(a)
		got := y.Float()
		want0 := math.Log(a.Float())
		want1 := math.Log(Fixed{a.int64 + 1}.Float())
		if want0 > want1 {
			want0, want1 = want1, want0
		}
		if got < want0-logEpsilon || got > want1+logEpsilon {
			t.Errorf("log(%v) => %v: got %v, want [%v,%v]", a.Float(), y, got, want0, want1)
			t.FailNow()
		}
	}
}

func BenchmarkFixed_Log(b *testing.B) {
	for i := 1; i < 100000; i++ {
		Result = Log(Fixed{int64(i)})
	}
}

func BenchmarkFixed_RefLog(b *testing.B) {
	for i := 1; i < 100000; i++ {
		Result = From(math.Log(Fixed{int64(i)}.Float()))
	}
}
