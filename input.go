package main

import "fmt"

func main() {
	var name, age string
	fmt.Println("请输入姓名年龄")
	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
}
