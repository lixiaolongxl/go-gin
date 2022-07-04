package main

import "fmt"

func main() {
	a := new(int)
	fmt.Println(a)
	*a = 99
	fmt.Println(*a) //可直接给指针赋值
	//  if
	if *a > 100 {
		fmt.Println("大于100") //可直接给指针赋值
	} else {
		fmt.Println("小于100") //可直接给指针赋值
	}

	//	switch
	b := 9
	switch {
	case b >= 100:
		fmt.Println("10")
	case b >= 50:
		fmt.Println("8")
	case b >= 10:
		fmt.Println("16")
	default:
		fmt.Println("default")
	}
	mouth := 10
	switch mouth {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31")
	case 2:
		fmt.Println("28/29")
	default:
		fmt.Println("31")

	}
}
