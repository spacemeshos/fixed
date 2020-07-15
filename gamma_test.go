package fixed

import (
	"fmt"
	"math"
	"testing"
)

var gammaE = (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>5)))))

func Test_GammaLn(t *testing.T) {
	for i := oneValue + 1; i < 100*oneValue; i += oneValue >> 3 {
		a := Fixed{i}
		y := GammaLn(a)
		got := y.Float()
		want := math.Log(math.Gamma(a.Float()))
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

func Test_GenParam(t *testing.T) {
	d := []float64{
		2.48574089138753565546e-5,
		1.05142378581721974210,
		-3.45687097222016235469,
		4.51227709466894823700,
		-2.98285225323576655721,
		1.05639711577126713077,
		-1.95428773191645869583e-1,
		1.70970543404441224307e-2,
		-5.71926117404305781283e-4,
		4.63399473359905636708e-6,
		-2.71994908488607703910e-9,
	}
	for _, a := range d {
		fmt.Printf("%#016x, // x >> 54-fracBits => %v \n", int64(int64(a*(1<<54))), a)
	}
}

func BenchmarkFixed_LogGamma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = GammaLn(Fixed{(int64(i) + oneValue) % (100 * oneValue)})
	}
}

func BenchmarkFixed_RefLogGamma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = From(math.Log(math.Gamma(Fixed{(int64(i) + oneValue) % (100 * oneValue)}.Float())))
	}
}
