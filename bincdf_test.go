package fixed

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFixed_BinCDF(t *testing.T) {
	fmt.Printf("%.10f\n", BinCDF(New(34), From(0.9), New(25)).Float())
}

func TestFixed_BinCDF1(t *testing.T) {
	fmt.Printf("%.10f\n", BinCDF(New(340), From(0.82), New(250)).Float())
}

func TestFixed_BinCDF2(t *testing.T) {
	require.NotPanics(t, func() {
		BinCDF(New(1471), From(.05), New(2))
	})
	require.NotPanics(t, func() {
		BinCDF(New(1472), From(.05), New(2))
	})
	require.PanicsWithValue(t, ErrOverflow, func() {
		BinCDF(New(6471), From(.05), New(2))
	})
	require.PanicsWithValue(t, ErrOverflow, func() {
		BinCDF(New(6472), From(.05), New(2))
	})
	require.PanicsWithValue(t, ErrOverflow, func() {
		BinCDF(New(20480), From(.05), New(2))
	})
	require.PanicsWithValue(t, ErrOverflow, func() {
		BinCDF(New(20481), From(.05), New(2))
	})
}
