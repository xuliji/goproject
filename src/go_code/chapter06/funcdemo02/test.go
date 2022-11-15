package main

import "fmt"

// 函数不能改变原始数据的值，相当于是复制了一份数据
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)
}
