package main

import (
	"fmt"
)

type power interface {
	Power() float64      // 计算功率的方法
	Resistance() float64 // 计算电阻的方法
}
type PowerImpl struct {
	Voltage    float64 // 电压
	Current    float64 // 电流
	Resistance float64 // 电阻
}

// PowerImpl 实现 power 接口的 Power 方法
func (p *PowerImpl) Power() float64 {
	return p.Voltage * p.Current // 返回功率
}

// PowerImpl 实现 power 接口的 Resistance 方法
func (p *PowerImpl) GetResistance() float64 {

	return p.Voltage / p.Current // 返回电阻
}
func main() {

	fmt.Println("请输入电流")
	var current float64
	fmt.Scanln(&current) // 从标准输入读取电流
	fmt.Println("请输入电压")
	var voltage float64
	fmt.Scanln(&voltage) // 从标准输入读取电压
	a := &PowerImpl{Voltage: current, Current: voltage}

	// 输出电压，电流，功率
	fmt.Printf("电压: %.2f, 电流: %.2f, 功率: %.2f ,电阻:%2f \n", a.Voltage, a.Current, a.Power(), a.GetResistance())
}
