package main

import (
	"fmt"
	"time"
)

func printOdd() {

	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
		time.Sleep(time.Second)
	}
}

func printEven() {

	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go printEven()
	printOdd()
}
