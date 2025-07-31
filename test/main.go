package main

import (
	"fmt"
	"runtime"
)

func main() {

	a := runtime.NumCPU()
	fmt.Print("当前系统的CPU数量:", a, "\n")
}
