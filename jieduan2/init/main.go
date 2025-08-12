package main

import "fmt"

// 全局变量
var config map[string]string
var dbConnected bool

func init() {
	// 初始化全局变量
	config = make(map[string]string)
	config["app_name"] = "MyApp"
	config["version"] = "1.0"

	// 模拟资源初始化，如数据库连接
	dbConnected = true // 假设连接成功

	// 检查环境变量（示例）
	// os.Getenv("ENV_VAR")

	fmt.Println("init 函数被执行：用于初始化操作")
	fmt.Println("配置已加载：", config)
	fmt.Println("数据库连接状态：", dbConnected)
}

func main() {
	fmt.Println("main 函数被执行：程序入口")
}
