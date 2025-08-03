package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64          // 原子操作计数器，必须使用 int64 类型
	wg      sync.WaitGroup // 等待组，用于等待所有协程完成
)

// 使用原子操作递增计数器
func atomicIncrement() {
	defer wg.Done() // 协程结束时通知等待组

	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1) // 原子操作递增计数器
	}
}

func main() {
	fmt.Println("使用原子操作实现无锁计数器")
	fmt.Println("启动10个协程，每个协程递增1000次")

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)            // 增加等待计数
		go atomicIncrement() // 启动协程
	}

	wg.Wait() // 等待所有协程完成

	// 使用原子操作读取最终值
	finalValue := atomic.LoadInt64(&counter)

	fmt.Printf("最终计数器的值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", 10*1000)

	if finalValue == 10000 {
		fmt.Println("结果正确！原子操作成功保护了共享资源")
	} else {
		fmt.Println("结果错误！存在并发问题")
	}
}
