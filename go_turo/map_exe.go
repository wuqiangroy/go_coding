package main

import (
	//"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount1(s string) map[string]int {
	wordcount := make(map[string]int)
	for _, st := range s {
		_, ok := wordcount[string(st)]
		if ok {
			//wordcount[string(st)] += 1
			wordcount[string(st)] ++
			} else {
			wordcount[string(st)] = 1
		}
		}
	return wordcount
	}

func WordCount2(s string) map[string]int {
	//lis := strings.Split(s, " ")
	lis := strings.Fields(s)
	dic := make(map[string]int)
	for _, st := range lis {
		_, ok := dic[st]
		if ok {
			dic[st] ++
		} else {
			dic[st] = 1
		}
	}
	return dic
}
func main() {
	//s := "namespace"
	//fmt.Print(string(s[1]))
	fmt.Print(WordCount1("woshiyigedahuaidan"), "\n")
	fmt.Print(WordCount2("my name is Roy"), "\n")
	fmt.Print(WordCount2("  foo bar  baz   "), "\n")
	//wc.Test(WordCount)
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}