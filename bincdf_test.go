package fixed

import (
	"fmt"
	"testing"
)

func TestFixed_BinCDF(t *testing.T) {
	fmt.Printf("%.8f\n", BinCDF(New(34), Frac(9, 1), New(25)).Float())
}
