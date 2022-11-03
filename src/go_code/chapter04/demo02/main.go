package main

import "fmt"

// 逻辑运算符 用于连接多个条件,最终结果是一个bool值
// && 逻辑与  || 逻辑或  ! 逻辑非
func main() {
	var age = 40
	if age > 30 && age < 50 {
		fmt.Println("中年")
	}
	var a, b = 10, 20
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
