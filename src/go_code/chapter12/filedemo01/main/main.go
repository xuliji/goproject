package main

import (
	"fmt"
	"os"
)

// 文件操作演示
func main() {
	// f 叫 file对象 file指针 file文件句柄 都可以
	f, err := os.Open("../files/test.txt")
	if err != nil {
		fmt.Printf("%v", f)
	}
	f.Close()
}
