package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello你好"
	// 统计字符串长度， 按字节返回
	length := len(str)
	fmt.Println("length=", length)

	// 遍历带中文的字符串
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	// 字符串转整数
	// 整数转字符串
	n, err := strconv.Atoi("20")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println(n)
	}

	str2 := strconv.Itoa(30)
	fmt.Printf("str2=%v str2=%T\n", str2, str2)

	// 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("btyes=%v\n", bytes)
	// byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	// 10进制转2，8，16进制 strconv.FormatInt()来完成
	str = strconv.FormatInt(123, 16)
	fmt.Println("123对应的二进制是:", str)

	// 查找字串是不是在指定的字符串中
	flag := strings.Contains("seafood", "foo")
	fmt.Printf("%t\n", flag)

	// 查找一个字符串中字串得个数
	count := strings.Count("abababababa", "aba")
	fmt.Println(count)

	// 不区分大小写字母的字符串比较(==是区分字符大小写的)：strings.EqualFold("abc", "ABC")
	fmt.Println(strings.EqualFold("abc", "ABC"))
}
