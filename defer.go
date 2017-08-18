package main

import (
	"fmt"
)

func main() {
	//延迟函数的执行直到上层函数返回
	//延迟调用的参数会立刻生成，但是不会在上层函数返回前执行
	//defer fmt.Print("World")
	//fmt.Print("Hello, ")

	fmt.Print("Counting\n")
	for i := 1; i < 10; i++ {
		defer fmt.Print(i, "\n")
	}
	fmt.Print("Done\n")
}
