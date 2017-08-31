package main

import (
	"fmt"
)

//this can use string to describe self type
type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Nell", 43,}
	b := Person{"Roy", 24,}
	fmt.Print(a, b)

}
