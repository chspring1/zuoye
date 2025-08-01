package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个整数类型的通道
	ch := make(chan int)

	// 创建 WaitGroup 等待协程完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动生产者协程
	go func() {
		defer wg.Done()
		defer close(ch) // 发送完毕后关闭通道

		// 生成并发送1到10的整数
		for i := 1; i <= 10; i++ {
			fmt.Printf("发送: %d\n", i)
			ch <- i
		}
	}()

	// 启动消费者协程
	go func() {
		defer wg.Done()

		// 从通道接收并打印整数
		for num := range ch {
			fmt.Printf("接收: %d\n", num)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
