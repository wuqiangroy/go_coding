package main

import (
	"fmt"
	"os"
	"log"
)

// global varibal
var pwd string

func init() {
	// define a err to void using err := short declare
	var err error
	// corret way to give value to pwd
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd error: %v\n", err)
	}
}

func main() {
	// using global varible pwd
	fmt.Println(pwd)
}