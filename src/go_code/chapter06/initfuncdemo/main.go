package main

import "fmt"

// 源文件中先执行全局变量定义，再执行init()函数，再执行main()函数
var age = test()

func test() int {
	fmt.Println("test()....")
	return 90
}

// init函数在main()函数之前被执行，用于初始化操作
func init() {
	fmt.Println("init()....")
}

func main() {
	fmt.Println("main()....")
	fmt.Println(age)
}
