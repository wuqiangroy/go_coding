package main

import (
	"fmt"
	"log"
	"os"

	"/Users/qiang/Learn/go/go_coding/gopl/ch4/github"
)

func main() {
	resutl, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", resutl.TotalCount)
	for _, item := range resutl.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
