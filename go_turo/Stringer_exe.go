package main

import "fmt"

type IPAddr struct {
	a, b, c, d int
}

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i.a, i.b, i.c, i.d)
}

func main() {
	address := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8,8,8,8},
	}
	for n, a := range address {
		fmt.Printf("%v:%v\n", n, a)
	}
}