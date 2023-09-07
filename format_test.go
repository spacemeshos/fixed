package fixed

import "testing"

var formatCases = []struct {
	x float64
	s string
}{
	{
		x: 0.0,
		s: "+0'00000000000000/56",
	},
	{
		x: -0.0,
		s: "+0'00000000000000/56",
	},
	{
		x: 1.0,
		s: "+1'00000000000000/56",
	},
	{
		x: 1.25,
		s: "+1'40000000000000/56",
	},
	{
		x: 2.5,
		s: "+2'80000000000000/56",
	},
	{
		x: 0.984375,
		s: "+0'fc000000000000/56",
	},
	{
		x: -0.5,
		s: "-0'80000000000000/56",
	},
	{
		x: -4.125,
		s: "-4'20000000000000/56",
	},
	{
		x: -7.75,
		s: "-7'c0000000000000/56",
	},
}

func TestFixed_Format(t *testing.T) {
	for _, tc := range formatCases {
		if q := from(tc.x); q.format() != tc.s {
			t.Errorf("fromat(%v) => %v|%v, want %v|%v", tc.x, q.format(), q.float(), tc.s, tc.x)
		}
	}
}
