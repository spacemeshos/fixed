// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fixed implements fixed-point integer types.
package fixed // import "github.com/spacemeshos/fixed/40_24"

import (
	"fmt"
	"math"
	"unsafe"
)

// TODO: implement fmt.Formatter for %f and %g.

const fracBits = 24
// fracBits cannot be more than half the total bits, otherwise the implementation of Mul() can overflow in the
// fractional part multiplication.

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

// Fixed is a signed 40.24 fixed-point number.
//
// The integer part ranges from -549755813888 to 549755813889,
// inclusive. The fractional part has 24 bits of precision.
//
// For example, the number one-and-a-quarter is Fixed(1<<24 + 1<<22).
type Fixed int64

// String returns a human-readable representation of a fixed-point number.
//
// For example, the number one-and-a-quarter becomes "1+4194304/16777216" (the divisor is 2^precision).
func (x Fixed) String() string {
	const shift, mask = fracBits, 1<<fracBits - 1
	format := fmt.Sprintf("%%d+%%0%dd/%d", fracDecDigits, 1<<fracBits)
	if x >= 0 {
		return fmt.Sprintf(format, x>>shift, x&mask)
	}
	x = -x
	if x >= 0 {
		return fmt.Sprintf("-"+format, x>>shift, x&mask)
	}
	return fmt.Sprintf("-"+format, -(1 << (totalBits - fracBits - 1)), 0) // This is the minimum value.
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

// Div returns x/y in fixed-point arithmetic.
func (x Fixed) Div(y Fixed) Fixed {
	ret := x / y << fracBits              // Calculate whole part.
	ret += x % y << fracBits / y          // Add fractional part.
	ret += x % y << fracBits % y << 1 / y // Round to nearest, instead of rounding down.
	return ret
}
