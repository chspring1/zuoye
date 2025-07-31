package main

func doubleSlice(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 10, 100, 200, 500}
	doubledSlice := doubleSlice(slice)
	for _, value := range doubledSlice {
		println(value)
	}
}
