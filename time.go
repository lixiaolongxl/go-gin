package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)

	tn := time.Now()
	fmt.Println(tn)
	//	通过纳秒试驾车
	tn1 := tn.UnixNano()
	fmt.Println(tn1)
	tn2 := time.Unix(0, tn1)
	fmt.Println(tn2)

	tn3 := time.Date(2022, 2, 4, 8, 12, 30, 0, time.Local)
	fmt.Println(tn3)
	//	格式转化
	s := "2024-02-04 15:04:05"
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(tn.Format("2006-01-02 15:04:05"))

	//	字符串转
	fmt.Println(time.Parse("2006-01-02 15:04:05", s))
}
