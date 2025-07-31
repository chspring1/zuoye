package main

import "fmt"

func twosum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twosum(nums, target)
	if result != nil {
		fmt.Println(result[0], result[1]) // 输出: 0 1
	} else {
		fmt.Println("没有找到结果")
	}
}
