package main

import (
	"fmt"
	"reflect"
)

type graphic interface {
	Area() float64
	Perimeter() float64
}

type square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

func (s square) Area() float64 {
	return s.Side * s.Side
}

func (s square) Perimeter() float64 {
	return 4 * s.Side
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 类型断言
func getInfo(g graphic) {
	switch shape := g.(type) {
	case square:
		fmt.Printf("正方形面积: %f, 周长: %f\n", shape.Area(), shape.Perimeter())
	case Circle:
		fmt.Printf("圆形面积: %f, 周长: %f\n", shape.Area(), shape.Perimeter())
	default:
		fmt.Println("Unknown shape")
	}
}

// 反射方法
func reflectInfo(g graphic) {
	val := reflect.ValueOf(g)
	typ := reflect.TypeOf(g)

	fmt.Printf("Type: %s\n", typ.Name())
	fmt.Printf("Value: %v\n", val)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("Field %d - Type: %s, Value: %v\n", i, field.Type(), field.Interface())
	}
	if typ.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fmt.Printf("Field %d - Type: %s, Value: %v\n", i, field.Type(), field.Interface())
		}
	}
	if typ.Kind() == reflect.Ptr {
		fmt.Printf("Pointer to: %s\n", typ.Elem().Name())
	}
}

func main() {
	s := square{Side: 5}
	c := Circle{Radius: 7}

	getInfo(s)
	getInfo(c)
	reflectInfo(s)
	reflectInfo(c)
}
