package main

import "fmt"

func main() {
	slice3 := make([]int, 10)
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
	// append()是创建了一个新的数组来保存追加的内容，需要赋值给slice3
	slice3 = append(slice3, 400, 500, 600)
	// append()也可以追加切片
	slice3 = append(slice3, slice3...)
	for i, v := range slice3 {
		fmt.Println(i, v)
	}

	// copy() 切片的拷贝
	// 修改slice3的数据slice4不会改变
	slice4 := make([]int, 50)
	copy(slice4, slice3)
	fmt.Println(slice4)
}
