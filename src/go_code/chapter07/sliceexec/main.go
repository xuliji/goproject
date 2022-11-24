package main

import (
	"errors"
	"fmt"
)

// 返回斐波那契数
func fbn(n int) []uint64 {
	defer func() {
		err := recover()
		if err != nil {
			panic(errors.New("请检查n的范围！！！"))
		}
	}()
	slice := make([]uint64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}

func main() {
	fmt.Println(fbn(10))
}
