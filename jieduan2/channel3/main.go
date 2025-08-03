/*编写一个程序，使用 sync.Mutex
来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。*/

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println("最终计数器的值:", counter)
}
