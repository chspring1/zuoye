package main

import "fmt"

// 定义一个接口
type Animal interface {
	Speak() string
}

// 实现接口的结构体
type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪！"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵！"
}

type Bird struct{}

func (b Bird) Speak() string {
	return "啾啾！"
}

type Cow struct{}

func (cw Cow) Speak() string {
	return "哞哞！"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println("Dog说:", a.Speak())

	a = Cat{}
	fmt.Println("Cat说:", a.Speak())

	a = Bird{}
	fmt.Println("Bird说:", a.Speak())

	a = Cow{}
	fmt.Println("Cow说:", a.Speak())

	fmt.Println("\n--- 接口切片多态遍历 ---")
	animals := []Animal{Dog{}, Cat{}, Bird{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
