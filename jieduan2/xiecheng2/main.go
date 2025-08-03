package main

import (
	"fmt"  // 导入 fmt 包，用于格式化输出
	"sync" // 导入 sync 包，用于同步操作
	// 导入 time 包，用于时间相关操作
)

var (
	intChan chan int       // 整型通道
	wg      sync.WaitGroup // 定义一个 WaitGroup，用于等待协程完成
)

func writeData() {

	defer wg.Done() // 协程结束时通知等待组
	for i := 0; i < 50; i++ {
		intChan <- i            // 向通道写入数据
		fmt.Println("写入数据:", i) // 打印写入的数据
	}
	close(intChan) // 关闭通道，表示不再写入数据
}

func readData() int {

	defer wg.Done() // 协程结束时通知等待组
	// 使用 for 循环从通道中读取数据
	// 这里使用 range 来自动处理通道关闭
	val := 0 // 初始化 val 变量
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("接收到数据:", v) // 打印接收到的数据
		val = v
	}
	return val // 返回读取到的数据
}
func main() {
	intChan = make(chan int, 2)   // 创建一个缓冲区大小为10的整型通道
	wg.Add(2)                     // 添加两个协程到等待组
	go writeData()                // 启动写数据的协程               // 读取数据
	val := readData()             // 调用 readData 函数获取返回	值
	fmt.Println("最终接收到的数据:", val) // 打印最终接收到的数据
	wg.Wait()                     // 等待所有协程完成

}
