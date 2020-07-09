// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fixed implements fixed-point integer types.
package fixed // import "github.com/spacemeshos/fixed"

import (
	"errors"
	"fmt"
	"math/bits"
)

var ErrOverflow = errors.New("overflow")

// TODO: implement fmt.Formatter for %f and %g.

// New creates Fixed from integer
func New(i int) Fixed {
	return Fixed{int64(i << fracBits)}
}

// From creates Fixed from float
func From(f float64) Fixed {
	return Fixed{int64(f * (1 << fracBits))}
}

// Raw creates Fixed from raw value
func Raw(i int64) Fixed {
	return Fixed{i}
}

// Float converts Fixed to float64
func (x Fixed) Float() float64 {
	return float64(x.int64) / (1 << fracBits)
}

// Fixed is a signed 52.12 fixed-point number.
//
// The integer part ranges from -2251799813685248 to 2251799813685247,
// inclusive. The fractional part has 12 bits of precision.
//
// For example, the number one-and-a-quarter is Fixed(1<<12 + 1<<10).
type Fixed struct {
	int64
}

const (
	// fracBits is the number of fractional bits. It cannot be more than half the total bits, otherwise the implementation
	// of Mul() can overflow in the fractional part multiplication.
	fracBits      int    = 12
	totalBits     int    = 64   // unsafe.Sizeof(Fixed(0)) * 8
	fracDecDigits int    = 4    // int(math.Log10(1<<fracBits)) + 1
	fracMask      int64  = 4095 // Fixed(1<<fracBits - 1)
	roundValue    uint64 = uint64(1) << (fracBits - 1)
	oneValue      int64  = int64(1) << fracBits
	carryMask     uint64 = ^(uint64(1)<<(64-fracBits) - 1)
)

var format = fmt.Sprintf("%%s%%d+%%0%dd/%d", fracDecDigits, 1<<fracBits)

// String returns a human-readable representation of a fixed-point number.
//
// For example, the number one-and-a-quarter becomes "1+1024/4096" (the divisor is 2^precision).
func (x Fixed) String() string {
	xs := int64(x.int64) >> 63
	v := (x.int64 ^ xs) - xs
	sign := "-"[:xs&1]
	return fmt.Sprintf(format, sign, v>>fracBits, v&fracMask)
}

// Floor returns the greatest integer value less than or equal to x.
//
// Its return type is int, not Fixed.
func (x Fixed) Floor() int {
	return int(x.int64 >> fracBits)
}

// Round returns the nearest integer value to x. Ties are rounded up.
//
// Its return type is int, not Fixed.
func (x Fixed) Round() int {
	return int((x.int64 + int64(roundValue)) >> fracBits)
}

// Ceil returns the least integer value greater than or equal to x.
//
// Its return type is int, not Fixed.
func (x Fixed) Ceil() int {
	return int((x.int64 + oneValue - 1) >> fracBits)
}

// Value returns interval value
func (x Fixed) Value() int64 {
	return x.int64
}

// Mul returns x*y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Mul(y Fixed) Fixed {
	hi, lo := bits.Mul64(uint64(x.int64), uint64(y.int64))
	hi = hi - uint64((x.int64>>63)&y.int64) - uint64((y.int64>>63)&x.int64)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	if carry != 0 || ((hi&carryMask) != 0 && (hi&carryMask) != carryMask) {
		panic(ErrOverflow)
	}
	return Fixed{int64(hi<<(64-fracBits) | (lo >> fracBits))}
}

// UnsafeMul returns x*y in fixed-point arithmetic.
// Does not have overflow check
func (x Fixed) UnsafeMul(y Fixed) Fixed {
	hi, lo := bits.Mul64(uint64(x.int64), uint64(y.int64))
	hi = hi - uint64((x.int64>>63)&y.int64) - uint64((y.int64>>63)&x.int64)
	lo, carry := bits.Add64(lo, roundValue, 0)
	hi, _ = bits.Add64(hi, 0, carry)
	return Fixed{int64(hi<<(64-fracBits) | (lo >> fracBits))}
}

// Div returns x/y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Div(y Fixed) Fixed {
	xs := x.int64 >> 63
	ys := y.int64 >> 63
	a := uint64((x.int64 ^ xs) - xs)        // abs
	b := uint64((y.int64 ^ ys) - ys)        // abs
	hi, lo := a>>(64-fracBits), a<<fracBits // а*frac
	// will panic when divides by zero or occurs overflow
	v, rem := bits.Div64(hi, lo, b)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ [0,1]
	// round to near
	v, carry := bits.Add64(v, (rem<<1)/b, 0)
	if carry != 0 {
		panic(ErrOverflow)
	}
	return Fixed{int64(v) * ((xs^ys)*2 + 1)}
}

// UsafeDiv returns x/y in fixed-point arithmetic.
// Does not have overflow check (but bits.Div64 has it's own)
func (x Fixed) UnsafeDiv(y Fixed) Fixed {
	xs := x.int64 >> 63
	ys := y.int64 >> 63
	a := uint64((x.int64 ^ xs) - xs)        // abs
	b := uint64((y.int64 ^ ys) - ys)        // abs
	hi, lo := a>>(64-fracBits), a<<fracBits // а*frac
	// will panic when divides by zero or occurs overflow
	v, rem := bits.Div64(hi, lo, b)
	// rem < b && (b>>63) == 0 => (rem<<1) < ^uint64(0)
	//                            (rem<<1)/b ∈ [0,1]
	// round to near
	v, _ = bits.Add64(v, (rem<<1)/b, 0)
	return Fixed{int64(v) * ((xs^ys)*2 + 1)}
}

// Add returns x+y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Add(y Fixed) Fixed {
	v := x.int64 + y.int64
	if x.int64>>63 == y.int64>>63 && x.int64>>63 != int64(v)>>63 {
		panic(ErrOverflow)
	}
	return Fixed{int64(v)}
}

// UnsafeAdd returns x+y in fixed-point arithmetic.
// Does not have overflow check
func (x Fixed) UnsafeAdd(y Fixed) Fixed {
	return Fixed{x.int64 + y.int64}
}

// Sub returns x-y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Sub(y Fixed) Fixed {
	v := x.int64 - y.int64
	if x.int64>>63 != y.int64>>63 && x.int64>>63 != int64(v)>>63 {
		panic(ErrOverflow)
	}
	return Fixed{int64(v)}
}

// UnsafeSub returns x-y in fixed-point arithmetic.
// Does not have overflow check
func (x Fixed) UnsafeSub(y Fixed) Fixed {
	return Fixed{x.int64 - y.int64}
}
