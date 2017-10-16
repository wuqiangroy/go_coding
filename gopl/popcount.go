package main

import "fmt"

var pc [256]byte

// like def __init__() in Python
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]+
		pc[byte(x>>(1*8))]+
		pc[byte(x>>(2*8))]+
		pc[byte(x>>(3*8))]+
		pc[byte(x>>(4*8))]+
		pc[byte(x>>(5*8))]+
		pc[byte(x>>(6*8))]+
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println(len(pc))
	fmt.Println(pc)
	fmt.Println(PopCount(323232))
}