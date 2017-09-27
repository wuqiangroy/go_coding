package main

import "fmt"

func quick_sort(array []int) {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	var low []int
	var big []int
	for i:=0;i<len(array);i++ {
		if array[i] < pivot {
			low.append(array[i])
		} else if array[i] > pivot {
			big.append(array[i])
		}
	}
	middle := []int {pivot}
	return quick_sort(low) + middle + quick_sort(big)
}

func main() {
	array := []int {1, 6, 7, 3, 2, 0, 11, 3, 4, 8, 13, 12}
	fmt.Print(quick_sort(array))
	fmt.Print("I do not Know how to code quick sort in recursively")
}