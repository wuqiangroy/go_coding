package main

import (
	"fmt"
	"os"
	"strings"
)
//how to run this program
//as python, go run command_line.go "first" "second"
//return: first second 

func main() {
	var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = "  "
	// }

	//  _ is index and args is value : os.Args[index]
	//  _ means blank indentifier
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = "   "
	}
	fmt.Println(s)
	// another way to print os.Args
	fmt.Println(strings.Join(os.Args[1:], "   "))
}
