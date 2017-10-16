package main

import "fmt"

func main() {
	// first define x 
	x := "hello!"
	for i := 0; i < len(x); i++ {
		// define a new x second time
		x := x[i]
		if x != '!' {
			// define a new x the third time
			x := x +'A' - 'a'
			fmt.Println(x)
			fmt.Println('a')
		}
	}
}
