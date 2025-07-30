package main

import (
	"fmt"
)

func main() {

	number := []int{1, 2, 3, 2, 1}

	long := len(number)

	for i := 0; i < long; i++ {
		fmt.Println(number[i], " ", number[long-i-1])
		if number[i] != number[long-i-1] {

			fmt.Println("不是回文数")
			return
		}

	}
	fmt.Println("是回文数")
}
