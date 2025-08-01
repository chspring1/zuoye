package main // 声明 main 包

import (
	"bufio" // 导入 bufio 包，用于缓冲读写
	"fmt"   // 导入 fmt 包，用于格式化输出
	"os"    // 导入 os 包，用于文件操作
)

func main() {
	filePath := "src/com/demo1/zuoye/fliedome/filedome02/test.txt"    // 定义文件路径
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) // 打开文件，写入模式，不存在则创建，权限为0666

	if err != nil { // 判断是否打开文件失败
		fmt.Println("Error opening file:", err) // 打印错误信息
		return                                  // 结束程序
	}
	defer file.Close() // 程序结束前关闭文件

	str := "Hello, World!\n"        // 要写入文件的字符串
	writer := bufio.NewWriter(file) // 创建带缓冲的写入对象
	for i := 0; i < 10; i++ {       // 循环10次
		writer.WriteString(str) // 每次写入一行字符串
	}
	writer.Flush() // 将缓冲区内容写入
}
