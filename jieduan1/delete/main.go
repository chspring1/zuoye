package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4}
	length := removeDuplicates(nums)
	fmt.Println("新长度:", length)
	fmt.Println("去重后的数组:", nums[:length])
}
