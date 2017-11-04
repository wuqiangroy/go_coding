package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Cycle struct {
	Point
	radius int 
}

// Cycle means a anonymous func
type Wheel struct {
	Cycle
	spokes int
}

// how to define a struct-Wheel
func main() {
	var wheel Wheel
	wheel.x = 5
	wheel.y = 8
	wheel.radius = 22
	wheel.spokes = 33
	fmt.Println(wheel)	// {{{5 8} 22} 33}
	// you can't define like this
	// wrong_wheel := Wheel{5, 8, 22, 33}
	// wrong_wheel := Wheel{x: 5, y: 8, radius: 22, spokes: 33}
	
	//another way to define
	var w Wheel
	w = Wheel{
		Cycle: Cycle{
			Point: Point{x: 5, y: 8},
			radius: 22,
		},
		spokes: 44,
	}
	fmt.Println(w)
	// %#v means that print value liked golang
	fmt.Printf("%#v\n", w)	//main.Wheel{Cycle:main.Cycle{Point:main.Point{x:5, y:8}, radius:22}, spokes:44}	
	w.x = 4
	fmt.Println(w)
	fmt.Printf("%#v\n", w)
}