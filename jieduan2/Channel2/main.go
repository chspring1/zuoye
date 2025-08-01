package main // 声明 main 包

/*实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
 */ // 多行注释说明程序功能

import (
	"fmt"  // 导入 fmt 包，用于格式化输出
	"sync" // 导入 sync 包，用于协程同步
)

var (
	// 定义全局变量
	ch = make(chan int, 2) // 创建缓冲通道，可以存储10个整数而不阻塞
	wg sync.WaitGroup      // 定义一个 WaitGroup，用于等待协程完成
)

func shenchan() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i                      // 将整数 i 发送到通道 ch
		fmt.Println("生产者协程发送数据:", i) // 打印发送的数据
	}
}

func xiaofeizje() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		data := <-ch                    // 从通道 ch 接收数据
		fmt.Println("消费者协程接收数据:", data) // 打印接收到的数据
	}
}

func main() { // 主函数入口
	wg.Add(2)       // 添加两个协程到 WaitGroup
	go shenchan()   // 启动生产者协程
	go xiaofeizje() // 启动消费者协程
	wg.Wait()       // 等待所有协程完成
} // 主函数结束
