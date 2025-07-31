package main

import (
	"fmt"
)

func main() {

	str := []string{"flower", "flow", "flight"}
	var d string
	for i := 1; i < 10; i++ {
		fmt.Println("-----------------------------------------------")
		a := str[1][:i]

		b := str[2][:i]

		c := str[0][:i]

		if a == b && b == c {
			d = str[1][:i]
		} else {
			break
		}

	}
	fmt.Println("最长公共前缀为：", d)
}
