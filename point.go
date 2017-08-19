package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	//&符号会生成一个指向其作用对象的指针
	p := &i
	fmt.Print(p, "\n")
	*p = 21
	fmt.Print(i, "\n")

	p = &j
	*p = *p / 37
	fmt.Print(j)
}
