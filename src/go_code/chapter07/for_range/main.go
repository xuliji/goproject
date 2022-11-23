package main

import "fmt"

// for range第一个返回值index是数组下标
// 第二个是value是该下标位置的值
// index,value都是for循环中的局部变量

func main() {
	intArr := [30]bool{}

	for i, v := range intArr {
		fmt.Println(i, v)
	}
}
