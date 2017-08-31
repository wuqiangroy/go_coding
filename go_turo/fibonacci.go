package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	//递归
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Print(fibonacci(4), "\n")
	for i := 1; i < 10; i++ {
		fmt.Print(fibonacci(i), "\n")
	}
}
