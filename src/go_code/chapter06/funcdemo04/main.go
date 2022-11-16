package main

import "fmt"

// 全局匿名函数
var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// 匿名函数使用方式
func main() {
	// 在定义匿名函数时就调用，这种方式只能调用一次
	// 用匿名函数球两个数字之和
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	// 将匿名函数赋值给一个变量, 此时a的数据类型就是函数类型
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 30)
	fmt.Println(res2)
	fmt.Println(fun1(40, 50))
}
