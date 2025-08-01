package main // 声明 main 包

import "fmt" // 导入 fmt 包，用于格式化输出

func twosum(nums []int, target int) []int { // 定义函数，查找和为 target 的两个元素的下标
	for i := 0; i < len(nums); i++ { // 外层循环，遍历每个元素
		for j := 0; j < len(nums); j++ { // 内层循环，遍历每个元素
			if i != j && nums[i]+nums[j] == target { // 如果不是同一个元素且两数之和等于目标值
				return []int{i, j} // 返回两个元素的下标
			}
		}
	}
	return nil // 如果没有找到，返回 nil
}

func main() { // 主函数入口
	nums := []int{2, 7, 11, 15}    // 定义整数切片
	target := 9                    // 定义目标值
	result := twosum(nums, target) // 调用 twosum 函数，查找结果
	if result != nil {             // 如果找到了结果
		fmt.Println(result[0], result[1]) // 输出两个下标
	} else { // 如果没有找到结果
		fmt.Println("没有找到结果") // 输出提示信息
	}
}
