package main

import (
	"fmt" // 导入 fmt 包，用于格式化输出
)

type Circle interface {
	Area() float64      // 计算面积的方法
	Perimeter() float64 // 计算周长的方法
}
type CircleImpl struct {
	Radius float64 // 圆的半径
}

// CircleImpl 实现 Circle 接口的 Area 方法
func (c *CircleImpl) Area() float64 {
	return 3.14 * c.Radius * c.Radius // 返回圆的面积
}

// CircleImpl 实现 Circle 接口的 Perimeter 方法
func (c *CircleImpl) Perimeter() float64 {
	return 2 * 3.14 * c.Radius // 返回圆的周长
}

type power interface {
	Power() float64 // 计算功率的方法
}
type PowerImpl struct {
	Voltage float64 // 电压
	Current float64 // 电流
}

// PowerImpl 实现 power 接口的 Power 方法
func (p *PowerImpl) Power() float64 {
	return p.Voltage * p.Current // 返回功率
}

func main() {
	c := &CircleImpl{Radius: 5}
	//输出圆的半径，周长，面积
	fmt.Printf("圆的半径: %.2f, 面积: %.2f, 周长: %.2f\n", c.Radius, c.Area(), c.Perimeter()) // 输出圆面积和周长
	a := &PowerImpl{Voltage: 220, Current: 5}
	//输出电压，电流，功率
	fmt.Printf("电压: %.2f, 电流: %.2f, 功率: %.2f\n", a.Voltage, a.Current, a.Power())
}
