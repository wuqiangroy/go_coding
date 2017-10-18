package main

import "fmt"

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i + 1, j-1 {
		a[i], a[j] = a[j] , a[i]
	}
}

func reverse2(a []int) {
	j := 0
	for i := len(a)-1; i >= int(len(a))/2; i-- {
		a[i], a[j] = a[j], a[i]
		j++
	}
}

func main() {
	array := []int{1, 5, 6, 7, 8, 9}
	reverse2(array)
	fmt.Println(array)
	// array = append(array, 5)
	// fmt.Println(array)
}