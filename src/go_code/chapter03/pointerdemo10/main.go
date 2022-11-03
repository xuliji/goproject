package main

import "fmt"

// go中指针类型使用
func main() {
	var i int = 10
	// i的地址是，&i
	fmt.Println("i的地址=", &i)

	// 指针类型，指针变量存的是个地址，这个地址指向的空间才是值
	// &取地址  *取值
	// 比如：var ptr *int = &num
	/*
		    1. ptr是一个指针变量
		    2. ptr类型是*int
			3. ptr本身的值是&i
	*/
	var ptr *int = &i
	fmt.Printf("ptr 的值%v\n", ptr)
	fmt.Printf("ptr 的地址%v\n", &ptr)
	fmt.Printf("ptr指向的值%v\n", *ptr)
}
