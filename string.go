package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串截取
	s := "smaSlming张"
	fmt.Println(len(s)) // 获取字节长度
	//fmt.Println(s[:5])  //第一个和最后一个可以省略
	//s1 := fmt.Sprintf("%c", s[0])
	//fmt.Println(s1)
	//arr := []rune(s)
	//fmt.Println(len(arr)) // 获取字节长度
	//fmt.Println(arr[9])
	//
	//for _, n := range s {
	//	fmt.Printf("%c\n", n)
	//}
	fmt.Println(strings.Index(s, "l"))
	fmt.Println(strings.LastIndex(s, "l"))
	fmt.Println(strings.HasSuffix(s, "a"))
	fmt.Println(strings.HasSuffix(s, "张"))
	fmt.Println(strings.Contains(s, "ll"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.Replace(s, "m", "k", -1))
	fmt.Println(strings.Repeat(s, 12))
	fmt.Println(strings.Trim(s, "s"))
	fmt.Println(strings.Split(s, "m"))
	arr := []string{"small", "ming"}
	fmt.Println(strings.Join(arr, "-"))
}
