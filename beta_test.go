package fixed

import (
	"math"
	"testing"
)

var betaRegDat = []struct{ a, b, x, r float64 }{
	{0.5, 0.5, 0.5, 0.5},
	{0.5, 0.5, 1.0, 1.0},
	{1.0, 0.5, 0.5, 0.292893218813452475599},
	{1.0, 0.5, 1.0, 1.0},
	{2.5, 0.5, 0.5, 0.07558681842161243795},
	{2.5, 0.5, 1.0, 1.0},
	{0.5, 1.0, 0.5, 0.7071067811865475244},
	{0.5, 1.0, 1.0, 1.0},
	{1.0, 1.0, 0.5, 0.5},
	{1.0, 1.0, 1.0, 1.0},
	{2.5, 1.0, 0.5, 0.1767766952966368811},
	{2.5, 1.0, 1.0, 1.0},
	{0.5, 2.5, 0.5, 0.92441318157838756205},
	{0.5, 2.5, 1.0, 1.0},
	{1.0, 2.5, 0.5, 0.8232233047033631189},
	{1.0, 2.5, 1.0, 1.0},
	{2.5, 2.5, 0.5, 0.5},
	{2.5, 2.5, 1.0, 1.0},
}

var betaEpsilon = 1 / (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>1)))))

func TestFixed_BetaReg(t *testing.T) {
	for i, tx := range betaRegDat {
		y := BetaReg(From(tx.a), From(tx.b), From(tx.x))
		got := y.Float()
		want := tx.r
		if got <= want-betaEpsilon || got >= want+betaEpsilon {
			t.Errorf("%d: BetReg(%v,%v,%v) => %v: got %v, want %v | %v in (%v,%v)",
				i, tx.a, tx.b, tx.x, y, got, want, From(want), want-betaEpsilon, want+betaEpsilon)
			t.FailNow()
		}
	}
}

func BenchmarkFixed_BetaReg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tx := betaRegDat[i%len(betaRegDat)]
		Result = BetaReg(From(tx.a), From(tx.b), From(tx.x))
	}
}
