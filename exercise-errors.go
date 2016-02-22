package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %s", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (float64, error)  {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		newz := z - (z * z - x) / (2 * z)
		if math.Abs(newz - z) < 1e-10 {
			break
		} else {
			z = newz
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
> go run exercise-errors.go
1.4142135623746899 <nil>
0 cannot sqrt negative number: -2
*/