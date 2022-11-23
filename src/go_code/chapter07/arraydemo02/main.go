package main

import "fmt"

func main() {
	var intArr [3]int
	fmt.Println(intArr)
	// 数组首地址就是第一个元素的地址
	fmt.Printf("intArr的地址%p intArr[0]的地址%p\n", &intArr, &intArr[0])

	// 四种数组初始化方式
	var numArr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr1)

	var numArr2 = [3]int{3, 2, 3}
	fmt.Println(numArr2)
	// [...]是固定的写法，不能更改，不写也可以
	var numArr3 = [...]int{4, 2, 3}
	fmt.Println(numArr3)

}
