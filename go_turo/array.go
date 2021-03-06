package main

import (
	"fmt"
)

//array有长度slice没长度

func printSlice(s string, x []int) {
	//cap为数组容量， len为数组长度
	fmt.Printf("\n%s len=%d cap=%d %v",
	s, len(x), cap(x), x)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Print(a[0], a[1])
	fmt.Print("\n", a)
	n := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Print("\n", n)
	for i := 0; i < len(n); i++ {
		fmt.Print("\n", n[i])
	}
	//二维数组
	num := [][]string{
		[]string{"-", "--", "---"},
		[]string{"---", "--", "-"},
		[]string{"yes", "ppg", "just do it"},
	}
	fmt.Print("\n", num)
	num[0][2] = "right now"
	num[1][1] = "by yourself"
	num[2][0] = "left"
	fmt.Print("\n", num)
	fmt.Print("\n", n[0:3])
	fmt.Print("\n", n[:2])
	fmt.Print("\n", n[4:])

	//slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组
	a1 := make([]int, 5)
	printSlice("a1", a1)
	a2 := make([]int, 0, 5)
	printSlice("a2", a2)
	a3 := a2[:2]
	printSlice("a3", a3)
	a4 := a3[2:5]
	printSlice("a4", a4)

	//slice 的零值是 nil
	// 一个 nil 的 slice 的长度和容量是 0
	var a5 []int
	fmt.Print("\n", a5, len(a5), cap(a5))
	if a5 == nil {
		fmt.Print("\nnil")
	}

	//append用法
	var a6 []int
	a6 = append(a6, 1)
	printSlice("a6", a6)
	a6 = append(a6, 2, 3, 4)
	printSlice("a6", a6)
	a6 = append(a6, 100, 101, 102)
	printSlice("a6", a6)

	//for循环slice
	pow := []int{1, 2, 4, 5, 7, 8, 9, 10}
	for i, v := range pow {
		fmt.Printf("\n%d, %d", i, v)
	}
	//可以通过赋值给 _ 来忽略序号和值
	//pow2 := make([]int, 10)
	pow2 := []int{22, 33, 44, 55}
	for _, value := range pow2 {
		fmt.Print("\n", value)
	}
}
