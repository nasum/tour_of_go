package main

import (
	"fmt"
	"math"
)

const eps = 1e-9

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	z_tmp := z
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-z_tmp) < eps {
			break
		}
		z_tmp = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
