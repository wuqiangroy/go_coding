package main

import "fmt"

func quick_sort(array []int) {
	if len(array) < 2 {
		return
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
	middle := []int
	middle.append(pivot)
	new_array := quick_sort(low) + middle + quick_sort(big)
	return new_array
}