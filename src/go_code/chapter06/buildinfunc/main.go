package main

import "fmt"

// 内建函数new()的使用
func main() {
	num1 := 100
	fmt.Printf("num1的类型%T， num1的值%v， num1的地址%v\n", num1, num1, &num1)

	// new()返回的是指针
	num2 := new(int)
	fmt.Printf("num2的类型%T， num2的值%v， num2的地址%v， num2指向的值为%v\n", num2, num2, &num2, *num2)
}
