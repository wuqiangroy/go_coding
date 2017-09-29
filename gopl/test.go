package main

import "fmt"

func main() {
	dic := make(map[string]string)
	dic["name"] = "wuqiangroy"
	dic["dream"] = "programmer"
	dic["salary"] = "50k per month"
	for key, value := range dic {
		fmt.Printf("the key is %s, and value is %s\n", key, value)
	}
}