package main

import (
	"fmt"
	"os"
)

// ioutil.ReadFile() 一次性读取文件, 适用于小文件(已弃用)
// 现在是os.ReadFile()
func main() {
	file, err := os.ReadFile("src/go_code/chapter12/filedemo03/files/test.txt")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Printf("%v", file)
	// 不需要显式的open和close文件, 不需要去管理
}
