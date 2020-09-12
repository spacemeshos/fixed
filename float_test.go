package fixed

import (
	"testing"
)

func TestFixed_From(t *testing.T) {
	values := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	for _, v := range values {
		if got := from(v); got.float() != v {
			t.Errorf("from(%v) => %v|%v, want %v", v, got, got.float(), v)
		}
	}
}

func TestFixed_From56(t *testing.T) {
	values := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}
	for _, v := range values {
		got := from(v)
		//t.Logf("*from(%v) => %v|%v, float56: %v, want %v", v, got, got.float(), float56(got.fixed56()), v)
		if float56(got.fixed56()) != v {
			t.Errorf("from(%v) => %v|%v, float56: %v, want %v", v, got, got.float(), float56(got.fixed56()), v)
		}
	}
}
