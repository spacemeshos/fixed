package fixed

import (
	"fmt"
	"testing"
)

func TestFixed_BinCDF(t *testing.T) {
	fmt.Printf("%.10f\n", BinCDF(New(34), From(0.9), New(25)).Float())
}

func TestFixed_BinCDF1(t *testing.T) {
	fmt.Printf("%.10f\n", BinCDF(New(340), From(0.82), New(250)).Float())
}
