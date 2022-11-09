package main

import "fmt"

// switch语句简单使用
func main() {
	var key byte
	fmt.Println("请输入字母：")
	fmt.Scanf("%c", &key)

	switch key {
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
