import os
import math
import scipy.stats
import random

prec = 56
total = 128

def mul(a, b):
    z = int(a * (1 << prec)) * int(b * (1 << prec))
    return z >> prec


def div(a, b):
    a = int(a * (1 << prec))
    b = int(b * (1 << prec))
    z = (a * (1 << prec)) // b
    return z


def fixed(x):
    return int(x * (1 << prec))


def float64(x):
    return x / (1 << prec)


def ceil(x):
    return math.ceil(x)


def floor(x):
    return math.floor(x)


def round(x):
    return float(int(x + (0.5 if x > 0 else -0.5)))


def string(x):
    # formats x as fixed with precision prec
    #maxval = (1 << (63-prec)) - 1
    v = abs(x)
    #if (v >> prec) >= maxval:
    #    return "overflow"
    fmt = "%%s%%d'%%0%dx/%d" % (prec / 4, prec)
    return fmt % (("+", "-")[x < 0], v >> prec, v & ((1 << prec) - 1))


def xstring(a, b, v):
    #maxval = (1 << (63-prec)) - 1
    #if abs(a) >= maxval:
    #    return "overflow"
    #if abs(b) >= maxval:
    #    return "overflow"
    return string(v)


def gen_mul_tests():
    f = open('mul_ts_test.go1', 'w+')
    f.write('''
package fixed
    
var mulTestCases = []struct {
    x float64
    y float64
    z float64
    s string
}{''')

    def case(a, b):
        a = float(a)
        b = float(b)
        f.write('''
{{
    x: {:1},
    y: {:2},
    z: {:3},
    s: "{:4}",
}},'''.format(a, b, float64(mul(a, b)), xstring(a, b, mul(a, b))))

    case(0, 1.5)
    case(1.25, 4)
    case(1.25, -4)
    case(-1.25, 4)
    case(-1.25, -4)
    case(1.25, 1.5)
    case(1234.5, -8888.875)
    case(1.515625, 1.531250)
    case(0.500244140625, 0.500732421875)
    case(0.015625, 0.000244140625)
    case(1.44140625, 1.44140625)
    case(1.44140625, 1.441650390625)
    f.write('\n}\n')
    f.close()
    if os.path.isfile('mul_ts_test.go'):
        os.remove('mul_ts_test.go')
    os.rename('mul_ts_test.go1', 'mul_ts_test.go')


def gen_div_tests():
    f = open('div_ts_test.go1', 'w+')
    f.write('''
package fixed
    
var divTestCases = []struct {
    x int
    y int
    z float64
    s string
}{''')

    def case(a, b):
        a = float(a)
        b = float(b)
        f.write('''
{{
    x: {:1},
    y: {:2},
    z: {:3},
    s: "{:4}",
}},'''.format(int(a), int(b), a/b, xstring(a, b, div(a, b))))

    case(2, 3)
    case(1, 3)
    case(10, 7)
    case(18, 7)
    case(10,600)
    case(100000,600)
    f.write('\n}\n')
    f.close()
    if os.path.isfile('div_ts_test.go'):
        os.remove('div_ts_test.go')
    os.rename('div_ts_test.go1', 'div_ts_test.go')


def gen_fixed_tests():
    f = open('fixed_ts_test.go1', 'w+')
    f.write('''
package fixed
    
var testCases = []struct {
    x float64
    s string
    floor int
    round int
    ceil int
}{''')

    def case(a):
        a = float(a)
        f.write('''
{{
    x: {:1},
    s: "{:2}",
    floor: {:3},
    round: {:4},
    ceil:  {:5}, 
}},'''.format(a,string(fixed(a)), floor(a), round(a), ceil(a)))

    case(0)
    case(1)
    case(1.25)
    case(2.5)
    case(63/64)
    case(-0.5)
    case(-4.125)
    case(-7.75)

    f.write('\n}\n')
    f.close()
    if os.path.isfile('fixed_ts_test.go'):
        os.remove('fixed_ts_test.go')
    os.rename('fixed_ts_test.go1', 'fixed_ts_test.go')


def gen_bincdf_tests(A,B,e):

    fname = 'bincdf_ts{:1}_test.go'.format(B)
    tmpfname = fname+'1'

    f = open(tmpfname, 'w+')
    f.write('''
package fixed
import "testing"

var bincdfTestCases_{:1} = []BinCDFCase{{'''.format(B))

    def case(n,x,p):
        n = int(n)
        x = int(x)
        p = float(p)

        cdf = scipy.stats.binom.cdf(x, n, p)
        if abs(cdf) < 1e-15:
            cdf = 0
        if abs(1-cdf) < 1e-15:
            cdf = 1

        if cdf == 1 or cdf == 0:
            return

        #print(cdf, x,n,p)
        if math.isnan(cdf):
            return

        f.write('''
{{
    n: {:1},
    x: {:2},
    p: {:3},
    cdf: {:4},
    s: "{:5}",
}},'''.format(n,x,p,cdf,string(fixed(cdf))))

    def case2(n,x,p):
        n = int(n)
        x = int(x)
        p = float(p)

        cdf = scipy.stats.binom.cdf(x, n, p)
        if abs(cdf) < 1e-15:
            cdf = 0
        if abs(1-cdf) < 1e-15:
            cdf = 1

        if cdf == 1 or cdf == 0:
            return

        if math.isnan(cdf):
            return

        f.write('''
{{
    n: {:1},
    x: {:2},
    p: {:3},
    cdf: {:4},
    s: "{:5}",
}},'''.format(n,x,p,cdf,string(fixed(cdf))))

    random.seed(42)

    def q(n):
        x = n//10 + random.randint(0,n//10)
        case(n, x, 0.2 + random.random()*0.2 - 0.1)
        x = n//2 + random.randint(0,n//2-1)
        case(n, x, 0.5 + random.random()*0.2 - 0.1)
        x = n//3+ random.randint(0,n//3)
        case(n, x, 0.5 + random.random()*0.2 - 0.1)
        x = n//3*2 + random.randint(0,n//3-1)
        case(n, x, 0.5 + random.random()*0.2 - 0.1)
        x = n-n//10 - random.randint(0,n//10)
        case(n, x, 0.8 + random.random()*0.2 - 0.1)

    for k in range(15):
        n = random.randint(A,B)
        q(n)

    for t in range(1,50):
        for k in range(1,B,B//100):
            case2(B, k, t/100)

    f.write('\n}\n')

    f.write(('''
func Benchmark_Fixed_BinCDF_{}(b *testing.B) {{
	for i := 0; i < b.N; i++ {{
		tc := bincdfTestCases_{}[i%len(bincdfTestCases_{})]
		bincdfResultFix = BinCDF(tc.n, From(tc.p), tc.x)
	}}

	bincdfResultFix.lo++
}}

func Benchmark_Float_BinCDF_{}(b *testing.B) {{
	for i := 0; i < b.N; i++ {{
		tc := bincdfTestCases_{}[i%len(bincdfTestCases_{})]
		bincdfResultFlt = bincdf_(tc.n, tc.p, tc.x)
	}}

	bincdfResultFlt++
}}

func TestFixed_BinCDF_{}(t *testing.T) {{
	acc := accuracy{{Epsilon: '''+e+'''}}
	for i, tc := range bincdfTestCases_{} {{
		p := From(tc.p)
		got := BinCDF(tc.n, p, tc.x)
		if ok := acc.update(got, tc.cdf); !ok {{
			t.Errorf("%d: BinCDF(%v,%v,%v) => got %v|%v, want %v|%v", i, tc.n, tc.p, tc.x, got, got.Float(), From(tc.cdf), tc.cdf)
		}}
	}}
	t.Log(acc)
}}

''').format(B,B,B,B,B,B,B,B))

    f.close()
    if os.path.isfile(fname):
        os.remove(fname)
    os.rename(tmpfname, fname)


if __name__ == '__main__':
    gen_mul_tests()
    gen_div_tests()
    gen_fixed_tests()
    gen_bincdf_tests(10,100,'1e-14')
    gen_bincdf_tests(100,1000,'1e-12')
    gen_bincdf_tests(1000,10000,'1e-10')
    gen_bincdf_tests(1000,1000000,'1e-6')
    gen_bincdf_tests(100000,10000000,'1e-4')
