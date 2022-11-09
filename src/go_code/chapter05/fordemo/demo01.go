package main

import "fmt"

// for循环简单应用
func main() {
	var times int
	fmt.Println("请输入要循环的次数：")
	fmt.Scanf("%d", &times)
	for i := 0; i < times; i++ {
		fmt.Printf("%v\n", i)
	}

}
