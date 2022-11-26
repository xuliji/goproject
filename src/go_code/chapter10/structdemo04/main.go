package main

import "fmt"

type A struct {
	Num int
}
type B struct {
	Num int
}

func main() {
	a := A{}
	b := B{}
	// 强转需要字段名字和类型相同
	a = A(b)
	fmt.Println(a)
}
