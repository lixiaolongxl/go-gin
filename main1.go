package main

import "fmt"

func main() {
	a := 12
	fmt.Println(a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%x %X\n", 12, 18)
	fmt.Printf("%T", 1)
	fmt.Printf("%p\n", &a)
	fmt.Printf("花费了%d%% \n 好 \t的", 20)
	fmt.Printf("%c", 65)
}
