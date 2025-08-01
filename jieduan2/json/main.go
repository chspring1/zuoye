package main

import (
	"encoding/json"
	"fmt"
)

// 导入 fmt 包，用于格式化输出

type Person struct { // 定义 Person 结构体
	Name     string
	Age      int     // 定义 Name 和 Age 字段
	Birthday string  // 定义 Birthday 字段
	Salary   float64 // 定义 Salary 字段
	skill    string  // 定义 skills 字段
}

func testPerson() { // 定义函数 testPerson
	p := Person{Name: "张三", Age: 18, Birthday: "2005-01-01", Salary: 10000.50, skill: "编程"}
	data, err := json.Marshal(p) // 将 Person 结构体转换为 JSON 格式
	if err != nil {              // 如果转换出错
		panic(err) // 抛出异常
	}
	fmt.Println(string(data)) // 输出 JSON 字符串
}

func testMap() { // 定义函数 testMap
	var a map[string]interface{} // 定义一个 map，键为字符串，值为任意类型
	a = make(map[string]interface{})
	a["Name"] = "李四"             // 设置键为 "Name" 的值为 "张三"
	a["Age"] = 19                // 设置键为 "Age" 的值为 18
	a["Birthday"] = "2005-01-01" // 设置键为 "Birthday" 的值为 "2005-01-01"
	a["Salary"] = 10000.50
	a["skill"] = "编程"            // 设置键为 "skill" 的值为 "编程"
	data, err := json.Marshal(a) // 将 map 转换为 JSON 格式
	if err != nil {              // 如果转换出错
		panic(err) // 抛出异常
	}
	fmt.Println(string(data)) // 输出 JSON 字符串
}

func testSlice() {
	var slice []map[string]interface{} // 定义一个切片，元素为 map
	var m1 map[string]interface{}      // 定义一个 map
	m1 = make(map[string]interface{})
	m1["Name"] = "王五"
	m1["Age"] = 20
	m1["Birthday"] = "2005-01-01"
	m1["Salary"] = 12000.50
	m1["skill"] = "编程"
	slice = append(slice, m1)     // 将 map 添加到切片中
	var m2 map[string]interface{} // 定义另一个 map
	m2 = make(map[string]interface{})
	m2["Name"] = "赵六"
	m2["Age"] = 21
	m2["Birthday"] = "2004-01-01"
	m2["Salary"] = 15000.50
	m2["skill"] = []string{"编程", "绘画"}
	slice = append(slice, m2)        // 将另一个 map 添加到切片中
	data, err := json.Marshal(slice) // 将切片转换为 JSON 格式
	if err != nil {                  // 如果转换出错
		panic(err) // 抛出异常
	}
	fmt.Println(string(data)) // 输出 JSON 字符串
	// 输出: [{"Name":"王五","Age":20,"Birthday":"2005-01-01","Salary":12000.5,"skill":"编程"},{"Name":"赵六","Age":21,"Birthday":"2004-01-01","Salary":15000.5,"skill":"编程"}]
}

// 定义函数 testSlice
// 对基本类型序列化
func testFloat64() {
	var num float64 = 123.456      // 定义一个 float64 类型的变量
	data, err := json.Marshal(num) // 将 float64 转换为 JSON 格式
	if err != nil {                // 如果转换出错
		panic(err) // 抛出异常
	}
	fmt.Println(string(data)) // 输出 JSON 字符串
}

func main() { // 主函数入口
	testPerson()  // 调用 testPerson 函数
	testMap()     // 调用 testMap 函数
	testSlice()   // 调用 testSlice 函数
	testFloat64() // 调用 testFloat64 函数
} // 输出 JSON 格式的 Person 结构体
// 输出: {"Name":"张三","Age":18,"Birthday":"2005-01-01","Salary":10000.5}
