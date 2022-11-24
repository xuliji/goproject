package main

import "fmt"

// 二维数组演示
func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 遍历打印数组
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < 4; i++ {
		fmt.Printf("%p\n", &arr[i])
	}

	// 二维数组遍历
	// 方式一:for循环 数组每一行个数可以不一样
	arr1 := [2][3]int{{1, 2, 3}, {4, 5}}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Print(arr1[i][j])
		}
		fmt.Println()
	}
	// 方拾二:for-range
	for i, ints := range arr1 {
		for j, num := range ints {
			fmt.Printf("arr1[%v][%v]=%v\t", i, j, num)
		}
		fmt.Println()
	}

}
