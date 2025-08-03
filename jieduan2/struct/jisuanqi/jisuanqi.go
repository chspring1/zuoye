package main

import (
	"fmt" // 导入 fmt 包，用于格式化输出
	"strings"
)

type calculator interface {
	add() float64      // 加法
	subtract() float64 // 减法
	multiply() float64 // 乘法
	divide() float64   // 除法

}
type CalculatorImpl struct {
	Operand1 float64 // 操作数1
	Operand2 float64 // 操作数2
}

func (c *CalculatorImpl) add() float64 {
	return c.Operand1 + c.Operand2 // 返回加法结果
}
func (c *CalculatorImpl) subtract() float64 {
	return c.Operand1 - c.Operand2 // 返回减法结果
}
func (c *CalculatorImpl) multiply() float64 {
	return c.Operand1 * c.Operand2 // 返回乘法结果
}
func (c *CalculatorImpl) divide() float64 {
	if c.Operand2 == 0 {
		fmt.Println("除数不能为零") // 如果除数为零，输出错误信息
		return 0
	}
	return c.Operand1 / c.Operand2 // 返回除法结果
}
func main() {
	fmt.Println("欢迎使用计算器")
	for {
		fmt.Println("请输入一个表达式，例如 3+5 或 10-2 等，或输入 'exit' 退出计算器")
		var op string
		fmt.Scanln(&op)

		if strings.ToLower(op) == "exit" {
			fmt.Println("感谢使用计算器，再见！")
			break
		}

		idx := strings.IndexAny(op, "+-*/")
		if idx == -1 || idx == 0 || idx == len(op)-1 {
			fmt.Println("无效的输入，请输入包含运算符 +, -, *, / 的表达式")
			continue
		}

		operand1Str := strings.TrimSpace(op[:idx])
		operand2Str := strings.TrimSpace(op[idx+1:])
		operator := op[idx]

		var operand1, operand2 float64
		_, err1 := fmt.Sscanf(operand1Str, "%f", &operand1)
		_, err2 := fmt.Sscanf(operand2Str, "%f", &operand2)
		if err1 != nil || err2 != nil {
			fmt.Println("输入的数字格式有误，请重新输入")
			continue
		}

		calc := &CalculatorImpl{Operand1: operand1, Operand2: operand2}
		switch operator {
		case '+':
			fmt.Printf("结果: %.2f\n", calc.add())
		case '-':
			fmt.Printf("结果: %.2f\n", calc.subtract())
		case '*':
			fmt.Printf("结果: %.2f\n", calc.multiply())
		case '/':
			if operand2 == 0 {
				fmt.Println("除数不能为零")
			} else {
				fmt.Printf("结果: %.2f\n", calc.divide())
			}
		default:
			fmt.Println("无效的运算符，请输入 +, -, *, /")
		}
	}
	fmt.Println("计算器已退出")
	fmt.Println("感谢使用计算器，再见！")
}
