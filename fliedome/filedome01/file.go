package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//只适合文件比较小的情况
	file := "D:/file/test.txt"
	contenr, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读取文件失败:", err)

	}
	fmt.Println("文件内容:", contenr)
	fmt.Println("文件内容:", string(contenr))

}
