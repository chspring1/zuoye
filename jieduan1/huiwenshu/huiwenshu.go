package main // 声明 main 包

import (
	"fmt" // 导入 fmt 包，用于格式化输出
)

func main() {
	number := []int{1, 2, 3, 2, 1} // 定义一个整数切片，待判断是否为回文数

	long := len(number) // 获取切片长度

	for i := 0; i < long; i++ { // 遍历切片
		fmt.Println(number[i], " ", number[long-i-1]) // 输出当前元素和对称位置的元素
		if number[i] != number[long-i-1] {            // 如果当前元素和对称位置元素不相等
			fmt.Println("不是回文数") // 输出不是回文数
			return               // 结束程序
		}
	}
	fmt.Println("是回文数") // 如果全部相等，输出是回文数
}
