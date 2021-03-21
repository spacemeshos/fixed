package fixed

import (
	"encoding/binary"
	"errors"
)

var ErrOverflow = errors.New("overflow")

func (x Fixed) String() string {
	return x.format()
}

func New(val int) Fixed {
	return fixed(int64(val))
}

func New64(val int64) Fixed {
	return fixed(val)
}

func From(val float64) Fixed {
	return from(val)
}

var One = fixedOne
var Zero = Fixed{}

func (x Fixed) Abs() Fixed {
	return x.abs()
}

func (x Fixed) Neg() Fixed {
	return x.neg()
}

func (x Fixed) Floor() int64 {
	return x.floor()
}

func (x Fixed) Ceil() int64 {
	return x.ceil()
}

func (x Fixed) Round() int64 {
	return x.round()
}

func (x Fixed) Float() float64 {
	return x.float()
}

func (x Fixed) Mul(y Fixed) Fixed {
	return mul(x, y)
}

func (x Fixed) Div(y Fixed) Fixed {
	return div(x, y)
}

func (x Fixed) Add(y Fixed) Fixed {
	return add(x, y)
}

func (x Fixed) Sub(y Fixed) Fixed {
	return sub(x, y)
}

func (x Fixed) LessThan(y Fixed) bool {
	return x.less(y)
}

// GreaterThan compares fixed values and returns true if x > y
func (x Fixed) GreaterThan(y Fixed) bool {
	return x.greater(y)
}

// EqualTo compares fixed values and returns true if x == y
func (x Fixed) EqualTo(y Fixed) bool {
	return x.equal(y)
}

func DivUint64(p, q uint64) Fixed {
	return udiv(ufixed(p), ufixed(q))
}

// Div64 creates new Fixed equal to p/q signed result
func Div64(p, q int64) Fixed {
	return div(fixed(p), fixed(q))
}

// FracFromBytes takes only fractional part from bytes array and return fixed value
func FracFromBytes(x []byte) Fixed {
	return rawfixed(int64(binary.LittleEndian.Uint64(x)) & fracMask)
}

// FromBytes creates fixed value from bytes array
func FromBytes(x []byte) Fixed {
	return Fixed{lo: binary.LittleEndian.Uint64(x[:8]), hi: binary.LittleEndian.Uint64(x[8:])}
}

// Bytes converts fixed value into bytes array
func (x Fixed) Bytes() []byte {
	b := [16]byte{}
	binary.LittleEndian.PutUint64(b[:8], x.lo)
	binary.LittleEndian.PutUint64(b[8:], x.hi)
	return b[:]
}

func BinCDF(n int, p Fixed, x int) Fixed {
	if x < 0 {
		return Zero
	} else if x >= n {
		return One
	} else {
		return incomplete(int64(n-x), int64(x+1), oneValue-p.fixed56())
	}
}

func BinCDF64(n int64, p Fixed, x int64) Fixed {
	if x < 0 {
		return Zero
	} else if x >= n {
		return One
	} else {
		return incomplete(n-x, x+1, oneValue-p.fixed56())
	}
}
