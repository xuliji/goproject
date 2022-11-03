package main

import "fmt"

// 基本数据类型转换
func main() {
	var i int32 = 100
	// 转float
	var f1 = float32(i)
	fmt.Printf("i=%v, f1=%f, f1 type=%T", i, f1, f1)
}
