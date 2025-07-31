package main // 声明 main 包

func doubleSlice(slice []int) []int { // 定义函数，接收一个整数切片，返回每个元素都乘以2后的切片
	for i := 0; i < len(slice); i++ { // 遍历切片
		slice[i] *= 2 // 当前元素乘以2
	}
	return slice // 返回修改后的切片
}

func main() { // 主函数入口
	slice := []int{1, 2, 3, 4, 5, 10, 100, 200, 500} // 定义一个整数切片
	doubledSlice := doubleSlice(slice)               // 调用 doubleSlice 函数，得到新切片
	for _, value := range doubledSlice {             // 遍历新切片
		println(value) // 输出每个元
	}
}
