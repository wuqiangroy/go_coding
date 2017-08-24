package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//func closure
//函数 adder 返回一个闭包。每个返回的闭包都被绑定到其各自的 sum 变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//closure to answer fibonacci
func fibonacci() func(int) int {
	x, y := 0, 1
	return func(i int) int {
		x, y = y, x+y
		return y
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Print(hypot(5, 12), "\n")

	fmt.Print(compute(hypot), "\n")
	fmt.Print(compute(math.Pow), "\n")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i ++ {
		fmt.Print(
			pos(i),
			neg(-2*i), "\n")
	}
	f := fibonacci()
	for i := 0; i < 10; i ++ {
		fmt.Print(f(i), "\n")
	}
}
