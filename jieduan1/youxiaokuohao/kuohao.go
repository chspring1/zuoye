package main // 声明 main 包

func isValid(s string) bool { // 定义判断括号有效性的函数
	n := len(s)   // 获取字符串长度
	if n%2 == 1 { // 如果长度为奇数，肯定无效
		return false
	}
	pairs := map[byte]byte{ // 定义右括号与左括号的映射关系
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}        // 用于存放左括号的栈
	for i := 0; i < n; i++ { // 遍历字符串每个字符
		if pairs[s[i]] > 0 { // 如果是右括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { // 栈为空或栈顶不是对应的左括号
				return false // 括号不匹配，返回 false
			}
			stack = stack[:len(stack)-1] // 匹配成功，弹出栈顶元素
		} else { // 如果是左括号
			stack = append(stack, s[i]) // 入栈
		}
	}
	return len(stack) == 0 // 栈为空则括号全部匹配，返回 true，否则返回 false
}

func main() { // 主函数
	s := "{{[()]}}"     // 测试字符串1
	println(isValid(s)) // 输出判断结果
	s = "{[(])}"        // 测试字符串2
	println(isValid(s)) // 输出判断结果
}
