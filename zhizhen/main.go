package main

import "fmt"

func addTen(x *int) int {
	*x += 10
	return *x
}

func main() {
	y := 5
	addTen(&y)
	fmt.Println(y) // 输出: 15
}
