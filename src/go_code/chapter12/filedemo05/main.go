package main

import "os"

// 把1.txt内容写到2.txt
func main() {
	file01, _ := os.ReadFile("src/go_code/chapter12/filedemo05/files/1.txt")
	err02 := os.WriteFile("src/go_code/chapter12/filedemo05/files/2.txt", file01, 777)
	if err02 != nil {
		panic(err02)
	}
}
