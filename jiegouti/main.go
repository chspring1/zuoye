package main // 声明 main 包

import (
	"fmt"  // 导入 fmt 包，用于格式化输出
	"math" // 导入 math 包，用于数学计算
)

// 定义 Shape 接口，包含 Area 和 Perimeter 方法
type Shape interface {
	Area() float64      // 计算面积的方法
	Perimeter() float64 // 计算周长的方法
}

// 定义 Rectangle 结构体
type Rectangle struct {
	Width  float64 // 矩形宽度
	Height float64 // 矩形高度
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // 返回矩形面积
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height) // 返回矩形周长
}

// 定义 Circle 结构体
type Circle struct {
	Radius float64 // 圆的半径
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius // 返回圆的面积
}

// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius // 返回圆的周长
}

func main() {
	r := Rectangle{Width: 3, Height: 4} // 创建一个矩形实例
	c := Circle{Radius: 5}              // 创建一个圆实例

	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", r.Area(), r.Perimeter()) // 输出矩形面积和周长
	fmt.Printf("圆形面积: %.2f, 周长: %.2f\n", c.Area(), c.Perimeter()) // 输出圆面积和周长

	// 使用 Shape 接口
	var s Shape // 定义 Shape 接口变量

	s = r                                                         // 将矩形赋值给接口变量
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter()) // 通过接口调用矩形方法

	s = c                                                         // 将圆赋值给接口变量
	fmt.Printf("圆形面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter()) // 通过接口调用圆方法
}
