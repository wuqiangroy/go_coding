package main

import (
	"fmt"
	//"os"
)


type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	//os.Stdout realise Writer
	//w = os.Stdout

	fmt.Print(w, "Hello, World\n")
}
