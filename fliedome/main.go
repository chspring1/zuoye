package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/file/test.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	fmt.Println("打开文件成功:", file)
	defer file.Close()
	//创建一个*Reader类型的变量
	reader := bufio.NewReader(file)
	//读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("文件读取完毕")
				break
			}

		}
		//输出内容
		fmt.Print(str)
	}
}
