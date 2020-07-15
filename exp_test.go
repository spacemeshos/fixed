package fixed

import (
	"fmt"
	"math"
	"testing"
)

var expEpsilon = 1 / (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>1)))))

func TestRange_NegExp(t *testing.T) {
	for i := -26 << fracBits; i <= 0; i += 1024 {
		a := Fixed{int64(i)}
		y := Exp(a)
		got := y.Float()
		want := math.Exp(Fixed{a.int64}.Float())
		if got < want-expEpsilon || got > want+expEpsilon {
			t.Errorf("exp(%v) => %v: got %v, want %v", a.Float(), y, got, want)
			t.FailNow()
		}
	}
}

func TestRange_Exp(t *testing.T) {
	for i := 0; i < 27<<fracBits; i += 1024 {
		a := Fixed{int64(i)}
		y := Exp(a)
		got := y.Float()
		want := math.Exp(Fixed{a.int64}.Float())
		expEpsilon = got / (math.Pow(10, math.Floor(math.Log10(float64(oneValue>>1)))))
		//t.Logf("exp(%v|%v) => %v: got %v, want %v", a, a.Float(), y, got, want)
		if got <= want-expEpsilon || got >= want+expEpsilon {
			t.Errorf("exp(%v|%v) => %v: got %v, want %v | %v", a, a.Float(), y, got, want, From(want))
			t.FailNow()
		}
	}
}

func TestRange_Exp27(t *testing.T) {
	func() {
		defer func() { recover() }()
		x := New(27)
		Exp(x)
		t.Errorf("exp(%v) must panic by overflow", x)
	}()
}

func BenchmarkFixed_Exp(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Exp(Fixed{int64(i % 130000)})
	}
}

func BenchmarkFixed_ExpRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Exp(Fixed{int64(i % 130000)}.Float()))
	}
}

func BenchmarkFixed_NegExp(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = Exp(Fixed{-int64(i % 130000)})
	}
}

func BenchmarkFixed_NegExpRef(b *testing.B) {
	for i := 1; i < b.N+1; i++ {
		Result = From(math.Exp(Fixed{-int64(i % 130000)}.Float()))
	}
}

func BenchmarkFixed_FullExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = Exp(Fixed{int64(i)*(144000*2)/int64(b.N) - 144000})
	}
}

func BenchmarkFixed_FullExpRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = From(math.Exp(Fixed{int64(i)*(144000*2)/int64(b.N) - 144000}.Float()))
	}
}

func Test_ExpGen(t *testing.T) {
	for i := 9; i >= 0; i-- {
		q := math.Exp(-float64(int64(1) << i))
		f, e := math.Frexp(q)
		fmt.Printf("{0x%x, %4d}, // exp(-2^%d) => %v / 1<<%v\n", int64(f*float64(uint64(1)<<63)), -e, i, f, -e)
	}
	q := math.Exp(0)
	f, e := math.Frexp(q)
	fmt.Printf("{0x%x, %4d}, // exp(0) => %v * 1<<%v\n", int64(f*float64(uint64(1)<<63)), e, f, e)
	for i := 0; i <= 9; i++ {
		q := math.Exp(float64(int64(1) << i))
		f, e := math.Frexp(q)
		fmt.Printf("{0x%x, %4d}, // exp(2^%d) => %v * 1<<%v\n", int64(f*float64(uint64(1)<<63)), e, i, f, e)
	}

	for i := 1; i <= 9; i++ {
		q := math.Exp(float64(1 / int64(1) << i))
		f, e := math.Frexp(q)
		fmt.Printf("{0x%x, %4d}, // exp(2^%d) => %v * 1<<%v\n", int64(f*float64(uint64(1)<<63)), e, i, f, e)
	}

}
