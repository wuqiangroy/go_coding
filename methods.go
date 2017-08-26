package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type MyFloat float64

//define method
func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (v *Vertex) Abs() float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := &Vertex{3, 4}
	fmt.Print(v.Abs(), "\n")
	m := MyFloat(-math.Sqrt(4))
	fmt.Print(m.Abs2(), "\n")
	fmt.Print("Before scaling: ", v.Abs(), "\n")
	v.Scale(5)
	fmt.Print("After scaling: ", v.Abs(), "\n")
	fmt.Print(v)
}

