package fixed

import (
	"math"
	"math/bits"
)

func (x Fixed) float() float64 {
	l := x.bitlen()
	if l > 53 {
		x = x.shr(l - 53)
	} else {
		x.lo = x.lo << (53 - l)
	}
	a := x.lo
	if a != 0 {
		a &= ^(uint64(1) << 52)
		a |= uint64(1023+l-57) << 52
		a |= x.sign_()
	}
	f := math.Float64frombits(a)
	return f
}

func from(f float64) Fixed {
	a := math.Float64bits(f)

	if a != 0 {
		v := a & ((uint64(1) << 52) - 1)
		e := int((a>>52)&((1<<11)-1)) - 1023 - 52 + 56
		v |= uint64(1) << 52 // MSB
		s := a & signMask
		if e < 0 {
			e = -e
			return fixed_(v>>e, s)
		}
		if bits.Len64(v)+e > 127 {
			panic(ErrOverflow)
		}
		return fixed_(v, s).shl(e)
	}

	return Fixed{}
}
