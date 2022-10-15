//GO程序执行入口为main()函数
//GO严格区分大小写
//GO编译器是一行一行编译的，一行最好只写一条语句
//GO定义的变量或者import的包如果没用用到代码编译不能通过

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	var num = 10
	fmt.Println(num)
}
