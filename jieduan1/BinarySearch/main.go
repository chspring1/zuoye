package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("low: %d\n", low)
		fmt.Printf("mid: %d\n", mid)
		fmt.Printf("high: %d\n", high)
		fmt.Println("--------------------------------")
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 3
	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}
