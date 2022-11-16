package main

import "fmt"

// init函数在main()函数之前被执行，用于初始化操作
func init() {
	fmt.Println("init()....")
}

func main() {
	fmt.Println("main()....")
}
