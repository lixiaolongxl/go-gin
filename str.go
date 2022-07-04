package main

import (
	"fmt"
	"strconv"
)

//数字字符串转化
func main() {
	s := "11"
	//i, _ := strconv.ParseInt(s, 10, 64)
	//fmt.Println(i + 6)

	i, _ := strconv.Atoi(s)
	fmt.Println(i + 6)

	var n int = 11
	//s1 := strconv.FormatInt(int64(n), 16)
	s1 := strconv.Itoa(n)
	fmt.Println(s1)
	fmt.Println('a')
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>")
	f := 1.5

	s3 := strconv.FormatFloat(f, 'f', 3, 64)
	fmt.Println(s3)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>")

}
