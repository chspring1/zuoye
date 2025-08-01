package main // 声明 main 包

import (
	"fmt"  // 导入 fmt 包，用于格式化输出
	"sync" // 导入 sync 包，用于协程同步
	"time" // 导入 time 包，用于时间处理
)

func printOdd(wg *sync.WaitGroup) { // 打印奇数的协程函数，参数为 WaitGroup 指针
	defer wg.Done()               // 协程结束时通知 WaitGroup
	for i := 1; i <= 10; i += 2 { // 遍历 1 到 10 的奇数
		fmt.Println("奇数:", i) // 输出奇数
		//循环一次等待1s

		time.Sleep(1 * time.Second)
	}
}

func printEven(wg *sync.WaitGroup) { // 打印偶数的协程函数，参数为 WaitGroup 指针
	defer wg.Done()               // 协程结束时通知 WaitGroup
	for i := 2; i <= 10; i += 2 { // 遍历 2 到 10 的偶数
		fmt.Println("偶数:", i) // 输出偶数
		time.Sleep(1 * time.Second)
	}
}

func main() { // 主函数入口
	var wg sync.WaitGroup // 定义 WaitGroup，用于等待协程结束
	wg.Add(2)             // 设置 WaitGroup 计数为 2，表示有两个协程
	go printOdd(&wg)      // 启动打印奇数的协程
	go printEven(&wg)     // 启动打印偶数的协程
	wg.Wait()             // 等待所有协程结束
}
