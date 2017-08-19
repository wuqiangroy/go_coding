package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Print(Vertex{2, 3})
	v := Vertex{1, 2}
	v.X = 3
	fmt.Print(v, v.X)
	//结构体字段可以通过结构体指针来访问。
	//通过指针间接的访问是透明的。
	v1 := Vertex{2, 4}
	p := &v1
	p.X = 1e9
	fmt.Print("\n", v1)
	v2 := Vertex{X: 2} // 省略Y: 0
	v3 := Vertex{} // 省略X: 0， Y: 0
	p1 := &Vertex{4, 5} // 类型为&Vertex
	fmt.Print("\n", v2, v3, p1)
}
