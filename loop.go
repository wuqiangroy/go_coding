package main

import (
	"fmt"
)

func main() {
	sum := 0
	//for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Print(sum)
	sum2 := 1
	for ; sum2 < 100; {
		sum2 += sum2
	}
	fmt.Print("\n", sum2)
	
	//as while in C
	sum3 := 0
	for sum3 < 1000 {
		sum3 += 1
	}
	fmt.Print("\n", sum3)
}