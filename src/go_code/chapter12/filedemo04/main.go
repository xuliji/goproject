package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件打开方式  os.OpenFile()
func main() {
	file, err := os.OpenFile("src/go_code/chapter12/filedemo04/files/test.txt",
		os.O_CREATE|os.O_RDWR, 777)
	// 要关闭文件指针
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 写入5句话
	str := "Hello, Garden\n"
	// 带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer是带缓存的, 因此使用WriteString方法时
	// 内容先写到缓存要通过writer.Flush()写到文件中
	err = writer.Flush() // 必须要写
	if err != nil {
		fmt.Println("写入文件失败")
	}

}
