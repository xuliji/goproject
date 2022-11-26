package main

import (
	"fmt"
	"sort"
)

// golang中map默认是无序的,每次遍历得到的结果可能不一样
func main() {
	// 新版本自动排序列
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	// 现将key放入到切片中
	// 对切片排序
	// 遍历切片,按key值输出map值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k) // append()会自动创建空间,不用make
	}
	// sort包里面还有很多排序方法
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, map1[v])
	}
}
