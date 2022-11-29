package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\23844\\Downloads\\test.txt")

	if err != nil {
		fmt.Println("open file err=", err)
	}

	// 当函数退出时, 自动关闭file
	defer file.Close()

	// 创建一个  *Reader 带缓冲
	/*
		默认缓冲区4096
	*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到换行符就结束
		// io.EOF表示文件的末尾
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}
