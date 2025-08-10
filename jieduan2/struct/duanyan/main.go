package main

import "fmt"

type Person interface {
	SayHello()
	Eat()
	Work()
}

type Student struct {
	Name  string
	Age   int
	Grade string
}

type Teacher struct {
	Name    string
	Age     int
	Subject string
}

func (s Student) SayHello() {
	fmt.Printf("大家好，我是%s，今年%d岁，读%s年级！\n", s.Name, s.Age, s.Grade)
}
func (t Teacher) SayHello() {
	fmt.Printf("大家好，我是%s，今年%d岁，教授%s！\n", t.Name, t.Age, t.Subject)
}
func (s Student) Eat() {
	fmt.Printf("%s正在吃饭。\n", s.Name)
}
func (t Teacher) Eat() {
	fmt.Printf("%s正在吃饭。\n", t.Name)
}
func (s Student) Work() {
	fmt.Printf("%s正在学习。\n", s.Name)
}
func (t Teacher) Work() {
	fmt.Printf("%s正在教书。\n", t.Name)
}

// 断言应用：根据接口变量的真实类型做不同处理
func handlePerson(p Person) {
	// 类型switch方式
	switch v := p.(type) {
	case Student:
		fmt.Printf("[断言] 这是学生，年级：%s，名字：%s\n", v.Grade, v.Name)
		v.SayHello()
		v.Eat()
		v.Work()
	case Teacher:
		fmt.Printf("[断言] 这是老师，科目：%s，名字：%s\n", v.Subject, v.Name)
		v.SayHello()
		v.Eat()
		v.Work()
	default:
		fmt.Println("[断言] 未知类型")
	}

	// 单一类型断言方式
	if s, ok := p.(Student); ok {
		fmt.Printf("[单一断言] %s 是学生，年级：%s\n", s.Name, s.Grade)
	}
	if t, ok := p.(Teacher); ok {
		fmt.Printf("[单一断言] %s 是老师，科目：%s\n", t.Name, t.Subject)
	}
}

func main() {
	var p Person

	p = Student{Name: "小明", Age: 18, Grade: "高三"}
	handlePerson(p)

	fmt.Println()

	p = Teacher{Name: "李老师", Age: 35, Subject: "数学"}
	handlePerson(p)
}
