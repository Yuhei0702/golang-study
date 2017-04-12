package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
  eps := 2.220446049250313e-16
	for math.Abs(math.Pow(z,2) - x)/(2 * z) > eps {
			z = z - (math.Pow(z,2) - x)/(2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
