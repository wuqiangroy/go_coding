package main

import "fmt"

func bubble_sort(array []int) {
	if len(array) < 2 {
		return 
	}
	for i:=0;i<len(array);i++ {
		for j:=0;j<len(array)-1;j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func main() {
	array := []int {1, 6, 7, 3, 2, 0, 11, 3, 4, 8, 13, 12}
	bubble_sort(array)
	fmt.Print(array)
}