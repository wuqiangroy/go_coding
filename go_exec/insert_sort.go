package main

import "fmt"

func insert_sort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j -= 1
		}
		array[j+1] = key
	}
	fmt.Print(array)
}

func main() {
	array := []int{1, 2, 4, 6, 8, 4, 0, 1, 3, 10, 11}
	insert_sort(array)
	// fmt.Print(array)
}