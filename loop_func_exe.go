package main

import (
	"fmt"
	//"math"
)

func newton() float64 {
	i := 1
	//as C while loop
	for i < 11 {
		x := i - (i*i-i)/(2*i)
		fmt.Print("i is: ", i, "\n")
		fmt.Print("x is: ", x, "\n")
		i += 1
	}
	//for loop
	for i:= 1; i < 11; i++ {
		x := i - (i*i-i)/(2*i)
		fmt.Print(x, "\n")
	}
	return 1
}

func sqrt(n float64) float64 {
	for i := 1; i < 11; i ++ {
		n = n - (n*n - n) / (2*n)
	}
	return n
}

func main() {
	fmt.Print(newton())
	fmt.Print("\n10 times calculate is: ", sqrt(4))
}
