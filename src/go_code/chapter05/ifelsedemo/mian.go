package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("你已经成年了,需要对自己的行为负责.")
	} else {
		fmt.Println("今天就放过你了.")
	}
}
