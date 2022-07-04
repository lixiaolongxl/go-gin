package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64 = 0.1
	fmt.Println(f)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	var a float64 = 1.3
	var b float32 = 3.4
	var c float64
	c = a + float64(b)
	fmt.Println(c)
}
