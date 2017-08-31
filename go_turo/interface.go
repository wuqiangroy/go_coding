package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct{
	x, y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt(2))
	v := Vertex{5, 12}

	a = f
	fmt.Print(a, "\n")
	a = &v
	//a = v

	fmt.Print(a.Abs())
}