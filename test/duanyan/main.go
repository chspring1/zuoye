package main

//类型断言实列
import (
	"fmt"
)

type MyInterface interface{}

type MyStruct struct {
	Name string
}

func main() {
	var myVar MyInterface = MyStruct{Name: "example"}
	// 类型断言

	a := myVar.(MyStruct) // 将 myVar 断言为 MyStruct 类型

	// 使用类型断言后的变量
	fmt.Println(a.Name) // 输出: example
}
