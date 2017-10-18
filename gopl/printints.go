package main

import (
	"fmt"
	"bytes"
)

func IntsToString(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	array := []int{1, 2, 3}
	fmt.Printf("%v\n", array)
	fmt.Println(IntsToString(array))
}