package main

import (
	"fmt"
	"unsafe"
)

// bool类型使用
func main() {
	var b bool = false
	fmt.Println(b)
	/*
		注意事项
		1. bool类型占用一个字节
		2. Go中bool类型只能是true和false
	*/
	fmt.Printf("b占用%d个字节", unsafe.Sizeof(b))
}
