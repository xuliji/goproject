package main

import "fmt"

// 函数不能改变原始数据的值，相当于是复制了一份数据
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

// getSum
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)
	n2 := 30
	sum := getSum(n1, n2)
	fmt.Println(sum)
}
