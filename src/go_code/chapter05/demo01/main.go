package main

import "fmt"

/*
顺序控制流程:按代码顺序执行
分支控制流程: if-else
循环控制流程:
*/
func main() {
	fmt.Println("第五章 流程控制\n顺序控制\n分支控制\n循环控制")
	var age byte
	fmt.Println("请输入你的年龄", &age)
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("你已经成年了,需要对自己的行为负责.")
	}
	if age1 := 20; age1 >= 18 {
		fmt.Println("你已经成年了,需要对自己的行为负责.")
	}
}
