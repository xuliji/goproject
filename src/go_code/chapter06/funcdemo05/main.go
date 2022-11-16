package main

import (
	"fmt"
	"funcdemo05/utils"
	"strconv"
)

// 闭包
// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += strconv.Itoa(x)
		fmt.Println(str)
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	// 返回的是一个匿名函数，但是这个匿名函数引用到了函数外的n
	// 所以这个函数就和n形成一个整体，构成闭包
	// 可以理解为，闭包是类，匿名函数时操作，n是字段
	func1 := utils.MakeSuffix(".jpg")
	name := func1("111111")
	fmt.Println(name)
	name = func1("22222.jpg")
	fmt.Println(name)

}
