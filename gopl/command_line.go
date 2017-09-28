package main

import (
	"fmt"
	"os"
)
//how to run this program
//as python, go run command_line.go "first" "second"
//return: first second 

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "  "
	}
	fmt.Println(s)
}
