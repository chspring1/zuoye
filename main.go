package main

import "fmt"

func main() {
	result := 0

	nums := []int{2, 2, 1, 1, 3, 3, 4, 4, 5}
	for _, num := range nums {

		result = num
		nubmber := 0
		for _, num := range nums {
			if result == num {
				nubmber = nubmber + 1
			}
		}
		if nubmber == 1 {
			fmt.Println("唯一的数字是%:", result)
			break
		}

	}
}
