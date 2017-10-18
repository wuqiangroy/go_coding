package main

import (
	"fmt"
	"strconv"
)

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

	// x := 123
	// y := fmt.Sprintf("%d", x)
	// fmt.Println(y, strconv.Itoa(x))
	// fmt.Println(strconv.FormatInt(int64(x), 2))
	// fmt.Println(strconv.FormatInt(int64(x), 8))
	// fmt.Println(strconv.FormatInt(int64(x), 10))
	// fmt.Println(strconv.FormatInt(int64(x), 16))
	x, err := strconv.Atoi("12n")
	y, err := strconv.ParseInt("12n", 10, 64)
	fmt.Println(x, y, err)
	months := []string{1:"January", 2: "Feberay"}
	fmt.Println(months)
}

func comma(s string) string {
	n := len(s)
	if n > 3 {
		return s[:n-3] + "," +s[n-3:]
	}
	return s
}