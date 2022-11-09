package main

import "fmt"

// 测试函数
func test(b byte) byte {
	return b + 1
}

// switch语句简单使用
func main() {
	var key byte
	fmt.Println("请输入字母：")
	fmt.Scanf("%c", &key)

	// case后面是表达式，可以是常量、变量、一个具有返回值的函数

	switch test(key) + 1 {
	case 'a':
		fmt.Println("这是周一")
	case 'b':
		fmt.Println("这是周二")
	case 'c':
		fmt.Println("这是周三")
	case 'd':
		fmt.Println("这是周四")
	case 'e':
		fmt.Println("这是周五")
	default:
		fmt.Println("输入有误")
	}

}
