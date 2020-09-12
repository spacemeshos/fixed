package fixed

import "fmt"

var formatString = fmt.Sprintf("%%v%%d'%%0%dx/%d", fracBits/4, fracBits)

func (x Fixed) format() string {
	s := []string{"+", "-"}[x.sign_()>>63]
	a := x.abs()
	return fmt.Sprintf(formatString, s, a.integer(), a.fraction())
}
