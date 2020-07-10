package fixed

import (
	"math"
	"testing"
)

const logEpsilon = 0.001

func TestRange_Log(t *testing.T) {
	for i := 1; i < 1000000; i++ {
		a := Fixed{int64(i)}
		y := Log(a)
		got := y.Float()
		want := math.Log(Fixed{a.int64}.Float())
		if got < want-logEpsilon || got > want+logEpsilon {
			t.Errorf("log(%v) => %v: got %v, want %v", a.Float(), y, got, want)
			t.FailNow()
		}
	}
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
