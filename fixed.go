// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fixed implements fixed-point integer types.
package fixed // import "github.com/spacemeshos/fixed"
import "encoding/binary"

// Fixed is a signed fixed-point number.
type Fixed struct {
	int64
}

// New creates Fixed from integer
func New(i int) Fixed {
	return Fixed{fixed(i)}
}

// From creates Fixed from float
func From(f float64) Fixed {
	return Fixed{from(f)}
}

// Raw creates Fixed from raw value
func Raw(i int64) Fixed {
	return Fixed{i}
}

// DivUint64 creates new Fixed equal to p/q unsigned result
func DivUint64(p, q uint64) Fixed {
	return Fixed{udiv(p, q)}
}

// Div64 creates new Fixed equal to p/q signed result
func Div64(p, q int64) Fixed {
	return Fixed{div(p, q)}
}

// FracFromBytes takes only fractional part from bytes array and return fixed value
func FracFromBytes(x []byte) Fixed {
	return Fixed{int64(binary.LittleEndian.Uint64(x)) & fracMask}
}

// FromBytes creates fixed value from bytes array
func FromBytes(x []byte) Fixed {
	return Fixed{int64(binary.LittleEndian.Uint64(x))}
}

// Bytes converts fixed value into bytes array
func (x Fixed) Bytes() []byte {
	b := [8]byte{}
	binary.LittleEndian.PutUint64(b[:], uint64(x.int64))
	return b[:]
}

// Float converts Fixed to float64
func (x Fixed) Float() float64 {
	return float(x.int64)
}

var One = Fixed{oneValue}

// Neg inverts sign
func (x Fixed) Neg() Fixed {
	return Fixed{-x.int64}
}

// String returns a human-readable representation of a fixed-point number.
func (x Fixed) String() string {
	return format(x.int64)
}

// Floor returns the greatest integer value less than or equal to x.
//
// Its return type is int, not Fixed.
func (x Fixed) Floor() int {
	return integer(x.int64)
}

// Round returns the nearest integer value to x. Ties are rounded up.
//
// Its return type is int, not Fixed.
func (x Fixed) Round() int {
	return integer(round(x.int64))
}

// Ceil returns the least integer value greater than or equal to x.
//
// Its return type is int, not Fixed.
func (x Fixed) Ceil() int {
	return integer(ceil(x.int64))
}

// Value returns interval value
func (x Fixed) Value() int64 {
	return x.int64
}

// Add returns x+y in fixed-point arithmetic.
// Panics if overflows
func (x Fixed) Add(y Fixed) Fixed {
	v := x.int64 + y.int64
	if x.int64>>63 == y.int64>>63 && x.int64>>63 != v>>63 {
		panic(ErrOverflow)
	}
	return Fixed{v}
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
	if x.int64>>63 != y.int64>>63 && x.int64>>63 != v>>63 {
		panic(ErrOverflow)
	}
	return Fixed{v}
}

// UnsafeSub returns x-y in fixed-point arithmetic.
// Does not have overflow check
func (x Fixed) UnsafeSub(y Fixed) Fixed {
	return Fixed{x.int64 - y.int64}
}

// Abs returns absolute value of the fixed-point number
func (x Fixed) Abs() Fixed {
	return Fixed{abs(x.int64)}
}

// LessThan compares fixed values and returns true if x < y
func (x Fixed) LessThan(y Fixed) bool {
	return x.int64 < y.int64
}

// GreaterThan compares fixed values abd returns true if x < y
func (x Fixed) GreaterThan(y Fixed) bool {
	return x.int64 > y.int64
}

// EqualTo compares fixed values and returns true if x == y
func (x Fixed) EqualTo(y Fixed) bool {
	return x.int64 == y.int64
}
