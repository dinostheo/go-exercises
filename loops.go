package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z  := float64(1)
	st := float64(0)

	for i := 0; i < 100; i++ {
		z = z - (z * z - x) / (2 * z)

		if math.Abs(z - st) < 1e-10 {
			break
		}

		st = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}
