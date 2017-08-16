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

func main() {
	//fmt.Print("This time is:", time.Now())
	//运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字
	//var a = rand.Intn(10)
	//fmt.Print("\nA random number is: ", a)
	//fmt.Print(math.Pi)
	fmt.Print(add(2, 5))
	fmt.Print(str("\nthis is a ", "good thing\n"))
}
