package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine 模拟异步数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自 ch1 的数据"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自 ch2 的数据"
	}()

	for i := 0; i < 2; i++ {
		select {
		case v1 := <-ch1:
			fmt.Println("收到：", v1)
		case v2 := <-ch2:
			fmt.Println("收到：", v2)
		case <-time.After(3 * time.Second):
			fmt.Println("超时，无数据可读")
		}
	}
}
