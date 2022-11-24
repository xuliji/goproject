package main

import (
	"fmt"
	"goproject/src/go_code/chapter08/sortdemo/sorts"
)

func main() {
	slice := []int{34, 21, 12, 45, 11}
	sorts.Bubble(slice)
	fmt.Println(slice)
}
