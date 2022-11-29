package main

import (
	"fmt"
	"os"
)

// 判断文件存不存在
// 如果err为nil表示文件或文件夹存在
// 如果os.IsNotExist(err)为true表示文件或文件夹不存在
// 其余情况,则不确定是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	path := "src/go_code/chapter12/filedemo06/files/1.txt"
	flag, err := PathExist(path)
	fmt.Println(flag, err)
}
