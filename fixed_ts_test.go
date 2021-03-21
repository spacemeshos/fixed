
package fixed
    
var testCases = []struct {
    x float64
    s string
    floor int
    round int
    ceil int
}{
{
    x: 0.0,
    s: "+0'00000000000000/56",
    floor:   0,
    round:  0.0,
    ceil:      0, 
},
{
    x: 1.0,
    s: "+1'00000000000000/56",
    floor:   1,
    round:  1.0,
    ceil:      1, 
},
{
    x: 1.25,
    s: "+1'40000000000000/56",
    floor:   1,
    round:  1.0,
    ceil:      2, 
},
{
    x: 2.5,
    s: "+2'80000000000000/56",
    floor:   2,
    round:  3.0,
    ceil:      3, 
},
{
    x: 0.984375,
    s: "+0'fc000000000000/56",
    floor:   0,
    round:  1.0,
    ceil:      1, 
},
{
    x: -0.5,
    s: "-0'80000000000000/56",
    floor:  -1,
    round: -1.0,
    ceil:      0, 
},
{
    x: -4.125,
    s: "-4'20000000000000/56",
    floor:  -5,
    round: -4.0,
    ceil:     -4, 
},
{
    x: -7.75,
    s: "-7'c0000000000000/56",
    floor:  -8,
    round: -8.0,
    ceil:     -7, 
},
}
