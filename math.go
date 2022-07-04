package main

import (
	"fmt"
	"math"
	"math/big"

	"crypto/rand"
)

func main() {
	//var a int8 = 8
	//var b int16 = 16
	//var c int32 = 32
	//var d int64 = 64
	//e := 45
	////类型转换 类型（值）
	//fmt.Println(a, b, c, d, e)
	//fmt.Println(int16(a) + b)

	d := 17
	o := 021
	x := 0x11
	fmt.Println(d, o, x)
	//	转化为二进制
	b := fmt.Sprintf("%b", d)
	fmt.Println(b)
	a, c := 12.4, 13.5
	fmt.Println(math.Max(a, c))
	fmt.Println(math.Min(a, c))
	//假随机
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(10))
	//fmt.Println(rand.Intn(10))
	//fmt.Println(rand.Intn(10))
	//fmt.Println(rand.Intn(10))
	//真随机
	for i := 0; i < 20; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Println(result)
	}

}
