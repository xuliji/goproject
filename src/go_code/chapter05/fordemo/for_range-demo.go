package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str = "hello, world"

	// 字符串遍历方式一
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c %d\n", str[i], unsafe.Sizeof(str[i]))
	}

	// 字符串遍历方式二
	str = "abc~ok"
	for index, val := range str {
		fmt.Printf("index=%d val=%c\n", index, val)
	}
}
