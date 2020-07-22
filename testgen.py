import os
import math

precisions = [12, 24, 32, 40, 48, 52]


def mul(a, b, prec):
    z = int(a * (1 << prec)) * int(b * (1 << prec))
    return ((z >> (prec-1)) + 1) >> 1


def div(a, b, prec):
    a = int(a * (1 << prec))
    b = int(b * (1 << prec))
    z = (a * (1 << prec) * 2) // b
    return (z + 1) >> 1


def fixed(x, prec):
    return int(x * (1 << prec))


def float64(x, prec):
    return x / (1 << prec)


def ceil(x):
    return math.ceil(x)


def floor(x):
    return math.floor(x)


def round(x):
    return float(int(x + (0.5 if x > -1 else -0.5)))


def string(x, prec):
    # formats x as fixed with precision prec
    maxval = (1 << (63-prec)) - 1
    v = abs(x)
    if (v >> prec) >= maxval:
        return "overflow"
    fmt = "%%s%%d+%%0%dx/%d" % (prec / 4, prec)
    return fmt % (("", "-")[x < 0], v >> prec, v & ((1 << prec) - 1))


def xstring(a, b, v, prec):
    maxval = (1 << (63-prec)) - 1
    if abs(a) >= maxval:
        return "overflow"
    if abs(b) >= maxval:
        return "overflow"
    return string(v, prec)


def gen_mul_tests():
    f = open('mul_ts_test.go1', 'w+')
    f.write('''
package fixed
    
var mulTestCases = []struct {
    x float64
    y float64
    z map[string]float64
    s map[string]string
}{''')

    def case(a, b):
        a = float(a)
        b = float(b)
        f.write('''
{{
    x: {:1},
    y: {:2},
    z: map[string]float64{{'''.format(a, b))
        for i in precisions:
            f.write('''
        "{:1}_{:2}":{:3},'''.format(64 - i, i, float64(mul(a, b, i), i)))
        f.write('''
    },
    s: map[string]string{''')
        for i in precisions:
            f.write('''
        "{:1}_{:2}":"{:3}",'''.format(64 - i, i, xstring(a, b, mul(a, b, i), i)))
        f.write('''
    },
},''')

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
    f.write('}\n')
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
    s map[string]string
}{''')

    def case(a, b):
        a = float(a)
        b = float(b)
        f.write('''
{{
    x: {:1},
    y: {:2},
    z: {:3},'''.format(a, b, a/b))
        f.write('''
    s: map[string]string{''')
        for i in precisions:
            f.write('''
        "{:1}_{:2}":"{:3}",'''.format(64 - i, i, xstring(a, b, div(a, b, i), i)))
        f.write('''
    },
},''')

    case(2, 3)
    case(1, 3)
    case(10, 7)
    case(18, 7)
    for i in precisions:
        case((1 << (64-i-1))-1, 31)
    f.write('}\n')
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
    s map[string]string
    floor int
    round int
    ceil int
}{''')

    def case(a):
        a = float(a)
        f.write('''
{{
    x: {:1},'''.format(a))
        f.write('''
    s: map[string]string{''')
        for i in precisions:
            f.write('''
        "{:1}_{:2}":"{:3}",'''.format(64 - i, i, string(fixed(a, i), i)))
        f.write('''
    }},
    floor: {:1},
    round: {:1},
    ceil:  {:1}, 
}},'''.format(floor(a), round(a), ceil(a)))

    case(0)
    case(1)
    case(1.25)
    case(2.5)
    case(63/64)
    case(-0.5)
    case(-4.125)
    case(-7.75)

    for i in precisions:
        case((1 << (64-i-1))-1/31)
    f.write('}\n')
    f.close()
    if os.path.isfile('fixed_ts_test.go'):
        os.remove('fixed_ts_test.go')
    os.rename('fixed_ts_test.go1', 'fixed_ts_test.go')


if __name__ == '__main__':
    gen_mul_tests()
    gen_div_tests()
    gen_fixed_tests()
