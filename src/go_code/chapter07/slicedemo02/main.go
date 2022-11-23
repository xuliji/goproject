package main

import "fmt"

// 切片的定义方式
// 方式一：让切片去索引一个已经创建好的数组
// 方式二：通过make来创建切片
func main() {
	// 对于切片必须先make，再使用
	var slice []int = make([]int, 10)
	fmt.Println(slice, cap(slice))
	// make创建的切片对应的数组是由make底层维护的，对外不可见
	// 只能通过slice访问
	for i, v := range slice {
		fmt.Println(i, v)
	}
}
