package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt ) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ",float64(e))
}

func Sqrt(x float64) (float64, error) {

	if(x < 0){
		return 0, ErrNegativeSqrt(x)
	}else{
		z := 1.0
    	eps := 2.220446049250313e-16
		for math.Abs((math.Pow(z,2) - x)/(2 * z)) > eps {
			z = z - (math.Pow(z,2) - x)/(2 * z)
		}
		return z, nil
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
