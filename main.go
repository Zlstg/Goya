package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("统计字符串的长度，按字节 len(str)")
	fmt.Println("golang的编码统一为utf-8（ascii的字符（字母和数字）占一个子节，汉字占三个子节）")
	fmt.Println("hello len=", len("hello"))
	str := "hello北"
	fmt.Println(str+" len=", len(str))
	//
	fmt.Println("字符串遍历，同时处理有中文的问题 r := []rune(str)")
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\t", r[i])
	}
	//
	fmt.Println("字符串转整数： n, err := strconv.Atoi(\"12\")")
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}
	nn, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", nn)
	}
}