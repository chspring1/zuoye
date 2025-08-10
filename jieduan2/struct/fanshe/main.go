package main

import (
	"fmt"
	"reflect"
)

// Person 接口
type Person interface {
	SayHello()
}

// Student 结构体实现 Person 接口
type Student struct {
	Name string
	Age  int
}

func (s Student) SayHello() {
	fmt.Printf("大家好，我是%s，今年%d岁！\n", s.Name, s.Age)
}

func reflectInfo(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Println("类型：", t)
	fmt.Println("值：", v)
	// 用反射判断类型并执行不同操作
	switch t.Kind() {
	case reflect.Struct:
		fmt.Println("这是一个结构体，遍历字段：")
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)
			fmt.Printf("字段名：%s，类型：%s，值：%v\n", field.Name, field.Type, value)
		}
		// 结构体类型名判断
		if t.Name() == "Student" {
			fmt.Println("[反射判断] 这是 Student 类型，可以执行学生相关操作")
		}
	case reflect.Ptr:
		// 如果是指针，递归处理其元素
		fmt.Println("这是一个指针，递归处理其元素：")
		reflectInfo(v.Elem().Interface())
	default:
		fmt.Println("[反射判断] 未知或不处理的类型")
	}
}

func main() {
	var p Person = Student{Name: "小明", Age: 18}
	p.SayHello()
	fmt.Println("--- 反射信息 ---")
	reflectInfo(p)
	fmt.Println("--- 反射结构体本体 ---")
	reflectInfo(Student{Name: "小明", Age: 18})
}
