package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file, err := os.Open("src/go_code/chapter12/filedemo08/files/1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	charcount := CharCount{}
	reader := bufio.NewReader(file)
	// 循环读取内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case (v >= 'a') && (v <= 'z'):
				fallthrough
			case v >= 'A' && v <= 'Z':
				charcount.ChCount++
			case v == ' ' || v == '\t':
				charcount.SpaceCount++
			case v >= '0' && v <= '9':
				charcount.NumCount++
			default:
				charcount.OtherCount++
			}
		}
	}
	fmt.Println(charcount)

}
