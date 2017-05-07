package main

import (
	"fmt"
	"math"
)

const eps = 1e-9

func Sqrt(x float64) float64 {
	z := 1.0
	z_tmp := z
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-z_tmp) < eps {
			break
		}
		z_tmp = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
