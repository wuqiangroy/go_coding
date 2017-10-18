package main

import "fmt"

func main() {
	// dic := make(map[string]string)
	// dic["name"] = "wuqiangroy"
	// dic["dream"] = "programmer"
	// dic["salary"] = "50k per month"
	// for key, value := range dic {
	// 	fmt.Printf("the key is %s, and value is %s\n", key, value)
	// }
	// fmt.Println("\a")
	fmt.Println(comma("namespace"), "\n")
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s, b, s2)
}

func comma(s string) string {
	n := len(s)
	if n > 3 {
		return s[:n-3] + "," +s[n-3:]
	}
	return s
}