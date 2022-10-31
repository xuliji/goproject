package main

import "fmt"

// 字符类型使用
// golang没有专门的字符类型,可以用byte保存
// Go的字符串不同,它是由字节组成的
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println(c1, c2)

	//如果想要输出字符,需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//输出中文字符可以用int rune存放
	var c3 rune = '北'
	fmt.Printf("c3=%c", c3)
}
