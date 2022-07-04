package main

import "fmt"

func main() {
	//const a = 1
	//const b = 1.5
	//const PI = 3.14
	//const name = "李小龙"
	//fmt.Println(a + b)

	const (
		a = 5
		b
		c
	)
	fmt.Println(a, b, c)
	const (
		d = iota
		e
		f = 23
		g
		x
		h = iota
		i
	)
	fmt.Println(d, e, f, g, x, h, i)
}
