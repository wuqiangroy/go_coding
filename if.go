package main
import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprintln(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//if语句定义的变量的作用域仅在 if 范围之内。
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//以下用v则不行
	return lim
}

func main() {
	//fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(2, 3, 10))
	fmt.Println(pow(2, 4, 10))
}
