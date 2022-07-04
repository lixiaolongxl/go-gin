package main

import "fmt"

func main() {
	//a := 1.3
	//fmt.Println(&a)
	//fmt.Printf("%T", &a)
	var a *int
	fmt.Println(a) //空指针

	b := 123
	a = &b
	//* 可以去除指针里面的值
	*a = 234
	fmt.Println(*a, b)
}
