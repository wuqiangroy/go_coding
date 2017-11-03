package main

import (
    "fmt"
)

type employee struct {
    ID int
    name string
    age int
    position string
    salary int
}

func main() {
    var roy employee
    roy.ID = 20120436
    fmt.Println(roy)
    roy.position = "Chengdu Pidu District"
    position := &roy.position
    fmt.Println(*position)
}
