package fixed

import "math/bits"

type fixed128 struct {
	hi int64 // high integer bits
	lo int64 // fixed value
}

func(a fixed128) fixed() int64 {
	if a.hi != 0 && a.hi != -1 || a.hi >> 63 != a.lo >> 63 {
		panic(ErrOverflow)
	}
	return a.lo
}

func tofixed128(x int64) fixed128 {
	return fixed128{ x>>63, x }
}

func add128(a ...fixed128) fixed128 {
	c := a[0]
	for _,x := range a[1:] {
		lo, curry := bits.Add64(uint64(x.lo),uint64(c.lo),0)
		hi, _ := bits.Add64(uint64(x.hi), uint64(c.hi), curry)
		if x.hi>>63 == c.hi>>63 && x.hi>>63 != int64(hi)>>63 {
			panic(ErrOverflow)
		}
		c = fixed128{ int64(hi), int64(lo) }
	}
	return c
}

func sub128(a ...fixed128) fixed128 {
	c := a[0]
	for _,x := range a[1:] {
		lo, borrow := bits.Sub64(uint64(c.lo), uint64(x.lo), 0)
		hi, _ := bits.Sub64(uint64(c.hi),uint64(x.hi),borrow)
		c = fixed128{ int64(hi), int64(lo) }
	}
	return c
}

func mul128_56(x, y int64) fixed128 {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	hi = hi - uint64((x>>63)&y) - uint64((y>>63)&x)
	lo, carry := bits.Add64(lo, roundValue56, 0)
	hi, carry = bits.Add64(hi, 0, carry)
	return fixed128{ int64(hi)>>56, int64(hi<<8 | lo>>56) }
}

