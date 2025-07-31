package main // 声明 main 包

import (
	"bufio" // 导入 bufio 包，用于缓冲读写
	"fmt"   // 导入 fmt 包，用于格式化输出
	"os"    // 导入 os 包，用于文件操作
)

func main() {
	file, err := os.Open("D:/file/test.txt") // 打开指定路径的文件
	if err != nil {                          // 如果打开文件失败
		fmt.Println("打开文件失败:", err) // 输出错误信息
		return                      // 结束程序
	}
	fmt.Println("打开文件成功:", file) // 打印文件打开成功信息
	defer file.Close()           // 程序结束前关闭文件

	reader := bufio.NewReader(file) // 创建一个带缓冲的 Reader 读取文件内容

	for { // 循环读取文件内容
		str, err := reader.ReadString('\n') // 读取一行内容
		if err != nil {                     // 如果读取出错
			if err.Error() == "EOF" { // 如果是文件结尾
				fmt.Println("文件读取完毕") // 输出文件读取完毕
				break                 // 跳出循环
			}
			// 其他错误不处理，继续读取
		}
		fmt.Print(str) // 输出读取到的内容
	}
}
