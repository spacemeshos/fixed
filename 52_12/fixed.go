// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fixed implements fixed-point integer types.
package fixed // import "github.com/spacemeshos/fixed/52_12"

import (
	"fmt"
	"math"
	"math/bits"
	"unsafe"
)

// TODO: implement fmt.Formatter for %f and %g.

// fracBits is the number of fractional bits. It cannot be more than half the total bits, otherwise the implementation
// of Mul() can overflow in the fractional part multiplication.
const fracBits = 12

var (
	totalBits     = unsafe.Sizeof(Fixed(0)) * 8
	fracDecDigits = int(math.Log10(1<<fracBits)) + 1
	fracMask      = Fixed(1<<fracBits - 1)
)

// I returns the integer value i as a Fixed.
//
// For example, if the precision is 6, passing the integer value 2 yields Fixed(128) because 2<<6=128.
func I(i int) Fixed {
	return Fixed(i << fracBits)
}

// Fixed is a signed 52.12 fixed-point number.
//
// The integer part ranges from -2251799813685248 to 2251799813685247,
// inclusive. The fractional part has 12 bits of precision.
//
// For example, the number one-and-a-quarter is Fixed(1<<12 + 1<<10).
type Fixed int64

// String returns a human-readable representation of a fixed-point number.
//
// For example, the number one-and-a-quarter becomes "1+1024/4096" (the divisor is 2^precision).
func (x Fixed) String() string {
	const shift, mask = fracBits, 1<<fracBits - 1
	format := fmt.Sprintf("%%d+%%0%dd/%d", fracDecDigits, 1<<fracBits) // e.g. "%d+%04d/4096" for 12 bit fractions
	return fmt.Sprintf(format, x>>shift, x&mask)
	// TODO: Improve negative number representation.
	//  E.g. With 6 bit fractions, -1/2 is printed as -1+32/64, but 0-32/64 would be more intuitive.
}

// Floor returns the greatest integer value less than or equal to x.
//
// Its return type is int, not Fixed.
func (x Fixed) Floor() int { return int(x >> fracBits) }

// Round returns the nearest integer value to x. Ties are rounded up.
//
// Its return type is int, not Fixed.
func (x Fixed) Round() int { return int((x + 1<<(fracBits-1)) >> fracBits) }

// Ceil returns the least integer value greater than or equal to x.
//
// Its return type is int, not Fixed.
func (x Fixed) Ceil() int { return int((x + (1 << fracBits) - 1) >> fracBits) }

// Mul returns x*y in fixed-point arithmetic.
func (x Fixed) Mul(y Fixed) Fixed {
	xw, yw := x>>fracBits, y>>fracBits     // Whole part of x and y.
	xf, yf := x&fracMask, y&fracMask       // Fractional part of x and y.
	ret := (xw * yw) << fracBits           // Multiply whole part.
	ret += xw*yf + yw*xf                   // Multiply each whole by other fraction.
	ret += (xf * yf) >> fracBits           // Multiply fractions by each other.
	ret += (xf * yf) >> (fracBits - 1) & 1 // Round to nearest, instead of rounding down.
	return ret
}

func (x Fixed) Mul2(y Fixed) Fixed {
	xs := x >> (totalBits - 1)                                 // x >= 0 ? 0 : -1
	ys := y >> (totalBits - 1)                                 // y >= 0 ? 0 : -1
	sign := 1 + 2*(xs^ys)                                      // x*y >= 0 ? 1 : -1
	hi, lo := bits.Mul64(uint64((x^xs)-xs), uint64((y^ys)-ys)) // abs(x): (x^xs)-xs
	rounding := (lo >> (fracBits - 1)) & 1
	println(x.String(), y.String(), rounding, hi, lo, uint64((x^xs)-xs), uint64((y^ys)-ys))
	return sign * Fixed(hi<<(totalBits-fracBits)|lo>>fracBits+rounding)
}

// Div returns x/y in fixed-point arithmetic.
func (x Fixed) Div(y Fixed) Fixed {
	ret := x / y << fracBits              // Calculate whole part.
	ret += x % y << fracBits / y          // Add fractional part.
	ret += x % y << fracBits % y << 1 / y // Round to nearest, instead of rounding down.
	return ret
}
