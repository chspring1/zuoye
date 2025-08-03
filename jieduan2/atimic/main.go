package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 全局唯一ID生成器
type IDGenerator struct {
	counter int64
}

func (g *IDGenerator) NextID() int64 {
	return atomic.AddInt64(&g.counter, 1)
}

func (g *IDGenerator) CurrentID() int64 {
	return atomic.LoadInt64(&g.counter)
}

// 用户ID生成器
var userIDGen = &IDGenerator{}

// 订单ID生成器
var orderIDGen = &IDGenerator{}

func createUser() int64 {
	userID := userIDGen.NextID()
	fmt.Printf("创建用户，ID: %d\n", userID)
	return userID
}

func createOrder() int64 {
	orderID := orderIDGen.NextID()
	fmt.Printf("创建订单，ID: %d\n", orderID)
	return orderID
}

func main() {
	var wg sync.WaitGroup

	// 并发创建用户和订单
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			createUser()
		}()
		go func() {
			defer wg.Done()
			createOrder()
		}()
	}

	wg.Wait()
	fmt.Printf("最终用户ID: %d, 订单ID: %d\n",
		userIDGen.CurrentID(), orderIDGen.CurrentID())
}
