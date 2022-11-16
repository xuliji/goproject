package main

import (
	"fmt"
	"goproject/src/go_code/chapter06/initfuncdemo/utils"
)

// 源文件中先执行全局变量定义，再执行init()函数，再执行main()函数
var age = test()

func test() int {
	fmt.Println("test()....")
	return utils.Age
}

// init函数在main()函数之前被执行，用于初始化操作
func init() {
	fmt.Println("init()....")
}

func main() {
	fmt.Println("main()....")
	fmt.Println(age)
	fmt.Println(utils.Name)
}
