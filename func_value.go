package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Print(hypot(5, 12), "\n")

	fmt.Print(compute(hypot), "\n")
	fmt.Print(compute(math.Pow), "\n")
}
