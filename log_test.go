package fixed

import (
	"fmt"
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

func Test_L(t *testing.T) {
	q := 1 / math.Log2(math.E)
	fmt.Printf("%v,%v,%d,%x",
		q,
		q*float64(uint64(1)<<54),
		int64(q*float64(uint64(1)<<54)),
		int64(q*float64(uint64(1)<<54)),
	)
}

func Test_L1(t *testing.T) {
	q := math.Log(2)
	fmt.Printf("%v,%v,%d,%x",
		q,
		q*float64(uint64(1)<<54),
		int64(q*float64(uint64(1)<<54)),
		int64(q*float64(uint64(1)<<54)),
	)
}

func Test_Y(t *testing.T) {
	q := 1 / math.Log(2)
	fmt.Printf("%v,%v,%d,%x\n",
		q,
		q*float64(uint64(1)<<54),
		int64(q*float64(uint64(1)<<54)),
		int64(q*float64(uint64(1)<<54)),
	)
	fmt.Println(Fixed{int64(q*float64(uint64(1)<<54)) >> (54 - fracBits)})
	fmt.Println(Fixed{int64(q*float64(uint64(1)<<54)) >> (54 - fracBits)}.Float())
}

func Test_Z(t *testing.T) {
	q := 90.253
	z := fracBits
	fmt.Printf("%v,%v,%d,%x\n",
		q,
		q*float64(uint64(1)<<z),
		int64(q*float64(uint64(1)<<z)),
		int64(q*float64(uint64(1)<<z)),
	)
	fmt.Println(Fixed{int64(q*float64(uint64(1)<<z)) >> (z - fracBits)})
	fmt.Println(Fixed{int64(q*float64(uint64(1)<<z)) >> (z - fracBits)}.Float())
	fmt.Println(Fixed{0x5a40c49b}.Float())
}

func Test_InvE(t *testing.T) {
	q := 10.90051
	fmt.Printf("%v,%v,%d,%x\n",
		q,
		q*float64(uint64(1)<<54),
		int64(q*float64(uint64(1)<<54)),
		int64(q*float64(uint64(1)<<54)),
	)
	fmt.Println(Fixed{int64(q*float64(uint64(1)<<54)) >> (54 - fracBits)})
	fmt.Println(Fixed{int64(q*float64(uint64(1)<<54)) >> (54 - fracBits)}.Float())
}

func Test_Gr(t *testing.T) {
	q := 10.90051
	//q := 0.6207822
	fmt.Printf("%v,%v,%d,%x\n",
		q,
		q*float64(uint64(1)<<fracBits),
		int64(q*float64(uint64(1)<<fracBits)),
		int64(q*float64(uint64(1)<<fracBits)),
	)
	fmt.Println(Fixed{int64(q * float64(uint64(1)<<fracBits))}.Float())
}
