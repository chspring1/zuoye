package main

import (
	"encoding/json"
	"fmt"
)

// 定义结构体，使用json标签
type Person struct {
	FirstName  string  `json:"first_name"`       // 映射到JSON中的first_name
	LastName   string  `json:"last_name"`        // 映射到JSON中的last_name
	Age        int     `json:"age"`              // 映射到JSON中的age
	IsEmployed bool    `json:"is_employed"`      // 映射到JSON中的is_employed
	Salary     float64 `json:"salary,omitempty"` // omitempty表示如果为零值则省略
	SecretKey  string  `json:"-"`                // "-"表示该字段不参与JSON序列化
}

func main() {
	// 结构体实例化
	person := Person{
		FirstName:  "张",
		LastName:   "三",
		Age:        30,
		IsEmployed: true,
		Salary:     0, // 零值
		SecretKey:  "abc123",
	}

	// 序列化为JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}
	fmt.Println("序列化结果:", string(jsonData))

	// 反序列化JSON
	var newPerson Person
	jsonStr := `{"first_name":"李","last_name":"四","age":25,"is_employed":false}`
	err = json.Unmarshal([]byte(jsonStr), &newPerson)
	if err != nil {
		fmt.Println("JSON解码错误:", err)
		return
	}
	fmt.Printf("反序列化结果: %+v\n", newPerson)
}
