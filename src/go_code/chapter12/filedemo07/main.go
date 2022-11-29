package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件拷贝演示  func Copy(dst Writer, src Reader) (written int64, err error)
func CopyFile(src_Name, dst_Name string) (wrritten int64, err error) {
	src, _ := os.OpenFile(src_Name, os.O_RDONLY, 777)
	dsc, _ := os.OpenFile(dst_Name, os.O_CREATE|os.O_WRONLY, 777)
	// 关闭文件
	defer func() {
		src.Close()
		dsc.Close()
	}()

	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dsc)
	written, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(written)
	return written, err
}

func main() {
	src_path := "src/go_code/chapter12/filedemo07/file01/1.png"
	dst_path := "src/go_code/chapter12/filedemo07/file02/2.png"
	CopyFile(src_path, dst_path)
}
