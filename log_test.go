package fixed

import (
	"math"
	"testing"
)

var logEpsilon = 1 / (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>1)))))

func TestRange_Log(t *testing.T) {
	for i := 1; i < 1000000; i++ {
		a := Fixed{int64(i)}
		y := Log(a)
		got := y.Float()
		want := math.Log(a.Float())
		if got < want-logEpsilon || got > want+logEpsilon {
			t.Errorf("log(%v|%v) => %v: got %v, want %v", a.int64, a.Float(), y, got, want)
			t.FailNow()
		}
	}
}

var log2Epsilon = 1 / (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>1)))))

func TestRange_Log2(t *testing.T) {
	for i := 1; i < 1000000; i++ {
		a := Fixed{int64(i)}
		y := Log2(a)
		got := y.Float()
		want := math.Log2(a.Float())
		if got < want-log2Epsilon || got > want+log2Epsilon {
			t.Errorf("log2(%v|%v) => %v: got %v, want %v", a.int64, a.Float(), y, got, want)
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

/*
func Test_Gen(t *testing.T) {
	g := []float64 {
		4.18938533204672725052e-01,  // 0x3FDACFE390C97D69
		8.33333333333329678849e-02,  // 0x3FB555555555553B
		-2.77777777728775536470e-03, // 0xBF66C16C16B02E5C
		7.93650558643019558500e-04,  // 0x3F4A019F98CF38B6
		-5.95187557450339963135e-04, // 0xBF4380CB8C0FE741
		8.36339918996282139126e-04,  // 0x3F4B67BA4CDAD5D1
		-1.63092934096575273989e-03, // 0xBF5AB89D0B9E43E4
	}
	//q := 0.6207822
	for i,q := range g {
		fmt.Printf("const g%d = %#x // %v\n",
			i,
			int64(q*float64(uint64(1)<<54)),
			q,
		)
	}
}
*/
