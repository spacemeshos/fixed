### Fixed-point math library

The library implements basic math operations for 72/56 fixed-point values, and Log, Exp, Lgamma, BinCDF functions.

### Motivation

There is no efficient golang implementation for fixed-point math with a high-precision bincdf function.

The library uses a 56-bit fractional part to achieve maximum performance on math functions like Log, Exp, and Lgamma.

### Implementation

[BinCDF](https://github.com/spacemeshos/fixed/blob/master/fixed.go#L110) is implemented with 
[incomplete Beta function](https://github.com/spacemeshos/fixed/blob/master/beta.go#L5) in standart way:   
```  
Iₓ(a,b) = (xᵃ*(1-x)ᵇ)/(a*B(a,b)) * (1/(1+(d₁/(1+(d₂/(1+...))))))   
(xᵃ*(1-x)ᵇ)/B(a,b) = exp(lgamma(a+b) - lgamma(a) - lgamma(b) + a*log(x) + b*log(1-x))   
d_{2m+1} = -(a+m)(a+b+m)x/((a+2m)(a+2m+1))   
d_{2m}   = m(b-m)x/((a+2m-1)(a+2m))   
```
Fixed-point arithmetics and 
[Log](https://github.com/spacemeshos/fixed/blob/master/log.go#L19),
[Exp](https://github.com/spacemeshos/fixed/blob/master/exp.go#L10),
[Lgamma](https://github.com/spacemeshos/fixed/blob/master/lgamma.go#L4)
functions are implemented with the **bits** go module and raw operations over 56-bit fractional part values in range [-7..7]. 
