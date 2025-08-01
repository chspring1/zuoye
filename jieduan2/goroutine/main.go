package main

import (
	"fmt"
	"sync"
)

// 计算1-200的各个数的阶乘，并且把各个数放到map中
// 打印出，需要用goroutine实现

var (
	myMap = make(map[int]int, 200) // 修改容量为200
	mutex sync.Mutex               // 互斥锁保护map的并发访问
	wg    sync.WaitGroup           // 等待所有协程完成
)

func test(n int) {
	defer wg.Done() // 协程完成时减少计数器

	res := 1
	for i := 1; i <= n; i++ {
		res *= 2
	}

	mutex.Lock() // 加锁保护map写入
	myMap[n] = res
	mutex.Unlock() // 解锁
}

func main() {
	// 启动200个协程计算阶乘
	for i := 1; i <= 200; i++ {
		wg.Add(1) // 增加等待计数
		go test(i)
	}

	wg.Wait() // 等待所有协程完成

	// 按顺序打印结果
	for i := 1; i <= 200; i++ {
		fmt.Printf("数字%d的阶乘是%d\n", i, myMap[i])
	}
}
