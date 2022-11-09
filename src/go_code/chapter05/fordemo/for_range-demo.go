package main

import (
	"fmt"
	"unsafe"
)

/*
传统方式对字符串的遍历是按照字节来遍历的，
而一个汉子在utf-8下是占3个字节的。
*/
func main() {
	var str = "hello, world！北京"
	str2 := []rune(str)
	// 字符串遍历方式一
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c %d\n", str2[i], unsafe.Sizeof(str[i]))
	}

	// 字符串遍历方式二 for-range
	//str = "abc~ok"
	for index, val := range str {
		fmt.Printf("index=%d val=%c\n", index, val)
	}
}
