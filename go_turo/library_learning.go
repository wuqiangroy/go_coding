package main

import (
	"fmt"
	//"time"
	//"math/rand"
	"math"
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

func trans() (int, int, float64, uint){
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	return x, y, f, z
}

const (
	Big = 1 << 100
	Small = Big >>99
)

func needInt(x int) int {
	return x*10 +1
}

func needFloadt(x float64) float64 {
	return x * 0.1
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
	//var a int
	//var f float64
	//var b bool
	//var s string
	//注意，format格式化需要Printf
	//fmt.Printf("\n%v, %v, %v, %q", a, f, b, s)

	var i int = 42
	var fl float64 = float64(i)
	var u uint = uint(fl)
	//i := 42
	//fl := float64(i)
	//u := uint(fl)
	fmt.Printf("\ni: %v, fl: %v, u: %v\n", i, fl, u)
	fmt.Print(trans())
	//print Type
	var t float64 = 324
	t = 22.3
	fmt.Printf("\ntype t is %T\n", t)
	const p = "你好，世界"
	const Truth = true
	const wrong = false
	if wrong {
		fmt.Print(p, Truth, wrong)
	}
	fmt.Print("\n", needInt(Small))
	fmt.Print("\n", needFloadt(Small))
	fmt.Print("\n", needFloadt(Big))
}
