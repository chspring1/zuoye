package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string)
	a["name"] = "侯文进"
	a["age"] = "18"
	a["address"] = "焦作"
	fmt.Println("我叫", a["name"], "今年", a["age"], "岁，住在", a["address"], "。")

}
