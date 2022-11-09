package main

import "fmt"

//	for 循环变量初始化；循环条件，循环变量迭代{
//	    循环语句操作
//	}
func main() {
	// for循环的第二种写法
	i := 0
	for i <= 10 {
		fmt.Printf("j=%v\n", i)
		i++
	}
	fmt.Println(i)
}
