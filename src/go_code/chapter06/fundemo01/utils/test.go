package utils

import "fmt"

// 同一个包下不能存在相同函数名
// 如果要编译成一个可执行的程序文件，就必须将这个包声明为main
func Print_hello() {
	fmt.Println("hello")
}
