package main

import (
	"fmt"
)
//map是key-value
//map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.122, -299999.111,
	}
	fmt.Print(m["Bell Labs"], "\n", m, "\n")

	var m1 = map[string]Vertex{
		"Chengdu": Vertex{
			104.06, 30.67,
		},
		"Beijing": Vertex{
			116.46, 39.92,
		},
	}
	fmt.Print(m1, "\n", m1["Chengdu"].Lat)

	var m2 = map[string]Vertex{
		"Shanghai": {121.48, 31.22,},
		"Shenzhen": {114.07, 22.62,},
	}
	fmt.Print("\n", m2, "\n", m2["Shenzhen"])
}
