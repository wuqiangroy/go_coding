package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	 "github.com/goinaction/code/tree/master/chapter2/sample/search"
)

// init called before main
func init() {
	// os standard output
	log.SetOutput(os.Stdout)
}

// the program's entrance
func main() {
	search.Run("president")
}