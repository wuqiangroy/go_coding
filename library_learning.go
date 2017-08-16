package main

import (
	"fmt"
	//"time"
	//"math/rand"
	//"math"
)

//func add(a int, b int) int{
//	return a + b
//}

func add(a, b int) int{
	return a % b
}

func str(a, b string) (string, string){
	return a+b, b
}

//return 裸返回, 返回所有值
func split(sum int) (a, b int){
	a = sum + 1
	b = sum - 2
	return
}

func main() {
	//fmt.Print("This time is:", time.Now())
	//运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字
	//var a = rand.Intn(10)
	//fmt.Print("\nA random number is: ", a)
	//fmt.Print(math.Pi)
	fmt.Print(add(2, 5))
	fmt.Print(str("\nthis is a ", "good thing\n"))
	fmt.Print(split(22))
	var j int = 10
	//函数外的每个语句都必须以关键字开始（ var 、 func 、等等）:= 结构不能使用在函数外。
	k := 20
	fmt.Print("\n", j, k)
	//变量在定义时没有明确的初始化时会赋值为 零值 。
	var a int
	var f float64
	var b bool
	var s string
	//注意，format格式化需要Printf
	fmt.Printf("\n%v, %v, %v, %q", a, f, b, s)
}
