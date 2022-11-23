package main

import "fmt"

// slice 切片
// 需要保存长度不定的数组
func main() {
	// 演示切片的基本使用
	intArr := [5]int{1, 2, 3, 4, 5}
	// 声明一个切片
	// slice引用intArr这个数组的第2-3个元素(没有做值拷贝)
	slice := intArr[1:3]
	fmt.Printf("slice长度%d\nslice内容%v\nslice容量%v", len(slice),
		slice, cap(slice)) // 切片的容量是动态变化的 一般是元素个数的2倍

}
